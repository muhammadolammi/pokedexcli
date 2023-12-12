package main

import "os"

func exitCallback(cfg *config) error {
	os.Exit(0)
	return nil
}
