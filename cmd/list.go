/*
Copyright © 2026 Anas Mirza <anasmirza534@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/anasmirza534/go-lang-todo-app/store"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `List all todos in database`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		doneFlag, err := cmd.Flags().GetBool("done")
		if err != nil {
			log.Fatal(err)
		}

		db, err := store.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var todos []store.Todo
		if doneFlag {
			todos, err = store.ListAllDoneTodos(db)
		} else {
			todos, err = store.ListAllTodos(db)
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(todos) == 0 {
			log.Println("No todos found.")
			return
		}

		log.Printf("%d todos found.", len(todos))
		for _, todo := range todos {
			log.Printf("[%s] %-70s(#%s)\n", todoStatus(todo.Done), todo.Title, todo.ID)
		}
	},
}

func todoStatus(done bool) string {
	if done {
		return "x"
	}
	return " "
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("done", "d", false, "show only completed todos")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
