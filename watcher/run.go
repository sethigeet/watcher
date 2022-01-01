package watcher

import (
	"io"
	"os"
	"os/exec"

	"github.com/sethigeet/watcher/watcher/cmd"
)

var prevCmd *exec.Cmd = nil

func run() {
	// NOTE: prevCmd.ProcessState only gets populated after the cmd has exited
	if prevCmd != nil && prevCmd.ProcessState == nil {
		log.Warning("Killing previously running cmd process as it has not exited yet!")
		prevCmd.Process.Kill()
	}

	prevCmd = exec.Command(cmd.Config.CmdParts[0], cmd.Config.CmdParts[1:]...)

	stderr, err := prevCmd.StderrPipe()
	if err != nil {
		log.Fatalf("An error occurred!\nerror: %s\n", err)
	}

	stdout, err := prevCmd.StdoutPipe()
	if err != nil {
		log.Fatalf("An error occurred!\nerror: %s\n", err)
	}

	err = prevCmd.Start()
	if err != nil {
		log.Fatalf("An error occurred!\nerror: %s\n", err)
	}

	// run the copying in a goroutine so that it does not block
	go func() {
		io.Copy(os.Stdout, stdout)
		io.Copy(os.Stderr, stderr)

		// wait for the cmd to end so that prevCmd.ProcessState can be populated
		prevCmd.Wait()
	}()
}
