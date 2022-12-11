package main

import (
	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&cli)

	// Call the Run() method of the selected parsed command.
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
	UpGrade   UpgradeCmd   `cmd help:"upgrade installed packages"`
}

type Context struct {
	Debug bool
}

type CleanCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *CleanCmd) Run(ctx *Context) error {

	return nil
}

type DepsCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *DepsCmd) Run(ctx *Context) error {

	return nil
}

type DependsCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *DependsCmd) Run(ctx *Context) error {

	return nil
}

type DownloadCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *DownloadCmd) Run(ctx *Context) error {

	return nil
}

type FixCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *FixCmd) Run(ctx *Context) error {

	return nil
}

type HelpCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *HelpCmd) Run(ctx *Context) error {

	return nil
}

type InfoCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *InfoCmd) Run(ctx *Context) error {

	return nil
}

type InstallCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *InstallCmd) Run(ctx *Context) error {

	return nil
}

type InstalledCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *InstalledCmd) Run(ctx *Context) error {

	return nil
}

type RemoveCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *RemoveCmd) Run(ctx *Context) error {

	return nil
}

type SearchCmd struct {
	Force     bool `help:"Force removal."`
	Recursive bool `help:"Recursively remove files."`

	Program string `arg name:"path" help:"Paths to remove." type:"path"`
}

func (r *SearchCmd) Run(ctx *Context) error {

	return nil
}

type UpdateCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *UpdateCmd) Run(ctx *Context) error {

	return nil
}

type UpgradeCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *UpgradeCmd) Run(ctx *Context) error {
	return nil
}
