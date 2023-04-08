package main

import "fmt"

func main() {
	gameWorld := &GameWorld{level: 1, score: 0}
	saveFileManager := &SaveFileManager{}
	// Save the initial state of the game
	saveFileManager.Save(gameWorld.Save())
	// Increase the level and score
	gameWorld.level++
	gameWorld.score += 1000
	// Save the new state of the game
	saveFileManager.Save(gameWorld.Save())
	// Restore the previous state of the game
	gameWorld.Restore(saveFileManager.Load(0))
	fmt.Println(gameWorld)
}
