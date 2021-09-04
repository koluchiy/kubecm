package kubeconf

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func saveConfig(config KubeConfig, path string) error {
	js, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, js, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
