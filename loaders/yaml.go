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

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"go.floofy.dev/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

// YamlLoader configures an internal Loader to use .toml files
// to load into the configuration.
type YamlLoader struct {
	path string
	data map[string]interface{}
}

func NewYAMLLoader(path string) Loader {
	return YamlLoader{
		path: path,
		data: make(map[string]interface{}),
	}
}

func (y YamlLoader) Get(key string) interface{} {
	nodes := strings.Split(key, ".")
	var value interface{}

	found := false

	for _, node := range nodes {
		// Check if the key is a `map[string]interface{}`
		map_, ok := y.data[node].(map[string]interface{})
		if !ok {
			found = true
			value = map_
			break
		}
	}

	if !found {
		return nil
	}

	return value
}

func (y YamlLoader) Set(key string, value interface{}) error {
	nodes := strings.Split(key, ".")
	found := false

	var struct_ map[string]interface{}
	for _, node := range nodes {
		map_, ok := y.data[node].(map[string]interface{})
		if ok {
			found = true
			struct_ = map_
			break
		}
	}

	if !found {
		return fmt.Errorf("unable to find key '%s' in config tree", key)
	}

	struct_[key] = value
	file, _ := os.OpenFile(y.path, os.O_CREATE, os.ModePerm)
	defer func() {
		_ = file.Close()
	}()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(y.data); err != nil {
		return err
	}

	return nil
}

func (y YamlLoader) Has(key string) bool {
	nodes := strings.Split(key, ".")
	found := false

	for _, node := range nodes {
		_, ok := y.data[node].(map[string]interface{})
		if ok {
			found = true
			break
		}
	}

	return found
}

func (_ YamlLoader) Format() config.Format {
	return config.YAMLFormat
}

func (y YamlLoader) Raw() map[string]interface{} {
	return y.data
}

func (y YamlLoader) Load() error {
	// Check if the file exists
	_, err := os.Stat(y.path)
	if err != nil {
		return err
	}

	// Let's read the contents of the file
	contents, err := ioutil.ReadFile(y.path)
	if err != nil {
		return nil
	}

	// Let's decode the contents of the file
	var data map[string]interface{}
	if err := yaml.Unmarshal(contents, &data); err != nil {
		return err
	}

	// Before we do, let's traverse through the tree,
	// since we can do ${ENV_VARIABLE} to override from the
	// system environment variable to this value.
	for key, value := range data {
		// Check if `value` is a string
		val, ok := value.(string)
		if ok {
			if v, ok := os.LookupEnv(val); ok {
				data[key] = v
			}
		}
	}

	y.data = data
	return nil
}
