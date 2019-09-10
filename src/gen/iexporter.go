package gen

// IExporter interface.
type IExporter interface {
	// Export oral calculation data with FormatType and how many columns.
	Export(t FormatType, cols int)
}
