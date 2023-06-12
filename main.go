/*
Copyright © 2022 Blooym

*/

package main

import (
	"github.com/Blooym/marketxiv/backend"
	"github.com/Blooym/marketxiv/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			backend.Fatal(r)
		}
	}()

	cmd.Execute()
}
