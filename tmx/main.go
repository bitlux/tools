package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

const (
	kNotRunning = "no server running"
	kNoSuchFile = "No such file or directory"
)

func must(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	if os.Getenv("TMUX") != "" {
		return
	}

	path, err := exec.LookPath("tmux")
	must(err)

	out, err := exec.Command(path, "ls").CombinedOutput()
	sessions := strings.TrimSpace(string(out))
	if strings.Contains(sessions, kNoSuchFile) || strings.HasPrefix(sessions, kNotRunning) {
		return
	}
	must(err)
	color.Green(sessions)

	if len(os.Args) > 1 && os.Args[1] == "-a" {
		must(syscall.Exec(path, []string{path, "attach"}, os.Environ()))
	}
}
