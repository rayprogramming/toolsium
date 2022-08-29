/*
Copyright © 2022 James Ray james@rayprogramming.com

This is the Toolsium config pkg.

This is used to handle all configuration needs that toolsium might need.
*/
package config

import "github.com/spf13/viper"

type Toolsium struct {
	Config         *viper.Viper
	ConfigDir      string
	ConfigFileName string
}

var t *Toolsium

func init() {
	t = New()
}

// Create a new Configuration object
func New() (t *Toolsium) {
	t = new(Toolsium)
	t.Config = viper.New()
	t.ConfigDir = DefaultConfigDirPath()
	t.ConfigFileName = DefaultConfigFileName()
	return
}

func GetViper() *viper.Viper { return t.GetViper() }
func (t *Toolsium) GetViper() *viper.Viper {
	return t.Config
}
