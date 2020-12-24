package naming

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Output returns the file output name based on set env variables
func Output() string {
	name := ""
	path := viper.GetString("output_path")
	if path != "" {
		name += path
	}
	prefix := viper.GetString("output_prefix")
	if prefix != "" {
		name += prefix
		if viper.GetBool("output_time") {
			name += "_"
		}
	}
	if viper.GetBool("output_time") || name == "" {
		name += currentMigrationTime()
	}
	return name + "." + viper.GetString("output_ext")
}

func currentMigrationTime() string {
	curr := time.Now()
	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d", curr.Year(), int(curr.Month()), curr.Day(),
		curr.Hour(), curr.Minute(), curr.Second())
}
