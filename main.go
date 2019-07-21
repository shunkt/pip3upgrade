package main

import (
	"encoding/json"
	"log"
	"os/exec"
)

type OutdatedList struct {
	Name           string `json:"name"`
	Version        string `json:"version"`
	LatestVersion  string `json:"latest_version"`
	LatestFiletype string `json:"latest_filetype"`
}

func main() {
	list, err := exec.Command("pip3", "list", "--outdated", "--format=json").Output()
	if err != nil {
		log.Fatalln(err)
	}
	var outdated_list []OutdatedList
	if err := json.Unmarshal(list, &outdated_list); err != nil {
		log.Fatalln(err)
	}
	for _, pack := range outdated_list {
		err := exec.Command("pip3", "install", "-U", pack.Name).Run()
		if err != nil {
			log.Fatalln(err)
		} else {
			println("Updated:", pack.Name)
		}

	}
}
