# Summary

[![Build Status](https://travis-ci.org/mono83/optional.svg?branch=master)](https://travis-ci.org/mono83/optional)
[![](https://godoc.org/github.com/mono83/optional?status.svg)](http://godoc.org/github.com/mono83/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/mono83/optional)](https://goreportcard.com/report/github.com/mono83/optional)
![](https://img.shields.io/badge/Go-1.7%2B-green.svg)
![](https://img.shields.io/github/tag/mono83/optional.svg?style=flat)


This library has two major goals:

1. Provide data types to be used during incoming and outgoing data mapping for optional values
2. Provide safe and sufficient mapping functions to work with optional values

# TL&DR;

Simple declare `optional.` data type in your struct

```go
type Configuration struct {
    Hostname optional.String `json:"hostname"`
    Port     optional.Int    `json:"port"`
}

func main() {
    bts, err := ioutil.ReadFile("config.json")
    if err != nil {
        panic(err)
    }
    var c Configuration
    if err := json.Unmarshal(bts, &v); err != nil {
        panic(err)
    }

    host := c.Hostname.OrElse("localhost")
    port := c.Port.OrElse(3306)
}
```

# Features

### Supported mapping:
* JSON
* SQL
* YAML compatible with https://github.com/go-yaml/yaml
* TOML decoding only, compatible with https://github.com/BurntSushi/toml

### Data types, all of them are immutable:
* `optional.Bool` - for booleans
* `optional.Int` - for integers
* `optional.String` - for strings
* `optional.Float64` - for floats
* `optional.Time` - for `time.Time`, mapping using `optional.TimeUnixSeconds`
* `optional.Duration` - for `time.Duration`, mapping using `optional.DurationSeconds`,
  `optional.DurationMillis` or `optional.DurationMinutes`

### Common methods for all data types:
* `Get` - returns value from optional, but panics if optional is empty
* `OrElse(els)` - returns value from optional or provided `els` value if optional is empty
* `IsPresent` - returns `true` if optional contains vaue, `false` otherwise
* `IfPresent(func)` - invokes provided `func` over value from optional if optional not empty
* `Filter(predicate)` - applies predicate over optional value
* `MapToString(func)`, `MapToInt(func)`, `MapToFloat(func)`, `MapToBool(func)` - maps optional value to other type using `func`
