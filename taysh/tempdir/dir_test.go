package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func TempDir() {

	prefix := "dev"

	// check if dev is true if so iterate for n and add n+1
	dirs, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dir := range dirs {
		if dir.IsDir() && len(dir.Name()) > len(prefix) && dir.Name()[:len(prefix)] == prefix {
			// if true increment end number in name
			num, err := strconv.Atoi(dir.Name()[len(prefix):])
			if err != nil {
				fmt.Println(err)
				return
			}
			prefix += strconv.Itoa(num + 1)
		}
	}

	// Create a temp directory to test your ideas
	tmpDir, err := ioutil.TempDir("", "prefix")
	if err != nil {
		fmt.Println(err)
		return
	}

	// cd and set temp Dir to pwd
	if err := os.Setenv("TEMP_DIR", tmpDir); err != nil {
		fmt.Println(err)
		return

	}

	// change the working directory to temp

	if err := os.Chdir(tmpDir); err != nil {
		fmt.Println(err)
		return
	}

	// generate go.mod file

	if err := ioutil.WriteFile("go.mod", []byte("module example.com"), 0644); err != nil {

		fmt.Println(err)
		return

	}

	// go.work file

	if err := ioutil.WriteFile("go.work", []byte("go 1.18\nuse ."), 0644); err != nil {

		fmt.Println(err)
		return

	}

	fmt.Printf("Your dev directory is set: %s\n", tmpDir)

	// Open the new dir in an elevated shell

	cmd := exec.Command("powershell.exe", "-Command", "Set-ExecutionPolicy Unrestricted '-Force'", "Start-Process powershell -Verb runAs -ArgumentList '-NoExit', '-Command', 'cd	"+tmpDir+"'")

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)

		if err := os.RemoveAll(tmpDir); err != nil {
			fmt.Printf("cannot remove temp dir: %v\n", err)
		}
	}
}
