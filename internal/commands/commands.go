package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/easbarba/zae/internal/config"
)

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}

	return currentUser.Username == "root"
}

func Run(cmd []string) {
	var su string

	if isRoot() == false {
		su = "sudo "
	}

	command := su + strings.Join(cmd, " ")
	fmt.Println(command)

	cm := exec.Command("sh", "-c", command)

	stdout, err := cm.Output()
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
