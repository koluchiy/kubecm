package kubeconf

import (
	"errors"
	"strings"
)

const inputSeparator = "::"
const aliasesSeparator = ","

func buildConfig(source string) (*Config, error) {
	parts := strings.Split(source, inputSeparator)
	configPath := parts[0]

	var configContextNames []string

	if len(parts) > 1 {
		configContextNames = strings.Split(parts[1], aliasesSeparator)
	}
	config, err := LoadConfigAsSingle(configPath)
	if err != nil {
		return nil, err
	}
	if len(configContextNames) > 0 {
		if len(configContextNames) != len(config.Contexts) {
			return nil, errors.New("bad context aliases count")
		}

		for i, ctxName := range configContextNames {
			config.Contexts[i].Name = ctxName
		}
	}

	result := Config{
		Path:       source,
		KubeConfig: config,
	}

	return &result, nil
}

func Merge(target string, sources []string) error {
	dUsers := NewDeduplicator()
	dClusters := NewDeduplicator()
	dContexts := NewDeduplicator()

	configs := make([]*Config, 0, len(sources))
	for _, input := range sources {
		config, err := buildConfig(input)
		if err != nil {
			return err
		}
		configs = append(configs, config)
	}

	targetConfig := KubeConfig{
		APIVersion:  "v1",
		Kind:        "Config",
		Preferences: nil,
	}

	for _, config := range configs {
		mapUsers := make(map[string]string)
		mapClusters := make(map[string]string)

		for _, user := range config.KubeConfig.Users {
			userName := dUsers.GetUniqueName(user.Name)
			mapUsers[user.Name] = userName

			targetConfig.Users = append(targetConfig.Users, User{
				Name: userName,
				User: user.User,
			})
		}

		for _, cluster := range config.KubeConfig.Clusters {
			clusterName := dClusters.GetUniqueName(cluster.Name)
			mapClusters[cluster.Name] = clusterName

			targetConfig.Clusters = append(targetConfig.Clusters, Cluster{
				Name:    clusterName,
				Cluster: cluster.Cluster,
			})
		}

		for _, ctx := range config.KubeConfig.Contexts {
			ctxName := dContexts.GetUniqueName(ctx.Name)

			targetConfig.Contexts = append(targetConfig.Contexts, Context{
				Name: ctxName,
				Context: ContextData{
					User:    mapUsers[ctx.Context.User],
					Cluster: mapClusters[ctx.Context.Cluster],
				},
			})
		}
	}

	targetConfig.CurrentContext = targetConfig.Contexts[0].Name

	err := saveConfig(targetConfig, target)
	if err != nil {
		return err
	}

	return nil
}
