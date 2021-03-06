// Copyright © 2018 Moises P. Sena <moisespsena@gmail.com>
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

package cmd

import (
	"fmt"

	"github.com/moisespsena-go/error-wrap"
	"github.com/moisespsena-go/goenv"
	"github.com/spf13/cobra"
)

var versionsGetCmd = &cobra.Command{
	Use:   "get VERSION...",
	Args:  cobra.MinimumNArgs(1),
	Short: "Download one or more GoLang versions and save files into $GOENVROOT/.versions dir",
	RunE: func(cmd *cobra.Command, args []string) error {
		env, err := goenv.NewGoEnv(db, false)
		if err != nil {
			return errwrap.Wrap(err, "New Env")
		}
		v := goenv.NewGoVersions(env)
		items, err := v.Download(args...)
		if err != nil {
			return err
		}
		for _, v := range items {
			fmt.Println(pad(v.Name), v.DownloadPath())
		}
		return nil
	},
}

func init() {
	versionsCmd.AddCommand(versionsGetCmd)
}
