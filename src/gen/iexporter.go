package gen

// IExporter interface.
type IExporter interface {
	// Export oral calculation data with FormatType, how many columns and whether
	// needs to export item index.
	Export(t FormatType, cols int, indexed bool)
}
