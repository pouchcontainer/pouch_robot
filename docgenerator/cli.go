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

package docgenerator

import (
	"fmt"
	"os/exec"
)

// generateCliDoc will generate Cli doc.
// First, use newly built binary to execute `binary gen-doc` to generate Cli doc.
// Second, git commit and push to github.
// Third, use github to create a new pull request.
//
// FIXME: this is specific for PouchContainer
// try to make it general.
func (g *Generator) generateCliDoc() error {
	// build a new binary cli client, since all cli doc is from newly built cli.
	cmd := exec.Command("make", "build-cli")
	if data, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to make client: output(%s), err(%v)", string(data), err)
	}

	// auto generate cli docs
	cmd = exec.Command("./bin/pouch", "gen-doc")
	if data, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to gen doc via cobra: output(%s), err(%v)", string(data), err)
	}

	return nil
}
