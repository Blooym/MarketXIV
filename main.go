/*
Copyright © 2022 BitsOfAByte

*/

package main

import (
	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/BitsOfAByte/marketxiv/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			backend.Fatal(r)
		}
	}()

	cmd.Execute()
}
