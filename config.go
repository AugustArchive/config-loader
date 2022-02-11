// Package config is Noel's configuration loader for Go. It supports the following formats:
//
//   - JSON
//   - YAML
//   - TOML
//   - Environment Variables
//
// You can simply call `config.NewLoader()` to create a new Loader object.
package config

// Loader is the main configuration loader
type Loader struct {
	Format Format
}

// NewLoader creates a new Loader object.
func NewLoader() Loader {
	return Loader{}
}
