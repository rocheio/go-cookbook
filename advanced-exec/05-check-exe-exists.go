package main

import (
	"fmt"
	"os/exec"
)

func checkLsExists() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Printf("didn't find 'ls' executable\n")
		return
	}
	fmt.Printf("'ls' executable is in '%s'\n", path)
}

func main() {
	checkLsExists()
}