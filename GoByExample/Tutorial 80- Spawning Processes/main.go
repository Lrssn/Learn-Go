package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	dateCmd := exec.Command("date") // Create external command

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output() // Check if the command gives an error
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err) // If calling the command fails
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode()) // If command returns code != 0
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()   // Grep inpipe
	grepOut, _ := grepCmd.StdoutPipe() // Grep outpipe
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep")) // Write to inpipe
	grepIn.Close()                                   // Close inpipe
	grepBytes, _ := io.ReadAll(grepOut)              // Read to outpipe
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h") // Call bash command using flags
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
