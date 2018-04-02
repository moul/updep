package updep

import (
	"os"
	"path"
)

var PmNpm = PackageManager{
	Name:     "npm",
	Language: LangJavascript,
	CheckFn: func(root string) bool {
		stat, err := os.Stat(path.Join(root, "package.json"))
		return err == nil && !stat.IsDir()
	},
}

func init() {
	PackageManagers["npm"] = PmNpm
}
