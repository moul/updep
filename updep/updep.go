package updep

var (
	Languages       = make(map[string]Language)
	PackageManagers = make(map[string]PackageManager)
)

type Language struct {
	Name string
}

type PackageManager struct {
	Name       string
	Language   Language
	CheckFn    func(string) bool
	InstallCmd string
	RestoreCmd string
	UpdateCmd  string
}

func DetectPackageManagers(root string) ([]PackageManager, error) {
	pms := make([]PackageManager, 0)
	for _, pm := range PackageManagers {
		if pm.CheckFn(root) {
			pms = append(pms, pm)
		}
	}
	return pms, nil
}
