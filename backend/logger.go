package backend

import (
	"log"
	"os"
	"strings"
)

var Verbose string = strings.ToLower(os.Getenv("MXIV_VERBOSE"))

// Logs a message if Verbose is true.
func verboseLog(message string) {
	if Verbose == "true" {
		log.Println(message)
	}
}
