package main

import (
	"fmt"
	"log"

	"github.com/funvit/schematic/cmd/schematic/genmodelcmd"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(genmodelcmd.GenModelCmd)
}

func main() {
	fmt.Println("Schematic model by schema generator")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
