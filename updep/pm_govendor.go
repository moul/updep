package updep

import (
	"os"
	"path"
)

var PmGovendor = PackageManager{
	Name:     "govendor",
	Language: LangGo,
	CheckFn: func(root string) bool {
		stat, err := os.Stat(path.Join(root, "vendor/vendor.json"))
		return err == nil && !stat.IsDir()
	},
	// VersionFn:
	InstallCmd: "github.com/kardianos/govendor",
	RestoreCmd: "govendor sync",
	UpdateCmd:  "govendor fetch",
}

func init() {
	PackageManagers["govendor"] = PmGovendor
}
