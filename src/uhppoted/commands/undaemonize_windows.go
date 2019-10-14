package commands

import (
	"errors"
	"fmt"
)

type Undaemonize struct {
}

func (c *Undaemonize) Execute(ctx Context) error {
	return errors.New("uhppoted undaemonize: NOT IMPLEMENTED")
}

func (c *Undaemonize) Cmd() string {
	return "daemonize"
}

func (c *Undaemonize) Description() string {
	return "Deregisters the uhppoted service"
}

func (c *Undaemonize) Usage() string {
	return ""
}

func (c *Undaemonize) Help() {
	fmt.Println("Usage: uhppoted undaemonize")
	fmt.Println()
	fmt.Println(" Deregisters uhppoted as a Windows service")
	fmt.Println()
}
