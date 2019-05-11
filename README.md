goconf
---
[![GoDoc](https://godoc.org/plimble/goconf?status.svg)](https://godoc.org/github.com/plimble/goconf)
[![Build Status](https://travis-ci.org/plimble/goconf.svg?branch=master)](https://travis-ci.org/plimble/goconf?branch=master)
[![Coverage Status](https://coveralls.io/repos/plimble/goconf/badge.svg?branch=master&service=github&foo)](https://coveralls.io/github/plimble/goconf?branch=master)
[![Go Report Card](https://goreportcard.com/badge/plimble/goconf)](https://goreportcard.com/report/plimble/goconf)

Combine yaml and environment config

## Features
- [Parse yaml](gopkg.in/yaml.v2)
- [Parse env](github.com/kelseyhightower/envconfig)
- Watch yaml config file

## Installation

```
go get github.com/onedaycat/goconf
```

## Struct Tags

```go
type SampleA struct {
	A               string
	CamelCase       bool 
	ManualOverride1 string 
	SplitWord1      string `default:"split"`
	ID              string 
	DefaultValue    string 
}
```

## Example

```go
type Sample struct {
  Value string
}


sample := &Sample{}
goconf.Parse("PREFIX", sample)
```
