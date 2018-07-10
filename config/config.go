// Copyright 2018 The Pouch Robot Authors
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

package config

// Config refers
type Config struct {
	Owner       string
	Repo        string
	HTTPListen  string
	AccessToken string
	Timeunit	string
	Time        int
}

// NewConfig creates a
func NewConfig() Config {
	return Config{
		Owner:      "",
		Repo:       "",
		HTTPListen: "",
		Timeunit: 	"s",
		Time:		1,
	}
}
