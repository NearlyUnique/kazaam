package transform

// Pass performs no manipulation of the passed-in data. It is useful
// for testing/default behavior.
func PassRaw(spec *Config, data []byte) ([]byte, error) {
	return data, nil
}
