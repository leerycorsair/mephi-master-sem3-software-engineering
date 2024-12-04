package model

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Model struct {
	rs map[string]Resource
}

func NewEmptyModel() *Model {
	return &Model{
		rs: map[string]Resource{},
	}
}

func (m *Model) IsEmpty() bool {
	return len(m.rs) == 0
}

func (m *Model) AddResource(nametag string, s ResourceStatus) error {
	r := NewResource(s)
	m.rs[nametag] = r
	return nil
}

func (m *Model) FreeResource(nametag string) error {
	r, ok := m.rs[nametag]
	if !ok {
		return ErrResourceNotFound
	}
	r.Free()
	m.rs[nametag] = r
	return nil
}

func (m *Model) OccupyResource(nametag string) error {
	r, ok := m.rs[nametag]
	if !ok {
		return ErrResourceNotFound
	}
	err := r.Occupy()
	if err != nil {
		return err
	}
	m.rs[nametag] = r
	return nil
}

func (m *Model) InitResources(resCnt int) {
	for i := 0; i < resCnt; i++ {
		m.AddResource("res"+strconv.Itoa(i), Free)
	}
}

func (m *Model) SearchResource(nametag string) (Resource, error) {
	r, ok := m.rs[nametag]
	if ok {
		return r, nil
	}
	return r, ErrResourceNotFound
}

func (m *Model) InitFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return ErrModelIncorrectFile
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return ErrModelIncorrectFile
		}
		resTag := strings.TrimSpace(parts[0])
		resStatus := strings.TrimSpace(parts[1])

		var status ResourceStatus

		switch resStatus {
		case string(Free):
			status = Free
		case string(Occupied):
			status = Occupied
		default:
			return ErrModelIncorrectFile
		}

		m.AddResource(resTag, status)
	}

	if err := scanner.Err(); err != nil {
		return ErrModelIncorrectFile
	}
	return nil
}
