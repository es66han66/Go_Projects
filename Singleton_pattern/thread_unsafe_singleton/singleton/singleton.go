package singleton

type Singleton struct {
	name string
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{name: "Golang Singleton"}
	}
	return instance
}

func (s *Singleton) GetName() string {
	return s.name
}
