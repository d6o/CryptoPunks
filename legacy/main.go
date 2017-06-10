package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Punk struct {
	ID          int
	Type        string
	Owned       bool
	Accessories []string
}

func main() {
	allPunks := make([]Punk, 10000)

	for j := 0; j <= 9999; j++ {
		p := NewPunk(j)
		if p.Type == "Alien" || p.Type == "Ape" {
			fmt.Println(p)
		}
		allPunks[j] = *p
	}

	jPunks, err := json.Marshal(allPunks)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("punks.json", jPunks, 0777)
	if err != nil {
		panic(err)
	}
}
