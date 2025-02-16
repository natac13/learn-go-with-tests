package main

import (
	"os"
	"time"

	"github.com/natac13/learn-go-with-tests/maths/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
