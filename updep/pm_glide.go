package updep

import (
	"os"
	"os/exec"
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
	UpdateDepsFn: func(workdir string) error {
		cmd := exec.Command("glide", "update")
		cmd.Dir = workdir
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stderr
		cmd.Stderr = os.Stderr
		return cmd.Run()
	},
}

func init() {
	PackageManagers["glide"] = PmGlide
}
