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

func DefaultConfigDir() string { return t.DefaultConfigDir() }
func (t *Toolsium) DefaultConfigDir() string {
	return cfgFolderName
}

// Returns the default Config Dir path
func DefaultConfigDirPath() string { return t.DefaultConfigDirPath() }
func (t *Toolsium) DefaultConfigDirPath() (path string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(home, t.DefaultConfigDir())
	return
}

// Returns the default config file path
func DefaultConfigFileName() string { return t.DefaultConfigFileName() }
func (t *Toolsium) DefaultConfigFileName() string {
	return cfgFileName
}

// Returns the full default config path
func DefaultConfigPath() string { return t.DefaultConfigPath() }
func (t *Toolsium) DefaultConfigPath() string {
	return filepath.Join(t.DefaultConfigDirPath(), t.DefaultConfigFileName())
}
