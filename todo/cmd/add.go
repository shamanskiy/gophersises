package cmd

import (
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task to your task list",
	Long: `Add new task to your task list. 

For example, to add a task "water the plants", execute:
  todo add water the plants`,
	Run: addCmdImpl,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addCmdImpl(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println(`Error: no task provided. Please provide a task to add.
		
For example:
  todo add water the plants`)
		return
	}

	newTask := strings.Join(args, " ")

	appDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TASK_BUCKET))
		err := b.Put([]byte(newTask), []byte(""))
		return err
	})

	fmt.Printf("Added task \"%s\" to your task list.\n", newTask)
}
