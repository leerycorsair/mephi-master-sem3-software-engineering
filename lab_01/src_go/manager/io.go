package manager

import (
	"fmt"
	"strings"
)

func (r Request) String() string {
	return fmt.Sprintf("\tRequest ResourseId=%s, Time Left=%d, Status=%s", r.resourceId, r.time, r.status)
}

type RequestSlice []Request

func (rs RequestSlice) String() string {
	strs := make([]string, len(rs))
	for i, r := range rs {
		strs[i] = fmt.Sprintln(r)
	}
	return fmt.Sprintf("[\n%s\n]", strings.Join(strs, "\n"))
}
