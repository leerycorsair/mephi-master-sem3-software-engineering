package model

import (
	"fmt"
	"strings"
)

func (m Model) String() string {
	strs := make([]string, 0, len(m.rs))
	for key, value := range m.rs {
		strs = append(strs, fmt.Sprintf("Resourse Tag=%s, %v", key, value))
	}
	return fmt.Sprintf("[\n%s\n]", strings.Join(strs, "\n"))
}

func (r Resource) String() string {
	return fmt.Sprintf("Status=%s", r.status)
}
