andpackage backend

import (
	"fmt"
	"os"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"
)

// returns string or error
func Update(version string) string {

	updater, _ := selfupdate.NewUpdater(selfupdate.Config{Validator: &selfupdate.ChecksumValidator{UniqueFilename: "checksums.txt"}})
	latest, found, err := updater.DetectLatest("BitsOfAByte/MarketXIV")

	// Unknown error occurred, abort update process.
	if err != nil {
		return fmt.Sprintf("error occurred while detecting version: %v", err)
	}

	// Specified OS or Architechture is not supported.
	if !found {
		return fmt.Sprintf("latest version for %s/%s could not be found from GitHub repository", runtime.GOOS, runtime.GOARCH)
	}

	// No update is available for the current version.
	if latest.LessOrEqual(version) {
		return "No updates are available, you are already running the latest version."
	}

	// Find the current executable's path.
	exe, err := os.Executable()

	//  Could not find the executable's path, abort update process.
	if err != nil {
		return "Could not locate executable path"
	}

	// Perform the update.
	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Sprintf("Error occurred while updating binary: %v", err)
	}

	return fmt.Sprintf("Successfully updated to version %s (OS: %s, Arch: %s) from %s", latest.Version(), latest.OS, latest.Arch, latest.PublishedAt)
}
