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
	"encoding/json"
	"fmt"
	"go.floofy.dev/config"
	"io/ioutil"
	"os"
	"strings"
)

// JsonLoader configures an internal Loader to use .json files
// to load into the configuration.
type JsonLoader struct {
	path string
	data map[string]interface{}
}

// NewJsonLoader creates a new JsonLoader object.
func NewJsonLoader(path string) Loader {
	return JsonLoader{
		path: path,
		data: make(map[string]interface{}),
	}
}

func (j JsonLoader) Get(key string) interface{} {
	nodes := strings.Split(key, ".")
	var value interface{}

	found := false

	for _, node := range nodes {
		// Check if the key is a `map[string]interface{}`
		map_, ok := j.data[node].(map[string]interface{})
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

func (j JsonLoader) Set(key string, value interface{}) error {
	nodes := strings.Split(key, ".")
	found := false

	var struct_ map[string]interface{}
	for _, node := range nodes {
		map_, ok := j.data[node].(map[string]interface{})
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
	file, _ := os.OpenFile(j.path, os.O_CREATE, os.ModePerm)
	defer func() {
		_ = file.Close()
	}()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(j.data); err != nil {
		return err
	}

	return nil
}

func (j JsonLoader) Has(key string) bool {
	nodes := strings.Split(key, ".")
	found := false

	for _, node := range nodes {
		_, ok := j.data[node].(map[string]interface{})
		if ok {
			found = true
			break
		}
	}

	return found
}

func (_ JsonLoader) Format() config.Format {
	return config.JSONFormat
}

func (j JsonLoader) Load() error {
	// Check if the file exists
	_, err := os.Stat(j.path)
	if err != nil {
		return err
	}

	// Let's read the contents of the file
	contents, err := ioutil.ReadFile(j.path)
	if err != nil {
		return nil
	}

	// Let's decode the contents of the file
	var data map[string]interface{}
	if err := json.Unmarshal(contents, &data); err != nil {
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

	j.data = data
	return nil
}

func (j JsonLoader) Raw() map[string]interface{} {
	return j.data
}
