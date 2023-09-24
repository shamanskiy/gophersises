package cmd

import (
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

const (
	msgTaskAdded      = "Added task \"%s\" to your task list.\n"
	errMsgNoTaskNoAdd = `Error: no task provided. Please provide a task to add.
		
For example:
	cli-todo add new awesome task`
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task to your task list",
	Run:   addCmdImpl,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addCmdImpl(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println(errMsgNoTaskNoAdd)
		return
	}

	newTask := strings.Join(args, " ")

	appDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TASK_BUCKET))
		err := b.Put([]byte(newTask), []byte(""))
		return err
	})

	fmt.Printf(msgTaskAdded, newTask)
}
