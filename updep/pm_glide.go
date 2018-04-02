package updep

import (
	"os"
	"path"
)

var PmGlide = PackageManager{
	Name:     "glide",
	Language: LangGo,
	CheckFn: func(root string) bool {
		stat, err := os.Stat(path.Join(root, "glide.yaml"))
		return err == nil && !stat.IsDir()
	},
	// VersionFn:
	InstallCmd: "github.com/Masterminds/glide",
	RestoreCmd: "glide install",
	UpdateCmd:  "glide update",
}

func init() {
	PackageManagers["glide"] = PmGlide
}
