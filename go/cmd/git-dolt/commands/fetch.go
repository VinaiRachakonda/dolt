// Copyright 2019 Liquidata, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"

	"github.com/dolthub/dolt/go/cmd/git-dolt/config"
	"github.com/dolthub/dolt/go/cmd/git-dolt/doltops"
	"github.com/dolthub/dolt/go/cmd/git-dolt/utils"
)

// Fetch takes the filename of a git-dolt pointer file and clones
// the specified dolt repository to the specified revision.
func Fetch(ptrFname string) error {
	config, err := config.Load(ptrFname)
	if err != nil {
		return err
	}

	if err := doltops.CloneToRevision(config.Remote, config.Revision); err != nil {
		return err
	}

	fmt.Printf("Dolt repository cloned from remote %s to directory %s at revision %s\n", config.Remote, utils.LastSegment(config.Remote), config.Revision)
	return nil
}
