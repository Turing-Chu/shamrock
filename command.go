// Author: Turing Zhu
// Date: 6/11/21 3:13 PM
// File: command.go

package shamrock

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

// capture Stderr andStdout
func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, _ = w.Write(d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

// run shell command
func Run(command string, args []string) (stdOutput, errOutput string, err error) {
	cmd := exec.Command(command, args...)
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	_ = cmd.Start()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()
	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
		wg.Done()
	}()
	err = cmd.Wait()
	if err != nil {
		return "", "", fmt.Errorf("cmd.Run() failed with %s", err)
	}

	wg.Wait()

	if errStdout != nil || errStderr != nil {
		return "", "", fmt.Errorf("failed to capture stdout or stderr")
	}
	return string(stdout), string(stderr), nil
}
