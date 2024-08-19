package main

import "os/exec"

func main() {

	output, err := exec.Command("git", "branch").Output()

	if err != nil {
		println(err.Error())
	}

	println(string(output))

}
