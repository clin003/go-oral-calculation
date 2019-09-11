package gen

// IFactor interface.
type IFactor interface {
	// Result returns the factor final result.
	Result() int

	// Format the factor fomula with FormatType and if needs to format with result.
	Format(t FormatType, withResult bool) string
}
