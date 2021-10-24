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

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coolstina/configure"
	"github.com/coolstina/configure/example/config"
)

func main() {

	conf := configure.NewConfigure(
		configure.WithSpecificConfigure(&config.Configure{}),
		configure.WithSpecificConfigPath("example"),
	).(*config.Configure)

	marshal, err := json.MarshalIndent(conf, "", "	")
	if err != nil {
		log.Panicf("JSON marshal has error: %s\n", err.Error())
	}

	//Out:
	//{
	//	"Application": {
	//	"Name": "coolstina",
	//		"Version": "v1.0.0"
	//},
	//	"Database": {
	//		"MYSQL": {
	//			"Host": "localhost:3306",
	//				"Username": "coolstina",
	//				"Password": "",
	//				"Database": "coolstina"
	//		},
	//		"Redis": {
	//			"Host": "localhost:6379",
	//				"Password": "",
	//				"Database": 5
	//		}
	//	}
	//}
	fmt.Println(string(marshal))
}
