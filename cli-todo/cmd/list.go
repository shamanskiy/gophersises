package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all incomplete tasks in your task list",
	Run:   listCmdImpl,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listCmdImpl(cmd *cobra.Command, args []string) {
	tasks := getTaskList()

	fmt.Println("You have the following tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func getTaskList() (tasks []string) {
	appDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TASK_BUCKET))

		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			tasks = append(tasks, string(k))
		}

		return nil
	})

	return tasks
}
