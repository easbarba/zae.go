package commands

import (
	"fmt"

	"github.com/easbarba/zae/internal/config"
)

func run() {
	configs := config.All()

	for _, cfg := range configs {
		fmt.Println(cfg.Commands.Clean)
	}
}
