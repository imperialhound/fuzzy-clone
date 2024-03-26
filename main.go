package main

import (
	"fmt"

	"github.com/imperialhound/fuzzy-clone/internal/config"
)

func main() {
  // Get fuzzy-cloner configuration file
  conf, err := config.New()
  if err != nil {
    panic(err)
  }

  fmt.Println(conf)

  // TODO(dpe): Fuzzy select source

  // TODO(dpe): Fuzzy select repo
}
