package joke_test

import (
	"testing"

	"github.com/gabrie30/joke/pkg/joke"
)

func TestJokeFetch(t *testing.T) {
	dbSetup()
	joke.FetchIfNeeded()
}
