package chat

import (
    "fmt"
    "strings"
    "net/pkg/common"
    "time"
)

func FormatMessage(name, text string) string {
    timestamp := time.Now().Format(common.TimestampFormat)
    return fmt.Sprintf("[%s][%s]: %s\n", timestamp, name, text)
}

func TrimName(name string) string {
    return strings.TrimSpace(name)
}