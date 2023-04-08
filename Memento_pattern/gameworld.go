package main

type GameWorld struct {
	level int
	score int
}

func (g *GameWorld) Save() Memento {
	return Memento{level: g.level, score: g.score}
}

func (g *GameWorld) Restore(m Memento) {
	g.level = m.level
	g.score = m.score
}
