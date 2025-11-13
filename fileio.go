package main

import (
	"os"
	"strings"
)

// loadMap loads a map from a file where each line is "key:value".
func loadMap(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		m[value] = key // Note: value (prefix) -> key (name)
	}
	return m, nil
}
