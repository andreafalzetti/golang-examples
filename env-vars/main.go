package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Env vars debug")
	cmd := exec.Command("node")
	env := cmd.Environ()
	fmt.Println(env)
}
