package singleton

import "sync"

type Singleton struct {
    name string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{name: "Safe Golang Singleton"}
    })
    return instance
}

func (s *Singleton) GetName() string {
    return s.name
}