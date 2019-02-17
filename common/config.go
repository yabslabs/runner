/*
 * Copyright (c) 2018 VRSG | Verwaltungsrechenzentrum AG, St.Gallen
 * All Rights Reserved.
 */

package common

import (
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConfigReader interface {
	Unmarshal(data []byte, o interface{}) error
}

type ValidateableConfiguration interface {
	Validate() error
}

type ConfigReaderFunc func(data []byte, o interface{}) error

func (c ConfigReaderFunc) Unmarshal(data []byte, o interface{}) error {
	return c(data, o)
}

func readConfigFile(configReader ConfigReader, configFile string, obj interface{}) error {
	configFile = os.ExpandEnv(configFile)

	if _, err := os.Stat(configFile); err != nil {
		fmt.Println("CONF-3e126710", err, configFile)
		return nil
	}

	configStr, err := ioutil.ReadFile(configFile) //nolint: gosec

	if err != nil {
		return status.Errorf(codes.Internal, "failed to read config file %s: %v", configFile, err)
	}

	configStr = []byte(os.ExpandEnv(string(configStr)))

	if err := configReader.Unmarshal(configStr, obj); err != nil {
		return status.Errorf(codes.Internal, "error parse config file %s: %v", configFile, err)
	}

	return nil
}

// ReadConfig deserializes each configfile to the target obj using the configReader
// env vars are replaced in the file path
func ReadConfig(configReader ConfigReader, obj interface{}, configFiles ...string) error {
	for _, cf := range configFiles {
		if err := readConfigFile(configReader, cf, obj); err != nil {
			return err
		}
	}

	if validateable, ok := obj.(ValidateableConfiguration); ok {
		if err := validateable.Validate(); err != nil {
			return err
		}
	}

	return nil
}
