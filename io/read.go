package io

import (
	"github.com/jsnctl/pirsig/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Script struct {
	Sequence model.Sequence `yaml:"sequence"`
}

func Read() *Script {
	script := Script{}
	file, _ := ioutil.ReadFile("test.yaml")
	yaml.Unmarshal(file, &script)
	return &script
}
