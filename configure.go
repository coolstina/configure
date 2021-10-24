// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configure

import (
	"log"

	"github.com/spf13/viper"
)

// NewConfigure Read config file configure into the Configure struct.
func NewConfigure(ops ...Option) interface{} {
	var i interface{}
	options := options{
		configure: i,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	if err := LoadConfigureFile(ops...); err != nil {
		log.Fatalf("Load config file has error: %s\n", err.Error())
	}

	if err := viper.Unmarshal(&options.configure); err != nil {
		log.Fatalf("Parse the config file into the structure has error: %s\n", err.Error())
	}

	return options.configure
}

// LoadConfigureFile Load config file configure into the viper tools.
func LoadConfigureFile(ops ...Option) error {
	options := options{
		configType: "toml",
		configName: "config",
		configPath: ".",
	}

	for _, o := range ops {
		o.apply(&options)
	}

	viper.SetConfigType(options.configType)
	viper.SetConfigName(options.configName)
	viper.AddConfigPath(options.configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
