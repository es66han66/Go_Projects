package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "hello world"

	wg.Add(2)

	go updateMessage("eshan")
	go updateMessage("eshan1")

	wg.Wait()

	if msg != "eshan" {
		t.Error("expected msg to be eshan but got", msg)
	}
}
