package joke

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gabrie30/joke/configs"
)

// Tell tells n jokes
func Tell(count int) {
	var jokeIDs []int

	if configs.LastNJokesToTell > 0 {
		jokeIDs = LastNJokeIDs(count)
	} else {
		// find a random number between 1 and lastjoke id
		jokeIDs = randomJokeIDs(count)
	}

	for _, id := range jokeIDs {
		j := GetJokeByID(id)
		fmt.Printf("================== #%v ==================\n", j.ID)
		fmt.Println(j.Setup)
		fmt.Println(j.Punchline)
		fmt.Println("")
	}
}

func randomJokeIDs(count int) []int {
	counts := map[int]bool{}
	min := 1
	max := LastJokeID()

	if count > max {
		count = max
	}

	for len(counts) < count {
		rand.Seed(time.Now().UnixNano())
		seed := rand.Intn((max - min + 1) + min)
		if seed == 0 {
			continue
		}

		counts[seed] = true
	}

	retKeys := []int{}

	for k := range counts {
		retKeys = append(retKeys, k)
	}

	return retKeys
}
