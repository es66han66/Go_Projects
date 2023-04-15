package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{
		name:      "Plato",
		leftFork:  4,
		rightFork: 0,
	},
	{
		name:      "Socrates",
		leftFork:  0,
		rightFork: 1,
	},
	{
		name:      "Aristotle",
		leftFork:  1,
		rightFork: 2,
	},
	{
		name:      "Pascal",
		leftFork:  2,
		rightFork: 3,
	},
	{
		name:      "Locke",
		leftFork:  3,
		rightFork: 4,
	},
}

var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second

func main() {

	fmt.Println("Dining philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("Table is empty")

	dine()

	fmt.Println("Table is empty")

}

func dine() {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	// wg is for everyone eating, i.e. it becomes 0 when everyone is done eating
	wg := &sync.WaitGroup{}

	wg.Add(len(philosophers))

	// we want everyone to be seated before they start eating, so we create a wait group and set it to 5

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	// wait for philosophers to finish. This blocks until wait group is 0.
	wg.Wait()

}

/*
diningProblem is the function fired off as a goroutine for each of our philosophers. It takes one philosopher, one waitgroup
to determine when everyone is done, a map consisting the mutexes for every fork on the table, and a waitgroup used to pause execution
of every instance of this goroutine until everyone is seated at table
*/
func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at table
	fmt.Printf("%s is seated at table.\n", philosopher.name)
	seated.Done()

	// wait until everyone is seated
	seated.Wait()

	// eat hunger aka 3 times
	for i := hunger; i > 0; i-- {
		/*
			get a lock on both forks. We have to choose the lower numbered fork first in order to avoid logical race
			condition, which is not detected by -race flag. If we don't do this, we have the potential for a deadlock, since 2
			philosophers will wait endlessly for same fork. Note that goroutine will block(pause) until it gets a lock on both the
			right and left forks
		*/
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)

		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s has put down the fork.\n", philosopher.name)

	}

	fmt.Println(philosopher.name, " is satisfied.")
	fmt.Println(philosopher.name, " left the table.")

}
