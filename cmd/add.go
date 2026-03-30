/*
Copyright © 2026 Anas Mirza <anasmirza534@gmail.com>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/anasmirza534/go-lang-todo-app/store"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new todo",
	Long:  `Add new todo by passing title`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		title, err := cmd.Flags().GetString("title")
		if err != nil {
			log.Fatal(err)
		}

		title = strings.TrimSpace(title)
		if len(title) == 0 {
			log.Fatal("Blank title is not allowed.")
		}

		db, err := store.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		todo, err := store.AddTodo(db, title)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Todo added #", todo.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.Flags().StringP("title", "t", "", "A title for todo")
	addCmd.MarkFlagRequired("title")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
