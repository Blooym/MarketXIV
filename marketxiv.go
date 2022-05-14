/*
Copyright © 2022 BitsOfAByte

*/

package main

import (
	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/BitsOfAByte/marketxiv/cmd"
)

func main() {
	// Catches any panics and logs them before exiting
	defer func() {
		if r := recover(); r != nil {
			backend.Fatal(r)
		}
	}()

	cmd.Execute()
}
