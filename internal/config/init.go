package config

import (
	"github.com/json-iterator/go"
	"os"
)

var env = os.Getenv
var json = jsoniter.ConfigCompatibleWithStandardLibrary
