// Copyright © 2018 Manfred Touron <m@42.am>
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
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/moul/updep/updep"
)

var (
	autoPr bool
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update DIR [DIR...]",
	Short: "Update dependencies",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := "updep-" + time.Now().Format("20060102-150405")
		updateMessage := "chore: update dependencies (updep)"

		if len(args) == 0 {
			args = []string{"."}
		}
		found := false
		for _, workdir := range args {
			// FIXME: if autoPr is on, check if workdir is clean before beginning, if yes -> stash or block
			pms, err := updep.DetectPackageManagers(workdir)
			if err != nil {
				log.Fatal(err)
			}
			if len(pms) == 0 {
				continue
			}
			found = true
			for _, pm := range pms {
				if err := pm.UpdateDepsFn(workdir); err != nil {
					log.Print("cannot update deps for %s: %v", workdir, err)
				}
			}
		}
		if !found {
			log.Fatalf("no ecosystem detected in %s", strings.Join(args, ", "))
		}

		if autoPr {
			for _, workdir := range args {
				cmd := exec.Command("git", "add", ".")
				cmd.Dir = workdir
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stderr
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
				cmd = exec.Command("git", "checkout", "-b", branchName)
				cmd.Dir = workdir
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stderr
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
				cmd = exec.Command("git", "commit", ".", "-m", updateMessage)
				cmd.Dir = workdir
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stderr
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
				cmd = exec.Command("git", "push", "-u", "origin", branchName)
				cmd.Dir = workdir
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stderr
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
				cmd = exec.Command("hub", "pull-request", "-m", updateMessage)
				cmd.Dir = workdir
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stderr
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVar(&autoPr, "pr", false, "Automatically commit, push and open a PR with updated dependencies")
	// FIXME: support --dry-run
}
