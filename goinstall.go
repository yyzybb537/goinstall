package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


func walk(dirpath string, ifSymLink bool) {
	installed := make(map[string]bool)
	filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}

		if f.IsDir() {
			return nil
		}

		if f.Mode()&os.ModeSymlink != 0 {
			symlinkPath, err := filepath.EvalSymlinks(f.Name())
			if err != nil {
				println("Error:", err.Error())
			}
			walk(symlinkPath, true)
			return nil
		}

		if strings.HasSuffix(path, ".go") {
			if ifSymLink {
				path = strings.TrimPrefix(path, dirpath)
				path = filepath.Join(filepath.Base(dirpath), path)
			}
			dir := filepath.Dir(path)
			if installed[dir] {
				return nil
			}
			installed[dir] = true
			println("Process: go install", dir)
			cmd := exec.Command("go", "install")
			cmd.Dir = dir
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				println("Error:", err.Error())
			}
		}
		return nil
	})
}

func main() {
	walk(".", false)
}
