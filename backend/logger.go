/*
Copyright © 2022 BitsOfAByte

*/

package backend

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

/**
 * The logger shouldn't be used in situations where the program just needs to output exactly what was asked
 * It should be used for debug messages, warnings & errors as this allows for easy tracking of errors.
 */

func verboseEnabled() bool {
	return viper.GetBool("app.verbose")
}

// Outputs a debug message if debug environment
func Debug(message string) {
	if verboseEnabled() {
		log.Printf("Debug: %v\n", message)
	}
}

// Outputs a warning message.
func Warning(message interface{}) {
	fmt.Printf("Warning: %v\n", message)
}

// Logs an error and continues the program.
func Error(message interface{}) {
	fmt.Printf("Error: %v\n", message)
}

// Logs a fatal error and exits the program.
func Fatal(message interface{}) {
	fmt.Printf("Fatal: %v\n", message)
	os.Exit(1)
}
