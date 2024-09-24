package server

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "sync"
    "net/pkg/chat"
    "net/pkg/common"
)

type Server struct {
    listenAddr string
    ln         net.Listener
    clients    map[net.Conn]string
    mu         sync.Mutex
    messages   []string
    quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
    return &Server{
        listenAddr: listenAddr,
        clients:    make(map[net.Conn]string),
        messages:   []string{},
        quitch:     make(chan struct{}),
    }
}

func (s *Server) Start() error {
    ln, err := net.Listen("tcp", s.listenAddr)
    if err != nil {
        return err
    }
    defer ln.Close()
    fmt.Printf("Listening on the port %s\n", s.listenAddr)
    s.ln = ln

    go s.acceptLoop()

    <-s.quitch
    return nil
}

func (s *Server) acceptLoop() {
    for {
        conn, err := s.ln.Accept()
        if err != nil {
            fmt.Println("Accept error: ", err)
            continue
        }

        go s.readLoop(conn)
    }
}

func (s *Server) readLoop(conn net.Conn) {
    defer conn.Close()
    name, err := s.getName(conn)
    if err != nil {
        return
    }

    s.mu.Lock()
    s.clients[conn] = name
    s.mu.Unlock()

    s.broadcast(fmt.Sprintf("%s has joined the chat...\n", name))

    s.sendPreviousMessages(conn)

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := scanner.Text()
        if text == "" {
            continue
        }

        formattedMessage := chat.FormatMessage(name, text)
        s.messages = append(s.messages, formattedMessage)
        s.broadcast(formattedMessage)
    }

    s.mu.Lock()
    delete(s.clients, conn)
    s.mu.Unlock()
    s.broadcast(fmt.Sprintf("%s has left the chat...\n", name))
}

func (s *Server) getName(conn net.Conn) (string, error) {
    conn.Write([]byte(common.WelcomeMessage(s.getClientsName())))
    name, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        conn.Write([]byte("Invalid name. Disconnection....\n"))
        return "", err
    }
    return chat.TrimName(name), nil
}

func (s *Server) broadcast(msg string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for conn := range s.clients {
        if _, err := conn.Write([]byte(msg)); err != nil {
            log.Println("Error writing to client: ", err)
        }
    }
}

func (s *Server) sendPreviousMessages(conn net.Conn) {
    for _, msg := range s.messages {
        conn.Write([]byte(msg))
    }
}

func (s *Server) getClientsName() []string {
    s.mu.Lock()
    defer s.mu.Unlock()
    names := make([]string, 0, len(s.clients))
    for _, name := range s.clients {
        names = append(names, name)
    }
    return names
}