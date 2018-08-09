package main

import (
	"os"
	"os/exec"
)

func main() {
	// go build -o D:\javascript\vue-d3-rollup\yarn.exe
	arguments := []string{"/S", "/C", "docker run -it --rm -p 4000:3000 -v %CD%:/home/node/app -u node --name dn dnr-docker-node-yarn"}
	for i, arg := range os.Args {
		if i > 0 {
			arguments = append(arguments, arg)
		}
	}
	// arguments := []string{"/C", "docker ps"}
	// fmt.Printf("os.Args: %v", os.Args)
	// fmt.Printf("arguments: %v", arguments)
	cmd := exec.Command("cmd", arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	// out, err := cmd.CombinedOutput()
	// fmt.Printf("%s\n", out)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
