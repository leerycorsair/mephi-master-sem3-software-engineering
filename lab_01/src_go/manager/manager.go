package manager

import (
	"fmt"
	"resourceRegistrator/model"
	"slices"
)

type Manager struct {
	requests RequestSlice
	model    *model.Model
}

func NewManager(m *model.Model) *Manager {
	return &Manager{
		requests: make(RequestSlice, 0),
		model:    m,
	}
}

func (m *Manager) CheckModel() bool {
	return m.model.IsEmpty()
}

func (m *Manager) SetModel(model *model.Model) {
	m.model = model
}

func (m *Manager) AddRequest(r Request) error {
	_, err := m.model.SearchResource(r.resourceId)
	if err != nil {
		return err
	}
	m.requests = append(m.requests, r)
	return nil
}

func (m *Manager) FreeResource(resId string) error {
	err := m.model.FreeResource(resId)
	if err != nil {
		return err
	}
	reqId := slices.IndexFunc(m.requests, func(r Request) bool {
		return r.GetResId() == resId && r.IsActive()
	})
	if reqId >= 0 {
		m.requests = slices.Delete(m.requests, reqId, reqId+1)
	}
	return nil
}

func (m *Manager) Work() {
	for i, r := range m.requests {
		if r.IsWaiting() {
			err := m.model.OccupyResource(r.GetResId())
			if err == model.ErrResourceIsBusy {
				continue
			}
		}
		m.requests[i].MakeActive()
		m.requests[i].RequestProcess()
		if m.requests[i].IsOver() {
			m.model.FreeResource(r.GetResId())
		}
	}
	m.requests = slices.DeleteFunc(m.requests, Request.IsOver)
}

func (m *Manager) InitModel(resCnt int) {
	m.model.InitResources(resCnt)
}

func (m *Manager) InitFromFile(path string) error {
	return m.model.InitFromFile(path)
}

func (m *Manager) RequestsInfo() string {
	return fmt.Sprint(m.requests)
}

func (m *Manager) ModelInfo() string {
	return fmt.Sprint(m.model)
}
