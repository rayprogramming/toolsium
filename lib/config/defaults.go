package config

import (
	"os"
	"path/filepath"
)

const (
	cfgFolderName = ".toolsium"
	cfgFileName   = "config"
	cfgFileType   = "json"
	cfgPrefix     = "TOOL"
)

// Returns the default Config Dir path or err
func GetDefaultConfigDir() string { return t.GetDefaultConfigDir() }
func (t *Toolsium) GetDefaultConfigDir() (path string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(home, cfgFolderName)
	return
}

// Returns the default config file path or err
func GetDefaultConfigFile() string { return t.GetDefaultConfigFile() }
func (t *Toolsium) GetDefaultConfigFile() (path string) {
	directory := t.GetDefaultConfigDir()
	path = filepath.Join(directory, cfgFileName)
	return
}
