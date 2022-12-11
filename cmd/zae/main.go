package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&cli)

	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}

type Context struct {
	Debug bool
}

type SearchCmd struct {
	Force     bool `help:"Force removal."`
	Recursive bool `help:"Recursively remove files."`

	Program string `arg name:"path" help:"Paths to remove." type:"path"`
}

func (r *SearchCmd) Run(ctx *Context) error {
	fmt.Println("searching:", r.Program)
	return nil
}

type UpdateCmd struct {
	Program string `arg name:"program" optional:"" help:"Program to remove." type:"string"`
}

func (r *UpdateCmd) Run(ctx *Context) error {
	fmt.Println("updating:", r.Program)
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Search SearchCmd `cmd help:"search for matching packages of term."`
	Update UpdateCmd `cmd help:"update single package, or all system packages."`
}
