package backend

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"
)

func Update(version string) error {

	updater, _ := selfupdate.NewUpdater(selfupdate.Config{Validator: &selfupdate.ChecksumValidator{UniqueFilename: "checksums.txt"}})
	latest, found, err := updater.DetectLatest("BitsOfAByte/MarketXIV")

	// Unknown error occurred, abort update process.
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %v", err)
	}

	// Specified OS or Architechture is not supported.
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from GitHub repository", runtime.GOOS, runtime.GOARCH)
	}

	// No update is available for the current version.
	if latest.LessOrEqual(version) {
		fmt.Printf("You currently have the latest version! (%s)", version)
		return nil
	}

	// Find the current executable's path.
	exe, err := os.Executable()

	//  Could not find the executable's path, abort update process.
	if err != nil {
		return errors.New("could not locate executable path")
	}

	// Perform the update.
	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %v", err)
	}

	fmt.Printf("Successfully updated to version %s (Arch: %s, OS: %s) ", latest.Version(), latest.Arch, latest.OS)

	return nil
}
