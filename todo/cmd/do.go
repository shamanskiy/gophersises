package cmd

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task in your task list complete",
	Run:   doCmdImpl,
}

func init() {
	rootCmd.AddCommand(doCmd)
}

func doCmdImpl(cmd *cobra.Command, args []string) {
	tasks := getTaskList()

	if len(args) == 0 {
		fmt.Print("Error: no task provided. Please provide a task number to remove.\n\n")
		printTaskListWithHints(tasks)
		return
	}

	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Print("Error: invalid task number provided. Please provide a valid task number.\n\n")
		printTaskListWithHints(tasks)
		return
	}
	taskId = taskId - 1

	if taskId < 0 || taskId >= len(tasks) {
		fmt.Print("Error: out-of-bound task number provided. Please provide a task number of one of the tasks.\n\n")
		printTaskListWithHints(tasks)
		return
	}

	taskToComplete := tasks[taskId]

	err = appDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TASK_BUCKET))
		err := b.Delete([]byte(taskToComplete))
		return err
	})
	shutDownOnErr(err)

	fmt.Printf("Mark task \"%s\" as complete.\n\n", taskToComplete)

	printTaskList(getTaskList())
}

func printTaskListWithHints(tasks []string) {
	printTaskList(tasks)
	if len(tasks) > 0 {
		fmt.Printf("\nFor example, you can mark \"%s\" task as complete with:\n", tasks[0])
		fmt.Println("  todo do 1")
	}
}
