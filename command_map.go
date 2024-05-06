package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = res.Next
	cfg.prevLocationURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}
	res, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = res.Next
	cfg.prevLocationURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
