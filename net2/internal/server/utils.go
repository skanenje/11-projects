package server

import (
    "net/pkg/chat"
)

func formatMessage(name, text string) string {
    return chat.FormatMessage(name, text)
}