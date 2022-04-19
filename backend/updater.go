package backend

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"
)

func Update(version string) error {
	latest, found, err := selfupdate.DetectLatest("BitsOfAByte/MarketXIV")

	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %v", err)
	}

	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from GitHub repository", runtime.GOOS, runtime.GOARCH)
	}

	if latest.LessOrEqual(version) {
		fmt.Printf("You currently have the latest version! (%s)", version)
		return nil
	}

	exe, err := os.Executable()

	if err != nil {
		return errors.New("could not locate executable path")
	}

	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %v", err)
	}

	fmt.Printf("Successfully updated to version %s", latest.Version())

	return nil
}
