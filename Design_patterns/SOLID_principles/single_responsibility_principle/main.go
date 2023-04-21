package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

/*
	breaks SRP because below we are doing stuff which is not the responsibility of journal type, it's responsibility was to maintain
	entries i.e. add and remove entries, the saving and read from file can be happening from other types too and it should be a
	separate concern not involved with journal type
*/

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

/*
below is the approach to separate concerns with using a method in which we will pass the desired type from which we will take out data
for this approach we are using a lineSeparator variable declared outside which can be hardcoded or reassigned everytime we need to use it
which is again not the right approach.
*/
var lineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

/*
below is the approach which actually resolves our purpose by utilizing a struct which in itself has a lineSeparator field which helps
us to utilize any value of line separator depending upon our usecase and then passing the desired type into the method attached to it
where we can store and load file
*/
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")
	fmt.Println(strings.Join(j.entries, "\n"))

	// separate function
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal1.txt")
}
