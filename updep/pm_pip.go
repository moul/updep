package updep

import (
	"os"
	"path"
)

var PmPip = PackageManager{
	Name:     "pip",
	Language: LangPython,
	CheckFn: func(root string) bool {
		stat, err := os.Stat(path.Join(root, "requirements.txt"))
		return err == nil && !stat.IsDir()
	},
}

func init() {
	PackageManagers["pip"] = PmPip
}
