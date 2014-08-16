package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"strings"
)

type ImageLoadConfig struct {
	Keyword      string        `json:"keyword"`
 	OutputDir    string        `json:"outputDir"`
	PageSettings []PageSetting `json:"pageSettings"`
}

const (
	configDirName  = ".web-scraping-example"
	configFileName = "config.json"
)

func createConfigPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	configPath := strings.Join([]string{u.HomeDir, configDirName, configFileName}, "/")
	return configPath, nil
}

func ReadConfigFile() (*ImageLoadConfig, error) {
	configPath, err := createConfigPath()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var c ImageLoadConfig
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
