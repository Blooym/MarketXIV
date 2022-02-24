package backend

import (
	"log"
)

// Enables logger if true, disabled during production build if specified
var Verbose string = "true"

// Logs a message if Verbose is true.
func verboseLog(message string) {
	if Verbose == "true" {
		log.Println(message)
	}
}
