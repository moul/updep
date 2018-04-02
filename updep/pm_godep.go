package updep

import (
	"os"
	"os/exec"
	"path"
)

var PmGodep = PackageManager{
	Name:     "godep",
	Language: LangGo,
	CheckFn: func(root string) bool {
		stat, err := os.Stat(path.Join(root, "Godeps", "Godeps.json"))
		return err == nil && !stat.IsDir()
	},
	UpdateDepsFn: func(workdir string) error {
		cmd := exec.Command("godep", "update", "./...")
		cmd.Dir = workdir
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stderr
		cmd.Stderr = os.Stderr
		return cmd.Run()
	},
}

func init() {
	PackageManagers["godep"] = PmGodep
}
