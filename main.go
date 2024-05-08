package main

import (
	"fmt"
	"os"
	"time"
	"math/rand"
	"github.com/ArrayOfLilly/pokedexcli/internal/pokeapi"
)

type config struct{
	pokeapiClient pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Error: too many arguments")
		fmt.Println("Usage: pokedex")
		return
	}

	// config
	cfg := config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Minute*5),
	}

	// greeting
	fmt.Printf("\nWelcome to the 🐙 Pokedex!\n\n")
	
	// REPL
	startREPL(&cfg)

	rand.Seed(time.Now().UnixMilli())
}
