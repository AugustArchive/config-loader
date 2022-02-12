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

package loaders

import "go.floofy.dev/config"

// NoopLoader represents a Loader struct that does no operations at all.
type NoopLoader struct{}

// NewNoopLoader creates a new NoopLoader object.
func NewNoopLoader() Loader {
	return NoopLoader{}
}

func (_ NoopLoader) Get(_ string) interface{} {
	return nil
}

func (_ NoopLoader) Set(_ string, _ interface{}) error {
	return nil
}

func (_ NoopLoader) Has(key string) bool {
	return false
}

func (_ NoopLoader) Format() config.Format {
	return config.NoopFormat
}

func (_ NoopLoader) Load() error {
	return nil
}
