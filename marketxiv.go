/*
Copyright © 2022 BitsOfAByte

*/

package main

import (
	"fmt"
	"os"

	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/BitsOfAByte/marketxiv/cmd"
)

func main() {

	// Check if the user is running a preview version (So regular users don't get distributed it)
	if backend.IsPreview() {
		fmt.Println("Detected a preview version of MarketXIV! If you are a user, please download a stable version from GitHub directly.")
		fmt.Println("If you are certain you want to run this preview build, set environment variable 'MXIV_PREVIEW_CONSENT' to 'true'")
		os.Exit(1)
	}
	// Catches any panics and logs them before exiting
	defer func() {
		if r := recover(); r != nil {
			backend.Fatal(r)
		}
	}()

	cmd.Execute()
}
