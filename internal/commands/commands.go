package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/easbarba/zae/internal/config"
)

func Run(command string, args ...string) {
	fmt.Println("commands: ", command, args)

	cmd := exec.Command(command, args...)

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(string(stdout))
}

func configs() (config.Raw, error) {
	cfgs := config.All()
	var err error

	for _, cfg := range cfgs {
		if _, err := os.Stat(cfg.Exec); err == nil {
			return cfg, nil
		}
	}

	// current distro is not listed in configuration files
	return config.Raw{}, err
}

func Found() config.Raw {
	result, err := configs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return result
}
