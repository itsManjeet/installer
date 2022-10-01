package main

import (
	"fmt"
	"os"

	"github.com/itsmanjeet/installer/installer"
)

type progressbar string

func (p progressbar) Update(status int, mesg string) {
	fmt.Printf("[%d/100] %s\n", status, mesg)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %s <config-file>\n", os.Args[0])
		os.Exit(1)
	}

	configfile := os.Args[1]
	installer, err := installer.LoadConfig(configfile)
	if err != nil {
		fmt.Println("failed to load configuration file", err)
	}
	var p progressbar
	if err := installer.Verify(p); err != nil {
		fmt.Printf("Error! Verification failed, %s\n", err)
		os.Exit(1)
	}
	defer os.Remove(installer.Workdir)

	if err := installer.Install(p); err != nil {
		fmt.Printf("Error! Installation failed, %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Installation success")
}
