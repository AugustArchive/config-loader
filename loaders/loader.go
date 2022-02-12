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

// FindConfigOptionFunc is the predicate function for Loader.Find
type FindConfigOptionFunc func(key string) bool

// Loader is the internal loader that config.Loader uses
// to override the options.
type Loader interface {
	// Get returns the property that was found as an `interface{}` or nil,
	// if nothing was found. This supports dot notation, so you can do:
	//
	// key.inner.inner
	//
	// it will traverse through the object if we can find it!
	Get(key string) interface{}

	// Set overrides a key-value pair and tries to write it into the disk.
	// Returns an error if anything goes wrong.
	Set(key string, value interface{}) error

	// Has checks if the `key` is in the config tree, if not, it will
	// return `false`, otherwise, it will return `true`.
	Has(key string) bool

	// Format returns the config.Format it is using
	Format() config.Format

	// Load initializes the config tree and returns a map[string]interface{}
	// and an `error` if anything has occurred.
	Load() error

	// Raw returns the raw data, just in case you need it.
	Raw() map[string]interface{}
}
