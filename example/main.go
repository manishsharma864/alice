// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"os"

	"github.com/manishsharma864/alice/example/dkg"
	"github.com/manishsharma864/alice/example/reshare"
	"github.com/manishsharma864/alice/example/signer"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "tss-example",
	Short: "TSS example",
	Long:  `This is a tss example`,
}

func init() {
	cmd.AddCommand(dkg.Cmd)
	cmd.AddCommand(signer.Cmd)
	cmd.AddCommand(reshare.Cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
