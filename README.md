# 💕 config-loader
> *Minimal and safe way to load in configuration files without any extra boilerplate, made for my own personal usage!*

## Why did you build this?
Because, I didn't want to repeat and copy/paste code into the Go projects I am creating, so I made it
into a tiny package that everyone can use, but you don't have to!

## Usage
```sh
$ go get go.floofy.dev/config
```

```go
package main

import (
	"fmt"
	"go.floofy.dev/config"
)

func main() {
	fmt.Println("reading config from ./config.yml!")
	loader := config.NewLoader(config.WithTOML(), config.AutomaticEnv("PREFIX"))
	cfg, err := loader.Load()
	if err != nil {
		panic(err)
    }
	
	val := cfg.Get("owo.da.uwu") // => interface{} (can be `nil`)
	fmt.Println("got owo.da.uwu => %v", val)
}
```

## License
**config-loader** is released under the **Apache 2.0** License.
