package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "schematic",
	Short: "schematic is a schema generator",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
