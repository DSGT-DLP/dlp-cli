//go:build windows

package pkg

func ExecBashCmd(runtime_os string, dir string, name string, arg ...string) string {
	return execBashCmdAny(dir, name, arg...)
}
