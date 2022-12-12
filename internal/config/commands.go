package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Raw structure of Configuration files
// log config files found
type Raw struct {
	Exec     string `json:"exec"`
	Commands struct {
		Clean      string `json:"clean"`
		Autoremove string `json:"autoremove"`
		Build      string `json:"build"`
		Deps       string `json:"deps"`
		Depends    string `json:"depends"`
		Download   string `json:"download"`
		Fix        string `json:"fix"`
		Help       string `json:"help"`
		Info       string `json:"info"`
		Install    string `json:"install"`
		Installed  string `json:"installed"`
		Remove     string `json:"remove"`
		Search     string `json:"search"`
		Update     string `json:"update"`
		Upgrade    string `json:"upgrade"`
		Version    string `json:"version"`
	} `json:"commands"`
}

// All configuration files unmarshallowed
func All() []Raw {
	var result []Raw
	files := files()

	fmt.Println("Configuration files found: ")

	for _, file := range files {
		p := path.Join(Folder(), file.Name())
		fileInfo, err := os.Stat(p)

		// ignore broken symbolic link
		if os.IsNotExist(err) {
			continue
		}

		// ignore directories
		if fileInfo.IsDir() {
			continue
		}

		// ignore csv files (legacy)
		if ext := filepath.Ext(p); ext == ".yaml" {
			continue
		}

		fmt.Println(p)
		parsed := parse(p)
		result = append(result, parsed)
	}

	fmt.Println()
	return result
}

// array of configuratiion file name
func files() []fs.FileInfo {
	files, err := ioutil.ReadDir(Folder())

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func parse(filepath string) Raw {
	file, err := ioutil.ReadFile(filepath)
	var proj Raw

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &proj)

	if err != nil {
		log.Fatal(err)
	}

	return proj
}

// home folder of user
func home() string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return home
}

// Folder that config files will be looked up for
func Folder() string {
	result := path.Join(home(), ".config", "zae")

	return result
}
