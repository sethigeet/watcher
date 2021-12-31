package watcher

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/sethigeet/watcher/watcher/cmd"
)

func run() (string, bool) {
	cmdParts := strings.Split(*cmd.Config.Command, " ")
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("An error occurred!\nerror: %s\n", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("An error occurred!\nerror: %s\n", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("An error occurred!\nerror: %s\n", err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
