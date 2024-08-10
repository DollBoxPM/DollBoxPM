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
	fmt.Println("DollBoxPM - Package Manager")
	fmt.Println("-----------------------------------")

	fmt.Println("Do you want to download the package from:")
	fmt.Println("1. Repository (https://github.com/DollBoxPM/{package})")
	fmt.Println("2. Branch of the repository (https://github.com/DollBoxPM/DollBoxPM)")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		downloadFromRepository()
	case 2:
		downloadFromBranch()
	default:
		fmt.Println("Invalid option. Please try again.")
		return
	}
}

func downloadFromRepository() {
	fmt.Println("Enter the package name:")
	var packageName string
	fmt.Scanln(&packageName)

	repoURL := fmt.Sprintf("https://github.com/DollBoxPM/%s", packageName)
	fmt.Println("Downloading package from:", repoURL)

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

	// Copy the executable to the Linux executables directory
	execPath := filepath.Join(tmpDir, packageName)
	if _, err := os.Stat(execPath); os.IsNotExist(err) {
		log.Fatal("Executable not found in repository.")
	}

	dstPath := "/usr/local/bin/" + packageName
	if err := copyFile(execPath, dstPath); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Package downloaded and installed successfully!")
}

func downloadFromBranch() {
	fmt.Println("Enter the branch name:")
	var branchName string
	fmt.Scanln(&branchName)

	repoURL := "https://github.com/DollBoxPM/DollBoxPM"
	fmt.Println("Downloading package from:", repoURL)

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

	// Checkout to the specified branch
	if err := repo.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branchName),
	}); err != nil {
		log.Fatal(err)
	}

	// Copy the executable to the Linux executables directory
	execPath := filepath.Join(tmpDir, "DollBoxPM")
	if _, err := os.Stat(execPath); os.IsNotExist(err) {
		log.Fatal("Executable not found in repository.")
	}

	dstPath := "/usr/local/bin/DollBoxPM"
	if err := copyFile(execPath, dstPath); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Package downloaded and installed successfully!")
}

func copyFile(src, dst string) error {
	cmd := exec.Command("cp", src, dst)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error copying file: %s", output)
	}
	return nil
}

func handleError(err error) {
	fmt.Println("Error:", err)
	fmt.Println("Please submit an issue to: https://github.com/DollBoxPM/DollBoxPM/issues")
	os.Exit(1)
}