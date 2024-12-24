package container

import "path/filepath"

// GetServiceDir retrives service directory
func GetServiceDir(service string) string {
	base := "../services/" + service
	path, err := filepath.Abs(base)
	if (err != nil) {
		panic(err)
	}
	
	return filepath.ToSlash(path)
}
