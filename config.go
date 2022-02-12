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

// Package config is Noel's configuration loader for Go. It supports the following formats:
//
//   - JSON
//   - YAML
//   - TOML
//   - Environment Variables
//
// You can simply call `config.NewLoader()` to create a new Loader object.
package config

import (
	"fmt"
	"go.floofy.dev/config/loaders"
)

// Loader is the main configuration loader
type Loader struct {
	options        Options
	internalLoader loaders.Loader
}

// NewLoader creates a new Loader object.
func NewLoader(path string, options ...OptionOverloadFunc) Loader {
	opts := NewDefaultOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Try to find the internal loader to use
	var loader loaders.Loader
	switch opts.format {
	case JSONFormat:
		loader = loaders.NewJsonLoader(path)
		break

	case YAMLFormat:
		loader = loaders.NewYAMLLoader(path)
		break

	case TOMLFormat:
		loader = loaders.NewTomlLoader(path)
		break

	default:
		panic(fmt.Errorf("unable to find loader for %s", opts.format.String()))
	}

	return Loader{
		options:        opts,
		internalLoader: loader,
	}
}

func (l Loader) Get(key string) interface{} {
	return l.internalLoader.Get(key)
}

func (l Loader) Set(key string, value interface{}) error {
	return l.internalLoader.Set(key, value)
}

func (l Loader) Has(key string) bool {
	return l.internalLoader.Has(key)
}

func (l Loader) Raw() map[string]interface{} {
	return l.internalLoader.Raw()
}

func (l Loader) Load() error {
	return l.internalLoader.Load()
}
