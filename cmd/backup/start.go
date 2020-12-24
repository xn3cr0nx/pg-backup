package main

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/xn3cr0nx/pg-backup/pkg/logger"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Export dump cron",
	Long:  `Dump pg periodically using pg_dump command`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Spider", "Spider starting...", logger.Params{})

		// pg, err := postgres.NewPg(postgres.Conf())
		// if err != nil {
		// 	logger.Error("Spider", err, logger.Params{})
		// 	os.Exit(-1)
		// }
		// if err := pg.Connect(); err != nil {
		// 	logger.Error("Spider", err, logger.Params{})
		// 	os.Exit(-1)
		// }
		// if err := migration.Migration(pg); err != nil {
		// 	logger.Error("Spider", err, logger.Params{})
		// 	os.Exit(-1)
		// }

		// target := viper.GetString("target")
		// if target != "" {
		// 	if target == "bitcoinabuse" {
		// 		btcabuse := bitcoinabuse.NewSpider(pg)
		// 		if err := btcabuse.Sync(); err != nil {
		// 			logger.Error("Spider", err, logger.Params{})
		// 			os.Exit(-1)
		// 		}
		// 		logger.Info("Spider", "bitcoinabuse sync ended, waiting for next schedule", logger.Params{"target": "bitcoinabuse.com"})
		// 		return
		// 	}

		// 	if target == "checkbitcoinaddress" {
		// 		checkbtcaddr := checkbitcoinaddress.NewSpider(pg)
		// 		if err := checkbtcaddr.Sync(); err != nil {
		// 			logger.Error("Spider", err, logger.Params{})
		// 			os.Exit(-1)
		// 		}
		// 		logger.Info("Spider", "checkbitcoinaddress sync ended, waiting for next schedule", logger.Params{"target": "bitcoinabuse.com"})
		// 		return
		// 	}

		// 	if target == "walletexplorer" {
		// 		wexplorer := walletexplorer.NewSpider(pg)
		// 		if err := wexplorer.Sync(); err != nil {
		// 			logger.Error("Spider", err, logger.Params{})
		// 			os.Exit(-1)
		// 		}
		// 		logger.Info("Spider", "walletexplorer sync ended, waiting for next schedule", logger.Params{"target": "walletexplorer.com"})
		// 		return
		// 	}
		// }

		// if !viper.GetBool("cron") {
		// 	btcabuse := bitcoinabuse.NewSpider(pg)
		// 	if err := btcabuse.Sync(); err != nil {
		// 		logger.Error("Spider", err, logger.Params{})
		// 		os.Exit(-1)
		// 	}
		// 	logger.Info("Spider", "bitcoinabuse sync ended, waiting for next schedule", logger.Params{"target": "bitcoinabuse.com", "crontime": viper.GetString("spider.crontime")})

		// 	checkbtcaddr := checkbitcoinaddress.NewSpider(pg)
		// 	if err := checkbtcaddr.Sync(); err != nil {
		// 		logger.Error("Spider", err, logger.Params{})
		// 		os.Exit(-1)
		// 	}
		// 	logger.Info("Spider", "checkbitcoinaddress sync ended, waiting for next schedule", logger.Params{"target": "bitcoinabuse.com", "crontime": viper.GetString("spider.crontime")})

		// 	wexplorer := walletexplorer.NewSpider(pg)
		// 	if err := wexplorer.Sync(); err != nil {
		// 		logger.Error("Spider", err, logger.Params{})
		// 		os.Exit(-1)
		// 	}
		// 	logger.Info("Spider", "walletexplorer sync ended, waiting for next schedule", logger.Params{"target": "walletexplorer.com", "crontime": viper.GetString("spider.crontime")})
		// 	return
		// }

		c := cron.New()
		defer c.Stop()

		// logger.Info("Spider", "Scheduling spider starter", logger.Params{"target": "bitcoinabuse.com", "crontime": viper.GetString("spider.crontime")})
		// if _, err = c.AddFunc(viper.GetString("spider.crontime"), func() {
		// 	btcabuse := bitcoinabuse.NewSpider(pg)
		// 	if err := btcabuse.Sync(); err != nil {
		// 		logger.Error("Spider", err, logger.Params{})
		// 		os.Exit(-1)
		// 	}
		// 	logger.Info("Spider", "bitcoinabuse sync ended, waiting for next schedule", logger.Params{"target": "bitcoinabuse.com", "crontime": viper.GetString("spider.crontime")})
		// }); err != nil {
		// 	logger.Error("Spider", err, logger.Params{})
		// 	os.Exit(-1)
		// }

		// logger.Info("Spider", "Scheduling spider starter", logger.Params{"target": "checkbitcoinaddress.com", "crontime": viper.GetString("spider.crontime")})
		// if _, err = c.AddFunc(viper.GetString("spider.crontime"), func() {
		// 	checkbtcaddr := checkbitcoinaddress.NewSpider(pg)
		// 	if err := checkbtcaddr.Sync(); err != nil {
		// 		logger.Error("Spider", err, logger.Params{})
		// 		os.Exit(-1)
		// 	}
		// 	logger.Info("Spider", "checkbitcoinaddress sync ended, waiting for next schedule", logger.Params{"target": "bitcoinabuse.com", "crontime": viper.GetString("spider.crontime")})
		// }); err != nil {
		// 	logger.Error("Spider", err, logger.Params{})
		// 	os.Exit(-1)
		// }

		// logger.Info("Spider", "Scheduling spider starter", logger.Params{"target": "walletexplorer.com", "crontime": viper.GetString("spider.crontime")})
		// if _, err = c.AddFunc(viper.GetString("spider.crontime"), func() {
		// 	wexplorer := walletexplorer.NewSpider(pg)
		// 	if err := wexplorer.Sync(); err != nil {
		// 		logger.Error("Spider", err, logger.Params{})
		// 		os.Exit(-1)
		// 	}
		// 	logger.Info("Spider", "walletexplorer sync ended, waiting for next schedule", logger.Params{"target": "walletexplorer.com", "crontime": viper.GetString("spider.crontime")})
		// }); err != nil {
		// 	logger.Error("Spider", err, logger.Params{})
		// 	os.Exit(-1)
		// }

		c.Run()
	},
}
