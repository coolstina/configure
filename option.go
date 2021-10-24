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

// Option interface implementation Option mode.
type Option interface {
	apply(*options)
}

// options
type options struct {
	// Specifies the structure type used to map file
	// configuration to a concrete structure.
	configure interface{}

	// The file type can be any type supported by Viper.
	// Such as `toml/yaml/yml/json` and so on. Default by toml.
	configType string

	// Config file name, notice contains not file extension,
	// configType field is the specified extension.
	// Specific file names such as `config`. Default by config.
	configName string

	// Specific configure file path.
	// Such as:
	//		test/data
	//		config
	// The default is the directory where the program runs
	configPath string
}

// option function use implementation Option interface.
type optionFunc func(*options)

// apply the specified option configuration field.
func (o optionFunc) apply(ops *options) {
	o(ops)
}

// WithSpecificConfigType Use options to specify the type of configuration file.
func WithSpecificConfigType(configType string) Option {
	return optionFunc(func(ops *options) {
		ops.configType = configType
	})
}

// WithSpecificConfigName Use options to specify the name of configuration file.
func WithSpecificConfigName(configName string) Option {
	return optionFunc(func(ops *options) {
		ops.configName = configName
	})
}

// WithSpecificConfigPath Use options to specify the path of configuration file.
func WithSpecificConfigPath(configPath string) Option {
	return optionFunc(func(ops *options) {
		ops.configPath = configPath
	})
}

// WithSpecificConfigPath Use options to specify the type of configure.
func WithSpecificConfigure(configure interface{}) Option {
	return optionFunc(func(ops *options) {
		ops.configure = configure
	})
}
