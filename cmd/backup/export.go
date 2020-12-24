package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xn3cr0nx/pg-backup/internal/exporter"
	"github.com/xn3cr0nx/pg-backup/pkg/logger"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export dump",
	Long:  `Dump pg using pg_dump command and export to supported exports`,
	Run: func(cmd *cobra.Command, args []string) {
		var e exporter.Exporter

		switch viper.GetString("target") {
		case "s3":
			e = exporter.NewS3Exporter()
		default:
			e = exporter.NewFileExporter()
		}

		if err := e.Export(); err != nil {
			logger.Error("Backup", err, logger.Params{})
			os.Exit(1)
		}
	},
}
