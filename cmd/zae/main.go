package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/easbarba/zae/internal/config"
)

func system(cmd []string) {
	command := strings.Join(cmd, " ")
	fmt.Println(command)

	cm := exec.Command("sh", "-c", command) // TODO
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

func cfgFound() config.Raw {
	meh, err := configs()

	if err != nil {
		panic(err)
	}
	return meh
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Clean     CleanCmd     `cmd help:"clean system residual packages dependencies"`
	Deps      DepsCmd      `cmd help:"install dependencies packages of package"`
	Depends   DependsCmd   `cmd help:"list all dependencies of package"`
	Download  DownloadCmd  `cmd help:"download package binary"`
	Fix       FixCmd       `cmd help:"fix system package manager issues"`
	Help      HelpCmd      `cmd help:"shows a list of commands or help for one command"`
	Info      InfoCmd      `cmd help:"view info about a specific package"`
	Install   InstallCmd   `cmd help:"install package(s) from repositories"`
	Installed InstalledCmd `cmd help:"list all installed packages on system"`
	Remove    RemoveCmd    `cmd help:"remove one or more installed packages"`
	Search    SearchCmd    `cmd help:"search for matching packages of term."`
	Update    UpdateCmd    `cmd help:"update packages database"`
	Upgrade   UpgradeCmd   `cmd help:"upgrade installed packages"`
}

type Context struct {
	Debug bool
}

type CleanCmd struct {
}

func (r *CleanCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Clean})

	return nil
}

type DepsCmd struct {
	Package string `arg name:"package" help:"get package dependencies."`
}

func (r *DepsCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Deps, r.Package})

	return nil
}

type DependsCmd struct {
	Package string `arg name:"package" help:"list package dependencies."`
}

func (r *DependsCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Depends, r.Package})

	return nil
}

type DownloadCmd struct {
	Package string `arg name:"package"  help:"Download package binary."`
}

func (r *DownloadCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Download, r.Package})

	return nil
}

type FixCmd struct {
	Package string `arg name:"package" optional:"" help:"Fix package error."`
}

func (r *FixCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Fix, r.Package})

	return nil
}

type HelpCmd struct {
	Package string `arg name:"package" help:"command help usage."`
}

func (r *HelpCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Help, r.Package})

	return nil
}

type InfoCmd struct {
	Package string `arg name:"package" help:"Package information."`
}

func (r *InfoCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Info, r.Package})

	return nil
}

type InstallCmd struct {
	Package string `arg name:"package" help:"Package to install."`
}

func (r *InstallCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Install, r.Package})

	return nil
}

type InstalledCmd struct {
}

func (r *InstalledCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Installed})

	return nil
}

type RemoveCmd struct {
	Package string `arg name:"package" help:"Package to remove."`
}

func (r *RemoveCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Remove, r.Package})

	return nil
}

type SearchCmd struct {
	Force     bool `help:"Force removal."`
	Recursive bool `help:"Recursively remove files."`

	Package string `arg name:"package" help:"package to remove"`
}

func (r *SearchCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Search, r.Package})

	return nil
}

type UpdateCmd struct {
}

func (r *UpdateCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Update})

	return nil
}

type UpgradeCmd struct {
	Package string `arg name:"package" optional:"" help:"package to upgrade"`
}

func (r *UpgradeCmd) Run(ctx *Context) error {
	cfg := cfgFound()
	system([]string{cfg.Exec, cfg.Commands.Upgrade, r.Package})

	return nil
}
