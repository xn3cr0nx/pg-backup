package pg

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

// Dump run pg_dump command
func Dump() (string, error) {
	dump, err := exec.Command("pg_dump", connectionString()).Output()
	if err != nil {
		return "", err
	}
	return string(dump), err
}

func connectionString() string {
	// os.Setenv("PGPASSWORD", viper.GetString("pg_pass"))
	// conn := fmt.Sprintf("-h %s -p %s -U %s %s", viper.GetString("pg_host"), viper.GetString("pg_port"), viper.GetString("pg_user"), viper.GetString("pg_db"))
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", viper.GetString("pg_user"), viper.GetString("pg_pass"), viper.GetString("pg_host"), viper.GetString("pg_port"), viper.GetString("pg_db"))
}
