package main

import "os"

func exitCallback(cfg *config, area string) error {
	os.Exit(0)
	return nil
}
