package config

// Format is a type that determines a configuration format that the config.Loader
// is using.
type Format string

var (
	// JSONFormat is the configuration format to use the encoding/json
	// package to load and lock in your configuration object.
	JSONFormat Format = "JSON"

	// YAMLFormat is the configuration format to use YAML/YML files
	// to load and lock in your configuration object.
	YAMLFormat Format = "YAML"

	// TOMLFormat is the configuration format to use TOML files
	// to load and lock in your configuration object.
	TOMLFormat Format = "TOML"

	// EnvFormat is the configuration format to use environment variables
	// to load and lock in your configuration object. This can be applied
	// using the config.AutomaticEnv option
	EnvFormat Format = "Environment"
)

// String stringifies a Format variable.
func (f Format) String() string {
	switch f {
	case JSONFormat:
		return "JSON"

	case YAMLFormat:
		return "YAML"

	case TOMLFormat:
		return "TOML"

	case EnvFormat:
		return "Environment"

	default:
		return "Unknown"
	}
}

// Extensions returns a list of strings of the given extensions
// of this Format object.
func (f Format) Extensions() []string {
	switch f {
	case JSONFormat:
		return []string{".json"}

	case YAMLFormat:
		return []string{".yaml", ".yml"}

	case TOMLFormat:
		return []string{".toml"}

	case EnvFormat:
		return []string{".env"}

	default:
		return []string{}
	}
}
