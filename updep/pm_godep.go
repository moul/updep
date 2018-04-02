package updep

import (
	"os"
	"path"
)

var PmGodep = PackageManager{
	Name:     "godep",
	Language: LangGo,
	CheckFn: func(root string) bool {
		stat, err := os.Stat(path.Join(root, "Godeps", "Godeps.json"))
		return err == nil && !stat.IsDir()
	},
	// VersionFn:
	InstallCmd: "github.com/tools/godep",
	RestoreCmd: "godep restore",
	UpdateCmd:  "godep update",
}

func init() {
	PackageManagers["godep"] = PmGodep
}
