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
	// var su string = ""

	fmt.Println(args)

	stdout, err = exec.Command(args[0], args[:1]...).Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

func configs() (config.Raw, error) {
	configs := config.All()
	var err error

	for _, cfg := range configs {
		if _, err := os.Stat(cfg.Exec); err == nil {
			return cfg, nil
		}
	}

	// current distro is not listed in configuration files
	return config.Raw{}, err
}

func Found() config.Raw {
	meh, err := configs()

	if err != nil {
		panic(err)
	}
	return meh
}
