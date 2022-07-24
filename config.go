package gilc

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/alexcoder04/friendly"
)

// utils {{{
func getConfigFile() string {
	return path.Join(Plugin.Data.SavePath, "var", Plugin.Name, "config.json")
}

func createConfigFolder() error {
	configFolder := path.Dir(getConfigFile())
	if !friendly.IsDir(configFolder) {
		return os.MkdirAll(configFolder, 0700)
	}
	return nil
}

func readConfig() (map[string]string, error) {
	configFile := getConfigFile()
	if !friendly.IsFile(configFile) {
		return map[string]string{}, nil
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return map[string]string{}, err
	}

	var config map[string]string
	err = json.Unmarshal(data, &config)
	return config, err
}

func writeConfig(config map[string]string) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(getConfigFile(), data, 0600)
}

// }}}

func ConfigGet(key string) (string, error) {
	err := createConfigFolder()
	if err != nil {
		return "", err
	}

	config, err := readConfig()
	if err != nil {
		return "", err
	}

	value, ok := config[key]
	if !ok {
		return "", nil
	}
	return value, nil
}

func ConfigSet(key string, value string) error {
	err := createConfigFolder()
	if err != nil {
		return err
	}

	config, err := readConfig()
	if err != nil {
		return err
	}

	config[key] = value
	return writeConfig(config)
}
