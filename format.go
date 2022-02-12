// 💕 config-loader: Minimal and safe way to load in configuration files without any extra boilerplate,
// made for my own personal usage!
// Copyright 2022 Noel <cutie@floofy.dev>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	// NoopFormat is the non-operational format that is used for testing.
	NoopFormat Format = "no-operation"
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

	case NoopFormat:
		return "no-operation"

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

	case NoopFormat:
		return []string{}

	default:
		return []string{}
	}
}
