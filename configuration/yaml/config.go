/*
 * Copyright (c) 2018 VRSG | Verwaltungsrechenzentrum AG, St.Gallen
 * All Rights Reserved.
 */

package yaml

import (
	"git.workshop21.ch/go/abraxas/configuration"
	"github.com/ghodss/yaml"
)

var ConfigReader = configuration.ConfigReaderFunc(yamlUnmarshalNoOpts)

func yamlUnmarshalNoOpts(y []byte, o interface{}) error {
	return yaml.Unmarshal(y, o)
}

func ReadConfig(obj interface{}, configFiles ...string) error {
	return configuration.ReadConfig(ConfigReader, obj, configFiles...)
}
