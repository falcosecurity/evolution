package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

func main() {
	commands := make([][]string, 0)
	currentIndex := -1
	var currentCommand []string
	for i, arg := range os.Args {
		if i == 0  || arg == "--" {
			if i != 0 {
				commands = append(commands, currentCommand)
			}
			currentCommand = make([]string, 0)
			currentIndex = currentIndex + 1
			continue
		}
		arg = strings.ReplaceAll(arg, "----", "--")
		currentCommand = append(currentCommand, arg)
	}
	if len(currentCommand) > 0 {
		commands = append(commands, currentCommand)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("cannot get cwd: %w", err))
	}

	for i, cmd := range commands {
		if i == 0 {
			continue
		}
		pid, err := syscall.ForkExec(cmd[0], cmd, &syscall.ProcAttr{
			Dir: cwd,
			Env: os.Environ(),
			Sys: &syscall.SysProcAttr{
				Foreground: false,
				Setsid: true,
			},
		})
		if err != nil {
			panic(fmt.Errorf("could not spawn %v: %w", cmd, err))
		}
		fmt.Printf("[Launcher] Spawned %v - pid %d\n", cmd, pid)
	}

	fmt.Printf("[Launcher] Performing exec %v\n", commands[0])
	err = syscall.Exec(commands[0][0], commands[0], os.Environ())
	if err != nil {
		panic(fmt.Errorf("could not exec: %w", err))
	}
}
