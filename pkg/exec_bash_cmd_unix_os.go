//go:build !windows

package pkg

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/creack/pty"
	"github.com/pterm/pterm"
	"golang.org/x/term"
)

func ExecBashCmd(runtime_os string, dir string, name string, arg ...string) string {
	if runtime_os == "windows" {
		return execBashCmdAny(dir, name, arg...)
	}
	// Code below found in pty examples: https://github.com/creack/pty
	bash_cmd := exec.Command(name, arg...)
	bash_cmd.Dir = dir
	pterm.DefaultHeader.WithFullWidth().Println(strings.Join(bash_cmd.Args, " "))
	ptmx, err := pty.Start(bash_cmd)
	if err != nil {
		panic(err)
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				pterm.Error.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Set stdin in raw mode.

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	// NOTE: The goroutine will keep reading until the next keystroke before returning.
	go func() { io.Copy(ptmx, os.Stdin) }()
	var buffer bytes.Buffer
	_, err = io.Copy(io.MultiWriter(os.Stdout, &buffer), ptmx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
