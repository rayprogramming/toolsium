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
		log.Debugf("Using custom cfgFile %v", cfgFile)
		// Use config file from the flag.
		t.Config.SetConfigFile(cfgFile)
		t.SetConfigFileName(filepath.Base(cfgFile))
		t.SetConfigDir(filepath.Dir(cfgFile))
	}
	t.configureViper()
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
		path = t.DefaultConfigDirPath()
	}
	return
}

// Set the Current Config directory
//
// If it is unable to find the supplied directory, it will use the default.
func SetConfigDir(confDir string) { t.SetConfigDir(confDir) }
func (t *Toolsium) SetConfigDir(confDir string) {
	if _, osErr := os.Stat(confDir); os.IsNotExist(osErr) {
		log.Debugf("Provided config directory didn't exist. %v", confDir)
		t.ConfigDir = t.DefaultConfigDirPath()
	}
	t.ConfigDir = confDir
}

// Set the Current Config File Name
//
// If it is unable to find the supplied file name, it will use the default.
func SetConfigFileName(fileName string) { t.SetConfigFileName(fileName) }
func (t *Toolsium) SetConfigFileName(fileName string) {
	if fileName == "" {
		t.ConfigFileName = t.DefaultConfigFileName()
	}
	t.ConfigFileName = fileName
}

// Returns the current Config Name
func GetConfigFileName() string { return t.GetConfigFileName() }
func (t *Toolsium) GetConfigFileName() (fileName string) {
	if t.ConfigFileName != "" {
		fileName = t.ConfigFileName
	} else {
		fileName = t.DefaultConfigFileName()
	}
	return
}

// Returns the current Config Directory
func GetConfigFilePath() string { return t.GetConfigFilePath() }
func (t *Toolsium) GetConfigFilePath() string {
	return filepath.Join(t.GetConfigDir(), t.GetConfigFileName())
}

// Configures viper based on the provided passed directory and uses default file type and name.
func (t *Toolsium) configureViper() {
	t.Config.AddConfigPath(t.GetConfigDir())
	t.Config.SetConfigType(cfgFileType)
	t.Config.SetConfigName(t.GetConfigFileName())
	t.Config.SetEnvPrefix(cfgPrefix)
	t.Config.AutomaticEnv()
}

// Creates the configuration directory
func (t *Toolsium) createConfigDirectory() {
	path := t.GetConfigDir()
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// Writes the toolsium configuration file
func CreateConfig() error { return t.CreateConfig() }
func (t *Toolsium) CreateConfig() error {
	t.createConfigDirectory()
	cfgFile := t.GetConfigFilePath()
	log.Infof("Wrote Config to %v", cfgFile)
	return t.Config.WriteConfigAs(cfgFile)
}
