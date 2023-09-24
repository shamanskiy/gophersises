package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var appDB *bolt.DB

const TASK_BUCKET = "task-bucket"

var rootCmd = &cobra.Command{
	Use:   "cli-todo",
	Short: "Simple CLI to-do list",
	Long:  `Simple CLI to-do list: add and list tasks, then mark them complete`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	homeDir, err := homedir.Dir()
	shutDownOnErr(err)

	db, err := bolt.Open(homeDir+"/.todo.db", 0600, nil)
	shutDownOnErr(err)

	err = db.Update(createTaskBuckerIfNotExists)
	shutDownOnErr(err)

	appDB = db
}

func createTaskBuckerIfNotExists(tx *bolt.Tx) error {
	_, err := tx.CreateBucketIfNotExists([]byte(TASK_BUCKET))
	if err != nil {
		return fmt.Errorf("create bucket: %s", err)
	}
	return nil
}

func shutDownOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
