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
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/moul/updep/updep"
)

var workDir string

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display project-wide information",
	Run: func(cmd *cobra.Command, args []string) {
		pms, err := updep.DetectPackageManagers(workDir)
		if err != nil {
			log.Fatal(err)
		}
		if len(pms) == 0 {
			log.Fatal(fmt.Errorf("no package manager detected in %s", workDir))
		}
		pmNames := []string{}
		langNames := []string{}
		for _, pm := range pms {
			pmNames = append(pmNames, pm.Name)
			langNames = append(langNames, pm.Language.Name)
		}
		fmt.Printf("Detected Languages:         %s\n", strings.Join(langNames, ", "))
		fmt.Printf("Detected Package Managers:  %s\n", strings.Join(pmNames, ", "))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVar(&workDir, "workdir", "", "Project directory")
	viper.SetDefault("workdir", ".")
}
