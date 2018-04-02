package updep

import (
	"errors"
	"os"
)

var (
	Languages       = make(map[string]Language)
	PackageManagers = make(map[string]PackageManager)
)

type Language struct {
	Name string
}

type PackageManager struct {
	Name         string
	Language     Language
	CheckFn      func(string) bool
	UpdateDepsFn func(string) error
}

func DetectPackageManagers(root string) ([]PackageManager, error) {
	stat, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New("not a directory")
	}
	pms := make([]PackageManager, 0)
	for _, pm := range PackageManagers {
		if pm.CheckFn(root) {
			pms = append(pms, pm)
		}
	}
	return pms, nil
}
