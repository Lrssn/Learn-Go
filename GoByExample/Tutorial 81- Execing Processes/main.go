package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("ls") // Create a process
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"} // Add args for process

	env := os.Environ() // Add envs for process

	execErr := syscall.Exec(binary, args, env) // Call process with args and envs
	if execErr != nil {                        // Check for errors
		panic(execErr) // Exit the go process
	} // And continue the created one
}
