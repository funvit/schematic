package genmodelcmd

import (
	"log"
	"os"

	"github.com/funvit/schematic/genrunner"
	"github.com/spf13/cobra"
)

var (
	target            string
	customPackageName string
)

var lInfo = log.New(os.Stdout, "INF ", log.LstdFlags)
var lErr = log.New(os.Stderr, "ERR ", log.LstdFlags)

var GenModelCmd = &cobra.Command{
	Use:     "model",
	Example: "schematic model ./schema --target ./model Schema1 Schema2",
	Args:    cobra.MinimumNArgs(2),
	Short:   "generate model by schema",
	Run: func(cmd *cobra.Command, args []string) {
		var p string
		if len(args) > 0 {
			p = args[0]
		}

		var models []string
		if len(args) > 1 {
			models = args[1:]
		} else {
			lErr.Fatalln("fatal: specify at least one schema name")
		}

		wd, err := os.Getwd()
		if err != nil {
			lErr.Println("get working dir:", err)
			os.Exit(1)
		}
		lInfo.Println("working dir:", wd)

		for _, m := range models {
			lInfo.Printf("processing schema %s\n", m)
			err := genrunner.Run(p, m, target, customPackageName)
			if err != nil {
				lErr.Printf("generate by schema %s: %s\n", m, err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	GenModelCmd.Flags().StringVar(&target, "target", ".", "target folder")
	GenModelCmd.Flags().StringVar(&customPackageName, "package", "model", "custom package name")
}
