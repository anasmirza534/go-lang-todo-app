/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/anasmirza534/go-lang-todo-app/store"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle todo item",
	Long:  `Toggle todo item for given id`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}
		id = strings.TrimSpace(id)

		db, err := store.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = store.ToggleTodo(db, id)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Todo is toggled.")
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)

	toggleCmd.Flags().String("id", "", "ID of the todo item")
	toggleCmd.MarkFlagRequired("id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
