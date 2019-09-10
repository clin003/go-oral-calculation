package gen

// IFactor interface.
type IFactor interface {
	// Result returns the factor final result.
	Result() int

	// Format the factor fomula with FormatType.
	Format(t FormatType) string
}
