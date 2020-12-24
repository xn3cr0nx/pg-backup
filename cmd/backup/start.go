package main

import (
	"os"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xn3cr0nx/pg-backup/internal/exporter"
	"github.com/xn3cr0nx/pg-backup/pkg/logger"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Export dump cron",
	Long:  `Dump pg periodically using pg_dump command`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Backup", "pg-backup service starting", logger.Params{})

		var e exporter.Exporter

		switch viper.GetString("target") {
		case "s3":
			e = exporter.NewS3Exporter()
		default:
			e = exporter.NewFileExporter()
		}

		logger.Info("Backup", "Running pg-backup", logger.Params{"target": viper.GetString("target")})
		if err := e.Export(); err != nil {
			logger.Error("Backup", err, logger.Params{})
			os.Exit(1)
		}

		c := cron.New()
		defer c.Stop()

		logger.Info("Backup", "Scheduling pg-backup starter", logger.Params{"target": viper.GetString("target"), "crontime": viper.GetString("crontime")})
		if _, err := c.AddFunc(viper.GetString("crontime"), func() {
			if err := e.Export(); err != nil {
				logger.Error("Backup", err, logger.Params{})
				os.Exit(1)
			}

			logger.Info("Backup", "pg-backup sync ended, waiting for next schedule", logger.Params{"target": viper.GetString("target"), "crontime": viper.GetString("crontime")})
		}); err != nil {
			logger.Error("Backup", err, logger.Params{})
			os.Exit(-1)
		}

		c.Run()
	},
}
