package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)
// The main Idea for tempdir is to be able to configure a temp environment within a drive location to test packaging and functions, 
// needs to be iterated "dev*n*" => "devn+1n" and not write the name per stdin iteration. 
// simple: call tempdir, load go.mod. go.work, set $ENV set TEMP_DRIVE && cd into new "dev" drive, if stdcout := ",_ "tayshell tempdir --done" then break and remove "dev", ask if 
// a log of the time and exit codes should be put to a log and saved, before recursively removing temp_dir *// 

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
