/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package quests

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor/cmd"
)

// questsCmd represents the quests command
var questsCmd = &cobra.Command{
	Use:   "quests",
	Short: "Commands for managing quests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quests called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(questsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
