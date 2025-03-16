package gomod

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/mod/modfile"
)

// Name godoc.
func Name() string {
	mod, err := os.ReadFile("go.mod")
	if err != nil {
		return currentPackage()
	}

	packagePath := modfile.ModulePath(mod)
	if packagePath == "" {
		return ""
	}

	return packagePath
}

// GoPath godoc.
func GoPath() string {
	return os.Getenv("GOPATH")
}

// GoPaths godoc.
func GoPaths() []string {
	gp := os.Getenv("GOPATH")
	if runtime.GOOS == "windows" {
		return strings.Split(gp, ";")
	}
	return strings.Split(gp, ":")
}

func importPath(path string) string {
	path = strings.TrimPrefix(path, "/private")
	for _, gopath := range GoPaths() {
		srcpath := filepath.Join(gopath, "src")
		rel, err := filepath.Rel(srcpath, path)
		if err == nil {
			return filepath.ToSlash(rel)
		}
	}

	rel := strings.TrimPrefix(path, filepath.Join(GoPath(), "src"))
	rel = strings.TrimPrefix(rel, string(filepath.Separator))
	return filepath.ToSlash(rel)
}

func currentPackage() string {
	pwd, _ := os.Getwd()
	return importPath(pwd)
}
