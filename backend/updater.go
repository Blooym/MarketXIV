/*
Copyright © 2022 Blooym

MIT License, see the LICENSE file for more information.
*/
package backend

import (
	"fmt"
	"os"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"
)

// returns string or error
func Update(version string) {

	updater, _ := selfupdate.NewUpdater(selfupdate.Config{Validator: &selfupdate.ChecksumValidator{UniqueFilename: "checksums.txt"}})
	latest, found, err := updater.DetectLatest("Blooym/MarketXIV")

	// Unknown error occurred, abort update process.
	if err != nil {
		Error(fmt.Sprintf("error occurred while detecting version: %v", err))
		return
	}

	// Specified OS or Architechture is not supported.
	if !found {
		Warning(fmt.Sprintf("version %s is not supported on %s/%s", version, runtime.GOOS, runtime.GOARCH))
		return
	}

	// No update is available for the current version.
	if latest.LessOrEqual(version) {
		fmt.Printf("no update available for version %s\n", version)
		return
	}

	// Find the current executable's path.
	exe, err := os.Executable()

	//  Could not find the executable's path, abort update process.
	if err != nil {
		Error(fmt.Sprintf("error occurred while finding executable's path: %v", err))
		return
	}

	// Perform the update.
	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		Error(fmt.Sprintf("error occurred while updating: %v", err))
		return
	}

	fmt.Printf("Successfully updated to version %s (OS: %s, Arch: %s) from %s\n", latest.Version(), latest.OS, latest.Arch, latest.PublishedAt)
}
