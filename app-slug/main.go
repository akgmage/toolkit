package main

import (
	"log"

	"github.com/akgmage/toolkit"
)

func main() {
	toSlug := "NOW!!!? is the time 12345"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)

	if err != nil {
		log.Println(err)
	}
	log.Println(slugified)
}