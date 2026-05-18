package config

import (
	"fmt"
	"strings"
)

func ParseKV(line string) (string, string, error) {
	key, value, found := strings.Cut(line, "=")
	if !found {
		return "", "", fmt.Errorf("Invalid line %s", line)
	}
	return strings.ToLower(strings.TrimSpace(key)), strings.TrimSpace(value), nil
}
