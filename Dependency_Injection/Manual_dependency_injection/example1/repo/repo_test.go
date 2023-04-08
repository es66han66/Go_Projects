package repo

import (
	"errors"
	"testing"
)

type MockDB struct {
	GetFunc func(id string) (string, error)
}

func (m *MockDB) Get(id string) (string, error) {
	return m.GetFunc(id)
}

func TestRepository_Get(t *testing.T) {
	mockDB := &MockDB{
		GetFunc: func(id string) (string, error) {
			if id != "123" {
				return "", errors.New("Not found")
			}
			return "Data", nil
		},
	}
	repo := NewRepository(mockDB)
	data, err := repo.Get("123")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if data != "Data" {
		t.Errorf("expected data to be 'Data', got '%v'", data)
	}
}
