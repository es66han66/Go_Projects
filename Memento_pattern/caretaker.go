package main

type SaveFileManager struct {
	mementos []Memento
}

func (s *SaveFileManager) Save(m Memento) {
	s.mementos = append(s.mementos, m)
}
func (s *SaveFileManager) Load(index int) Memento {
	return s.mementos[index]
}
