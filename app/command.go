package main

import (
    "fmt"
    "os/exec"
    "runtime"
)

const (
    OPEN_URL_CMD = "OPEN"
)

type runCmd func() error

func Parse(cmd string) runCmd {
    switch cmd {
    case OPEN_URL_CMD:
        return func() error { return open(URL) }
    default:
        fmt.Println("no command: ", cmd)
        return func() error { return nil }
    }
}

func open(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start"}
    case "darwin":
        cmd = "open"
    default:
        cmd = "xdg-open"
    }
    args = append(args, url)
    return exec.Command(cmd, args...).Start()
}
