package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/easbarba/zae/internal/config"
)

func Run(args ...string) {
	var stdout []byte
	var err error

	fmt.Println(args)

	stdout, err = exec.Command(args[0], args[:1]...).Output()

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
