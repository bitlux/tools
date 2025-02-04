package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const (
	kNotRunning = "no server running"
)

func must(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	path, err := exec.LookPath("tmux")
	must(err)

	out, err := exec.Command(path, "ls").CombinedOutput()
	sessions := strings.TrimSpace(string(out))
	if strings.HasPrefix(sessions, kNotRunning) {
		return
	}

	must(err)

	lines := strings.Split(sessions, "\n")
	if n := len(lines); n > 1 {
		fmt.Println("found", n, "tmux sessions")
		return
	}
	must(syscall.Exec(path, []string{"attach"}, os.Environ()))
}
