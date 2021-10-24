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
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// configure configure map value from config.toml(or other format) configure.
type configure struct {
	Application *struct {
		Name    string
		Version string
	}
	Database *struct {
		MYSQL *struct {
			Host     string
			Username string
			Password string
			Database string
		}
	}
}

type swaggerx struct {
	Title               string
	Version             string
	Schemes             []string
	Consumes            []string
	Produces            []string
	SecurityDefinitions SecurityDefinitions `mapstructure:"security_definitions"`
}

type SecurityToken struct {
	Type string
	Name string
	In   string
}

type SecurityDefinitions struct {
	AppTokenHeader    SecurityToken `mapstructure:"app_token_header"`
	AppTokenQuery     SecurityToken `mapstructure:"app_token_query"`
	ClientTokenHeader SecurityToken `mapstructure:"client_token_header"`
	ClientTokenQuery  SecurityToken `mapstructure:"client_token_query"`
}

func TestLoadConfigFile(t *testing.T) {
	err := LoadConfigureFile(WithSpecificConfigPath("test/data"))
	assert.NoError(t, err)
	actual := viper.Get("application.name")
	assert.Equal(t, "hellocms", actual)

	actual = viper.Get("database.mysql.host")
	assert.Equal(t, "localhost:3306", actual)
}

func TestNewConfigure(t *testing.T) {
	conf := NewConfigure(
		WithSpecificConfigPath("test/data"),
		WithSpecificConfigure(&configure{}),
	)
	assert.NotNil(t, conf)

	application := conf.(*configure).Application
	assert.Equal(t, "hellocms", application.Name)
	assert.Equal(t, "v1.2.3", application.Version)

	database := conf.(*configure).Database
	assert.Equal(t, "localhost:3306", database.MYSQL.Host)
	assert.Equal(t, "mysql", database.MYSQL.Database)
}

func TestLoadConfigFile_WithSwaggerxConfig(t *testing.T) {
	err := LoadConfigureFile(
		WithSpecificConfigPath("test/data"),
		WithSpecificConfigName("swaggerx"),
	)
	assert.NoError(t, err)
	actual := viper.Get("version")
	assert.Equal(t, "v1.2.3", actual)

	actual = viper.Get("security_definitions.app_token_header.type")
	assert.Equal(t, "apiKey", actual)
}

func TestNewConfigure_WithSwaggerxConfig(t *testing.T) {
	conf := NewConfigure(
		WithSpecificConfigPath("test/data"),
		WithSpecificConfigName("swaggerx"),
		WithSpecificConfigure(&swaggerx{}),
	)
	assert.NotNil(t, conf)

	swaggerx := conf.(*swaggerx)
	assert.Equal(t, "X-Access-Token", swaggerx.SecurityDefinitions.ClientTokenHeader.Name)
}
