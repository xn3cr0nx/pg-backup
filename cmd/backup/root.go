package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xn3cr0nx/pg-backup/pkg/logger"
)

var (
	debug                                       bool
	ct                                          string
	target, host, port, user, pass, db          string
	output, outputPrefix, outputPath, outputExt string
	outputTime                                  bool
	awsAcessKey,
	awsSecretKey,
	awsS3Region,
	awsS3Bucket string
)

var rootCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup pg dump export service",
	Long:  `Backup service to export pg dump to external storage`,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	logger.Setup()

	fmt.Println("System checking...")
	OS := readOS()
	fmt.Println("Running on: " + OS)
	pgDumpInstalled := pgDumpCheck()
	if !pgDumpInstalled {
		fmt.Println("pg_dump is missing. Install postgresql-client.")
	}
	fmt.Println("pg_dump in installed")
	fmt.Println("Running pg-backup")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(startCmd)

	// Adds root flags and persistent flags
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Sets logging level to Debug")
	rootCmd.PersistentFlags().StringVar(&ct, "crontime", "@daily", "Sets crontime - [default: @daily]")

	// pg variables
	rootCmd.PersistentFlags().StringVar(&host, "pg_host", "localhost", "Sets postgres host")
	rootCmd.PersistentFlags().StringVar(&port, "pg_port", "5432", "Sets postgres port")
	rootCmd.PersistentFlags().StringVar(&user, "pg_user", "postgres", "Sets postgres user")
	rootCmd.PersistentFlags().StringVar(&pass, "pg_pass", "", "Sets postgres password")
	rootCmd.PersistentFlags().StringVar(&db, "pg_db", "postgres", "Sets postgres db")

	rootCmd.PersistentFlags().StringVar(&target, "target", "", "Sets export target between s3, file - [default: s3]")

	rootCmd.PersistentFlags().StringVar(&outputPath, "output_path", "./backups/", "Adds directory path to output name.")
	rootCmd.PersistentFlags().StringVar(&outputPrefix, "output_prefix", "", "Adds prefix to output name.")
	rootCmd.PersistentFlags().BoolVar(&outputTime, "output_time", true, "Sets if output name should include time")
	rootCmd.PersistentFlags().StringVar(&outputExt, "output_ext", "psql", "Sets output extension - [default: psql]")

	// s3 variables
	rootCmd.PersistentFlags().StringVar(&awsAcessKey, "aws_access_key", "", "Sets AWS access key")
	rootCmd.PersistentFlags().StringVar(&awsSecretKey, "aws_secret_key", "", "Sets AWS secret key")
	rootCmd.PersistentFlags().StringVar(&awsS3Region, "aws_s3_region", "us-east-1", "Sets AWS S3 region")
	rootCmd.PersistentFlags().StringVar(&awsS3Bucket, "aws_s3_bucket", "bucket", "Sets AWS S3 bucket")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault("debug", false)
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	viper.SetDefault("crontime", "@daily")
	viper.BindPFlag("crontime", rootCmd.PersistentFlags().Lookup("crontime"))

	// pg variables
	viper.SetDefault("pg_host", "localhost")
	viper.BindPFlag("pg_host", rootCmd.PersistentFlags().Lookup("pg_host"))
	viper.SetDefault("pg_port", "5432")
	viper.BindPFlag("pg_port", rootCmd.PersistentFlags().Lookup("pg_port"))
	viper.SetDefault("pg_user", "postgres")
	viper.BindPFlag("pg_user", rootCmd.PersistentFlags().Lookup("pg_user"))
	viper.SetDefault("pg_pass", "")
	viper.BindPFlag("pg_pass", rootCmd.PersistentFlags().Lookup("pg_pass"))
	viper.SetDefault("pg_db", "postgres")
	viper.BindPFlag("pg_db", rootCmd.PersistentFlags().Lookup("pg_db"))

	viper.SetDefault("target", "file")
	viper.BindPFlag("target", rootCmd.PersistentFlags().Lookup("target"))

	viper.SetDefault("output_time", true)
	viper.BindPFlag("output_time", rootCmd.PersistentFlags().Lookup("output_time"))
	viper.SetDefault("output_prefix", "")
	viper.BindPFlag("output_prefix", rootCmd.PersistentFlags().Lookup("output_prefix"))
	viper.SetDefault("output_path", "./backups/")
	viper.BindPFlag("output_path", rootCmd.PersistentFlags().Lookup("output_path"))
	viper.SetDefault("output_ext", "psql")
	viper.BindPFlag("output_ext", rootCmd.PersistentFlags().Lookup("output_ext"))

	// s3 variables
	viper.SetDefault("aws_access_key", "")
	viper.BindPFlag("aws_access_key", rootCmd.PersistentFlags().Lookup("aws_access_key"))
	viper.SetDefault("aws_secret_key", "")
	viper.BindPFlag("aws_secret_key", rootCmd.PersistentFlags().Lookup("aws_secret_key"))
	viper.SetDefault("aws_s3_region", "us-east-1")
	viper.BindPFlag("aws_s3_region", rootCmd.PersistentFlags().Lookup("aws_s3_region"))
	viper.SetDefault("aws_s3_region", "bucket")
	viper.BindPFlag("aws_s3_region", rootCmd.PersistentFlags().Lookup("aws_s3_region"))

	viper.AutomaticEnv()

	if value, ok := os.LookupEnv("CONFIG_FILE"); ok {
		viper.SetConfigFile(value)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/pg-backup/")
		viper.AddConfigPath("$HOME/.pg-backup/backup")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
	}

	viper.ReadInConfig()
	f := viper.ConfigFileUsed()
	if f != "" {
		fmt.Printf("Found configuration file: %s \n", f)
	}
}

func readOS() string {
	out, err := exec.Command("uname").Output()
	if err != nil {
		logger.Error("Backup", err, logger.Params{})
		os.Exit(1)
	}
	OS := strings.Trim(string(out), " \n")
	if OS != "Linux" {
		logger.Error("Backup", fmt.Errorf("%s unsopported", OS), logger.Params{})
		os.Exit(1)
	}
	return OS
}

func pgDumpCheck() bool {
	_, err := os.Stat("/usr/bin/pg_dump")
	return err == nil
}
