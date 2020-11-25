package main

import (
	"github.com/alexellis/deploylister/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
