package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	installed := make(map[string]bool)

	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
        }

		if strings.HasSuffix(path, ".go") {
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
			err = cmd.Run()
        }

		return err
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
