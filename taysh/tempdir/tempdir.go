package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	// Create a temp directory to test your ideas
	tmpDir, err := ioutil.TempDir("", "dev")
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
