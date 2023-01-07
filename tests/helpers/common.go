package helpers

// GetVariablePointer returns pointer to provided variable.
func GetVariablePointer[T string | int64](str T) *T {
	return &str
}
