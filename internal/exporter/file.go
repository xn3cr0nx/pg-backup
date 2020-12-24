package exporter

import (
	"os"

	"github.com/xn3cr0nx/pg-backup/internal/naming"
	"github.com/xn3cr0nx/pg-backup/internal/pg"
	"github.com/xn3cr0nx/pg-backup/pkg/logger"
)

// FileExporter implement export method for export dump to file
type FileExporter struct{}

// NewFileExporter returns a new instance of FileExporter
func NewFileExporter() *FileExporter {
	return &FileExporter{}
}

// Export runs pg_dump and export the dump to selected target
func (e *FileExporter) Export() (err error) {
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

	return
}
