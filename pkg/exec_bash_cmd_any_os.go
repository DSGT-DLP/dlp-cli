package pkg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pterm/pterm"
)

func execBashCmdAny(dir string, name string, arg ...string) string {
	// Use this if the pty one doesn't work
	bash_cmd := exec.Command(name, arg...)
	bash_cmd.Dir = dir
	pterm.DefaultHeader.WithFullWidth().Println(strings.Join(bash_cmd.Args, " "))

	var stdoutBuf, stderrBuf bytes.Buffer
	bash_cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	bash_cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := bash_cmd.Run()
	if err != nil {
		pterm.Error.Println("Error starting cmd: ", err)
		return fmt.Sprint(err)
	}

	outStr, errStr := string(stdoutBuf.String()), string(stderrBuf.String())
	return fmt.Sprint(outStr, errStr)
}
