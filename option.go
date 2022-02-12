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

// OptionOverloadFunc is a type to represent an overriding, overloading
// option.
type OptionOverloadFunc func(o Options)

// Options represents the configuration options to use
// when configuring the config.Loader!
type Options struct {
	format       Format
	automaticEnv bool
	envPrefix    *string
}

// NewDefaultOptions creates a new Options object with the default settings.
func NewDefaultOptions() Options {
	return Options{
		format:       JSONFormat,
		automaticEnv: false,
		envPrefix:    nil,
	}
}

// WithToml is a OptionOverloadFunc to override the Options.format property to be the
// TOMLFormat.
func WithToml() OptionOverloadFunc {
	return func(o Options) {
		o.format = TOMLFormat
	}
}

// WithYAML is a OptionOverloadFunc to override the Options.format property to be
// YAMLFormat.
func WithYAML() OptionOverloadFunc {
	return func(o Options) {
		o.format = YAMLFormat
	}
}

// AutomaticEnv sets up using Environment Variables as a secondary option
// if no configuration path was found OR it was found, but can be overwritten.
func AutomaticEnv(prefix string) OptionOverloadFunc {
	return func(o Options) {
		o.automaticEnv = true
		o.envPrefix = &prefix
	}
}
