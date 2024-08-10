package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dollboxpm <command> [options]")
		fmt.Println("Commands:")
		fmt.Println("  install      Install a package")
		fmt.Println("  list        List installed packages")
		fmt.Println("  update      Update a package")
		fmt.Println("  remove      Remove a package")
		fmt.Println("  --help      Show this help message")
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Usage: dollboxpm install <repository|branch> <package name>")
			return
		}
		installPackage(os.Args[2], os.Args[3])
	case "list":
		listPackages()
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Usage: dollboxpm update <package name>")
			return
		}
		updatePackage(os.Args[2])
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: dollboxpm remove <package name>")
			return
		}
		removePackage(os.Args[2])
	case "--help":
		fmt.Println("Usage: dollboxpm <command> [options]")
		fmt.Println("Commands:")
		fmt.Println("  install      Install a package")
		fmt.Println("  list        List installed packages")
		fmt.Println("  update      Update a package")
		fmt.Println("  remove      Remove a package")
		fmt.Println("  --help      Show this help message")
	default:
		fmt.Println("Unknown command:", cmd)
		return
	}
}

func installPackage(installType, packageName string) {
	var repoURL string
	if installType == "repository" {
		repoURL = fmt.Sprintf("https://github.com/DollBoxPM/%s", packageName)
	} else if installType == "branch" {
		repoURL = "https://github.com/DollBoxPM/DollBoxPM"
	} else {
		fmt.Println("Invalid install type:", installType)
		return
	}

	fmt.Println("Installing package from:", repoURL)

	// Create a temporary directory to clone the repository
	tmpDir, err := ioutil.TempDir("", "dollboxpm-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Clone the repository
	repo, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Copy the executable to the packages directory
	execPath := filepath.Join(tmpDir, packageName)
	if _, err := os.Stat(execPath); os.IsNotExist(err) {
		log.Fatal("Executable not found in repository.")
	}

	dstPath := filepath.Join("packages", packageName)
	if err := copyFile(execPath, dstPath); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Package installed successfully!")
}

func listPackages() {
	fmt.Println("Installed packages:")
	// List packages in the packages directory
	packagesDir := "packages"
	files, err := ioutil.ReadDir(packagesDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func updatePackage(packageName string) {
	fmt.Println("Updating package:", packageName)
	// Ask if the package was installed from a repository or branch
	fmt.Println("Was the package installed from a repository or branch?")
	fmt.Println("1. Repository")
	fmt.Println("2. Branch")
	var option int
	fmt.Scanln(&option)

	var repoURL string
	if option == 1 {
		repoURL = fmt.Sprintf("https://github.com/DollBoxPM/%s", packageName)
	} else if option == 2 {
		repoURL = "https://github.com/DollBoxPM/DollBoxPM"
	} else {
		fmt.Println("Invalid option. Please try again.")
		return
	}

	// Update the package
	fmt.Println("Updating package from:", repoURL)

	// Create a temporary directory to clone the repository
	tmpDir, err := ioutil.TempDir("", "dollboxpm-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Clone the repository
	repo, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Copy the executable to the packages directory
	execPath := filepath
