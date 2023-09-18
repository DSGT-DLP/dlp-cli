//go:build windows

package pkg

func ExecBashCmd(runtime_os string, dir string, name string, arg ...string) string {
	execBashCmdAny(dir, name, arg...)
}
