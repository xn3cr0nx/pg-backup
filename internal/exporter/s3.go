package exporter

import (
	"os"

	"github.com/spf13/viper"
	"github.com/xn3cr0nx/pg-backup/internal/naming"
	"github.com/xn3cr0nx/pg-backup/internal/pg"
	"github.com/xn3cr0nx/pg-backup/pkg/logger"
	"github.com/xn3cr0nx/pg-backup/pkg/s3"
)

// S3Exporter implement export method for export dump to s3 bucket
type S3Exporter struct{}

// NewS3Exporter returns a new instance of S3Exporter
func NewS3Exporter() *S3Exporter {
	return &S3Exporter{}
}

// Export runs pg_dump and export the dump to selected target
func (e *S3Exporter) Export() (err error) {
	logger.Debug("Backup", "Pg dump running...", logger.Params{})

	dump, err := pg.Dump()
	if err != nil {
		return
	}

	logger.Debug("Backup", "Pg correctly dumped", logger.Params{})

	logger.Debug("Backup", "Exporting to file", logger.Params{"name": naming.Output()})

	f, err := os.Create(naming.Output())
	defer f.Close()
	if err != nil {
		return
	}
	if _, err = f.WriteString(dump); err != nil {
		return
	}

	switch viper.GetString("target") {
	case "s3":
		logger.Debug("Backup", "Exporting to S3 target", logger.Params{"bucket": viper.GetString("aws_s3_bucket"), "region": viper.GetString("aws_s3_region")})
		session, e := s3.NewS3Session(viper.GetString("aws_s3_region"))
		if e != nil {
			return e
		}
		if err = s3.UploadFile(session, viper.GetString("aws_s3_bucket"), viper.GetString("output"), f); err != nil {
			return
		}
	}

	return
}
