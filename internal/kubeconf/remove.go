package kubeconf

import (
	"github.com/pkg/errors"
)

func RemoveContexts(target string, contexts []string) error {
	config, err := LoadConfigAsSingle(target)
	if err != nil {
		return err
	}

	for _, ctxName := range contexts {
		err = removeContext(&config, ctxName)
		if err != nil {
			return err
		}
	}

	err = saveConfig(config, target)
	if err != nil {
		return err
	}

	return nil
}

func removeContext(config *KubeConfig, ctxName string) error {
	var rmUser string
	var rmCluster string
	canDeleteUser := true
	canDeleteCluster := true

	for i, ctx := range config.Contexts {
		if ctx.Name == ctxName {
			rmUser = ctx.Context.User
			rmCluster = ctx.Context.Cluster
			config.Contexts = append(config.Contexts[:i], config.Contexts[i+1:]...)
			break
		}
	}

	if len(rmUser) == 0 {
		return errors.New("Context not found:" + ctxName)
	}

	for _, ctx := range config.Contexts {
		if ctx.Context.User == rmUser {
			canDeleteUser = false
		}
		if ctx.Context.Cluster == rmCluster {
			canDeleteCluster = false
		}
	}

	if canDeleteUser {
		for i, user := range config.Users {
			if user.Name == rmUser {
				config.Users = append(config.Users[:i], config.Users[i+1:]...)
				break
			}
		}
	}

	if canDeleteCluster {
		for i, cluster := range config.Clusters {
			if cluster.Name == rmCluster {
				config.Clusters = append(config.Clusters[:i], config.Clusters[i+1:]...)
				break
			}
		}
	}

	return nil
}
