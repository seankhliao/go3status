package main

import (
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/seankhliao/go3status/mod"
)

// ModuleNames is a global map of names to their initializer funcs
var ModuleNames = map[string]func() mod.Module{
	"battery": mod.DefaultBattery,
	"static":  mod.DefaultStatic,
	"time":    mod.DefaultTime,
	"volume":  mod.DefaultVolume,
	"wifi":    mod.DefaultWifi,
}

// ParseConfig parses a string of TOML to generate the array of modules
func ParseConfig(config string) ([]mod.Module, error) {
	var blocks []mod.Module
	var base map[string][]map[string]toml.Primitive
	meta, err := toml.Decode(config, &base)
	if err != nil {
		return blocks, err
	}

	for _, raw := range base["conf"] {
		var key string
		for k := range raw {
			key = k
		}

		newBlockFunc, ok := ModuleNames[key]
		if !ok {
			return blocks, errors.New("module not found: " + key)
		}

		block := newBlockFunc()
		if err = meta.PrimitiveDecode(raw[key], block); err != nil {
			return blocks, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}
