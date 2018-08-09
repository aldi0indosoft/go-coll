package brick

import (
	"os"
	"os/exec"
)

func main() {
	// go build -o D:\javascript\vue-d3-rollup\yarn.exe
	arguments := []string{"/S", "/C", "docker run -it --rm -p 4000:8000 -v %CD%:/web --name brick andrius/alpine-webrick"}
	for i, arg := range os.Args {
		if i > 0 {
			arguments = append(arguments, arg)
		}
	}
	cmd := exec.Command("cmd", arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
