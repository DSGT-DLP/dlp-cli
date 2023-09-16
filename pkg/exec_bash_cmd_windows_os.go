//go:build windows

package pkg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func ExecBashCmd(dir string, name string, arg ...string) string {
	// Use this if the pty one doesn't work
	bash_cmd := exec.Command(name, arg...)
	bash_cmd.Dir = dir
	fmt.Println(strings.Join(bash_cmd.Args, " "))

	var stdoutBuf, stderrBuf bytes.Buffer
	bash_cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	bash_cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := bash_cmd.Run()
	if err != nil {
		fmt.Println("Error starting cmd: ", err)
		return fmt.Sprint(err)
	}

	outStr, errStr := string(stdoutBuf.String()), string(stderrBuf.String())
	return fmt.Sprint(outStr, errStr)
}
