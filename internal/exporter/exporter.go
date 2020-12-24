package exporter

// Exporter interface implement methods for export targets
type Exporter interface {
	Export() error
}
