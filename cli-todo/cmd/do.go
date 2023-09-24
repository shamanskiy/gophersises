package cmd

import (
	"fmt"

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
	fmt.Println("do called")
}
