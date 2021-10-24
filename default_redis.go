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

// Redis no-sql database configure structure.by default.
// You can use your custom define structure.
type Redis struct {
	// Redis server hostname, contain port section,
	// such as: localhost:6379.
	Host     string

	// Redis security verifies the password,
	// specify it if it exists,
	// otherwise you can leave it blank,
	// depending on how you use it.
	Password string


	// Database Specify the database number to use.
	Database int
}
