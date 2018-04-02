package updep

import (
	"os"
	"path"
)

var PmGovendor = PackageManager{
	Name:     "govendor",
	Language: LangGo,
	CheckFn: func(root string) bool {
		stat, err := os.Stat(path.Join(root, "vendor", "vendor.json"))
		return err == nil && !stat.IsDir()
	},
}

func init() {
	PackageManagers["govendor"] = PmGovendor
}
