package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need to provide a project name")
		os.Exit(1)
	}

	err := createNewScript(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func createNewScript(name string) error {
	parentDir := filepath.Join("src", name)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return fmt.Errorf("creating parent directory: %w", err)
	}

	childDir := filepath.Join(parentDir, name)
	if err := os.MkdirAll(childDir, 0755); err != nil {
		return fmt.Errorf("creating child directory: %w", err)
	}

	srcMod := filepath.Join("src", "day1", "go.mod")
	dstMod := filepath.Join(parentDir, "go.mod")
	modContent, err := os.ReadFile(srcMod)
	if err != nil {
		return fmt.Errorf("reading go.mod: %w", err)
	}

	newContent := strings.Replace(string(modContent), "module day1", "module "+name, 1)
	if err := os.WriteFile(dstMod, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("writing modified go.mod: %w", err)
	}

	srcToml := filepath.Join("src", "day1", "gomod2nix.toml")
	dstToml := filepath.Join(parentDir, "gomod2nix.toml")
	if err := copyFile(srcToml, dstToml); err != nil {
		return fmt.Errorf("copying gomod2nix.toml: %w", err)
	}

	testData := filepath.Join(childDir, "test-data.txt")
	if err := os.WriteFile(testData, []byte{}, 0644); err != nil {
		return fmt.Errorf("creating test-data.txt: %w", err)
	}

	return nil
}

func copyFile(src, dest string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, input, 0644)
}
