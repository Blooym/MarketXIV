/*
Copyright © 2022 BitsOfAByte

*/

package backend

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/BitsOfAByte/marketxiv/build"
	"github.com/nedpals/supabase-go"
	"github.com/spf13/viper"
)

var client *supabase.Client

// Returns the supabase client.
func getClient() *supabase.Client {
	if client == nil {
		client = supabase.CreateClient(build.SUPABASE_URL, build.SUPABASE_KEY)
	}

	return client
}

// Checks to see if analytics are enabled.
func analyticsEnabled() bool {
	return viper.GetBool("analytics.enabled")
}

// Uploads a log to the supabase analytics database.
func SendLog(logType string, message interface{}) {

	if !analyticsEnabled() {
		return
	}

	client := getClient()
	Debug("Uploading data to log table...")
	client.DB.From("logs").Insert(map[string]interface{}{
		"command":      "marketxiv " + strings.Join(os.Args[1:], " "),
		"log_type":     logType,
		"log_message":  message,
		"arch_os":      fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		"mxiv_version": fmt.Sprintf("%s/%s", build.Version, build.Commit),
		"uuid":         viper.GetString("analytics.uuid"),
	}).Execute(interface{}(nil))
}
