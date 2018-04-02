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
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/moul/updep/updep"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info DIR [DIR...]",
	Short: "Display project-wide information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		maxlen := maxArgLength(args)
		for _, workdir := range args {
			workdirstr := fmt.Sprintf(fmt.Sprintf("%%-%ds", maxlen+1), fmt.Sprintf("%s:", workdir))
			pms, err := updep.DetectPackageManagers(workdir)
			if err != nil {
				fmt.Printf("%s %v\n", workdirstr, err)
				continue
			}
			if len(pms) == 0 {
				fmt.Printf("%s unknown ecosystem\n", workdirstr)
				continue
			}
			pmNames := []string{}
			langNames := []string{}
			for _, pm := range pms {
				pmNames = append(pmNames, pm.Name)
				langNames = append(langNames, pm.Language.Name)
			}
			fmt.Printf("%s lang=%s pm=%s\n", workdirstr, strings.Join(langNames, ","), strings.Join(pmNames, ", "))
		}
	},
}

func maxArgLength(args []string) int {
	maxlen := 0
	for _, arg := range args {
		if newlen := len(arg); newlen > maxlen {
			maxlen = newlen
		}
	}
	return maxlen
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
