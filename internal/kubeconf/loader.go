package kubeconf

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func LoadConfigAsSingle(configPath string) (KubeConfig, error) {
	config, err := loadConfig(configPath)
	if err != nil {
		return KubeConfig{}, err
	}

	config, err = makeConfigSingle(config, configPath)
	if err != nil {
		return KubeConfig{}, err
	}

	return config, nil
}

func loadConfig(configPath string) (KubeConfig, error) {
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		return KubeConfig{}, err
	}

	var config KubeConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return KubeConfig{}, err
	}

	return config, nil
}

func makeConfigSingle(config KubeConfig, configPath string) (KubeConfig, error) {
	for i, u := range config.Users {
		user := u.User

		if len(user.ClientCertificate) > 0 && len(user.ClientCertificateData) == 0 {
			keyPath := resolveKeyPaths(configPath, user.ClientCertificate)

			data, err := ioutil.ReadFile(keyPath)
			if err != nil {
				return KubeConfig{}, err
			}

			encoded := base64.StdEncoding.EncodeToString(data)

			user.ClientCertificateData = encoded
			user.ClientCertificate = ""
		}

		if len(user.ClientKey) > 0 && len(user.ClientKeyData) == 0 {
			keyPath := resolveKeyPaths(configPath, user.ClientKey)

			data, err := ioutil.ReadFile(keyPath)
			if err != nil {
				return KubeConfig{}, err
			}

			encoded := base64.StdEncoding.EncodeToString(data)

			user.ClientKeyData = encoded
			user.ClientKey = ""
		}

		u.User = user
		config.Users[i] = u
	}

	for i, c := range config.Clusters {
		cluster := c.Cluster

		if len(cluster.CertificateAuthority) > 0 && len(cluster.CertificateAuthorityData) == 0 {
			keyPath := resolveKeyPaths(configPath, cluster.CertificateAuthority)

			data, err := ioutil.ReadFile(keyPath)
			if err != nil {
				return KubeConfig{}, err
			}

			encoded := base64.StdEncoding.EncodeToString(data)

			cluster.CertificateAuthorityData = encoded
			cluster.CertificateAuthority = ""
		}

		c.Cluster = cluster

		config.Clusters[i] = c
	}

	return config, nil
}

func resolveKeyPaths(configAbsolutePath string, keyPath string) string {
	if keyPath[0] == '/' {
		return keyPath
	}

	baseDir := filepath.Dir(configAbsolutePath)

	keyPublicPath := baseDir + string(filepath.Separator) + keyPath

	keyPublicPath = filepath.Clean(keyPublicPath)

	return keyPublicPath
}
