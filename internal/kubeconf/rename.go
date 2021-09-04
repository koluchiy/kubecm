package kubeconf

import (
	"errors"
)

func RenameContext(target string, from string, to string) error {
	config, err := LoadConfigAsSingle(target)
	if err != nil {
		return err
	}

	for _, ctx := range config.Contexts {
		if ctx.Name == to {
			return errors.New("Context already exists: " + to)
		}
	}

	for i, ctx := range config.Contexts {
		if ctx.Name == from {
			ctx.Name = to
			config.Contexts[i] = ctx
			err := saveConfig(config, target)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("Context not found: " + from)
}
