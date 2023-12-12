package main

import "os"

func exitCallback() error {
	os.Exit(0)
	return nil
}
