package poker_test

import (
	"strings"
	"testing"

	poker "github.com/natac13/learn-go-with-tests/build-an-application"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {

		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Chris"
		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Cleo"
		poker.AssertPlayerWin(t, playerStore, want)
	})
}
