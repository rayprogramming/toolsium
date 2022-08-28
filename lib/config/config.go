/*
Copyright © 2022 James Ray james@rayprogramming.com

This is the Toolsium config pkg.

This is used to handle all configuration needs that toolsium might need.
*/
package config

import (
	"errors"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Configures Toolsium based on the cfgFile path passed in.
func Configure(cfgFile string) { t.Configure(cfgFile) }
func (t *Toolsium) Configure(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		t.Config.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".toolsium" (without extension).
		t.ConfigureViper(t.GetConfigDir())
	}
	// If a config file is found, read it in.
	if err := t.Config.ReadInConfig(); err == nil {
		log.Infof("Using config file: %v", t.Config.ConfigFileUsed())
	}
}

// Returns the current Config Directory
func GetConfigDir() string { return t.GetConfigDir() }
func (t *Toolsium) GetConfigDir() (path string) {
	if t.ConfigDir != "" {
		path = t.ConfigDir
	} else {
		path = t.GetDefaultConfigDir()
	}
	return
}

// Set the Current Config directory
//
// If it is unable to find the supplied directory, it will use the default.
func SetConfigDir(confDir string) error { return t.SetConfigDir(confDir) }
func (t *Toolsium) SetConfigDir(confDir string) (err error) {
	if _, osErr := os.Stat(confDir); os.IsNotExist(osErr) {
		log.Infof("Provided config directory didn't exist. %v", confDir)
		t.ConfigDir = t.GetDefaultConfigDir()
	}
	t.ConfigDir = confDir
	return
}

// Returns the current Config Directory
//
// TODO(JR): Change usage of cfgFileName to allow changes to name
func GetConfigFile() string { return t.GetConfigFile() }
func (t *Toolsium) GetConfigFile() string {
	return filepath.Join(t.GetConfigDir(), cfgFileName)
}

// TODO(JR): Add SetConfigFile
// TODO(JR): Add GetConfigPath to fetch full path and reduce GetConfigFile

// Configures viper based on the provided passed directory and uses default file type and name.
func ConfigureViper(installDir string) { t.ConfigureViper(installDir) }
func (t *Toolsium) ConfigureViper(installDir string) {
	t.Config.AddConfigPath(installDir)
	t.Config.SetConfigType(cfgFileType)
	t.Config.SetConfigName(cfgFileName)
	t.Config.SetEnvPrefix(cfgPrefix)
	t.Config.AutomaticEnv()
}

func (t *Toolsium) createConfigDirectory() {
	path := t.GetConfigDir()
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Errorln(err)
		}
	}
}

// Writes the toolsium configuration file
func CreateConfig() error { return t.CreateConfig() }
func (t *Toolsium) CreateConfig() (err error) {
	t.createConfigDirectory()
	cfgFile := t.GetConfigFile()
	log.Infof("Wrote Config to %v", cfgFile)
	err = t.Config.WriteConfigAs(cfgFile)
	return
}
