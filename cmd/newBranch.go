/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// newBranchCmd represents the newBranch command
var newBranchCmd = &cobra.Command{
	Use:   "newBranch",
	Short: "This command will create a new branch",
	Long: `This command will create a new branch respecting the template defined. 
	The template will be users/name/branch-name. It will force users to not create branche a root of the repository
	users is a statics name (folder) to hold all branches
	name is the current user name (from git config) and hold all the user branches
	branch-name is the feature branch name
	For example:

	go-sample-cli newBranch <mybranchname>.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newBranch called")
		arglength := len(os.Args)
		if arglength != 3 {
			log.Fatal("No branch name provided")
		}
		//Get the current user name from the config
		log.Println("Get the current user name")
		var user = getCurrentUser()
		var branchName = fmt.Sprintf("users/%s/%s", user, os.Args[arglength-1])

		fmt.Println(branchName)
		_, err := exec.Command("git", "checkout", "-b", branchName).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Branch %s created.", branchName)
	},
}

func init() {
	rootCmd.AddCommand(newBranchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newBranchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newBranchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCurrentUser() string {
	//This function return the name of the current user
	response, err := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(response))
}
