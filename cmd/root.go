package cmd

import (
	"errors"
	"fmt"
	"go-ripgrep/logics"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "grg",
		Short: "Go RipGrep",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, errors.New("Err: Not enough args"))
				os.Exit(1)
			}
			ignoreCase, _ := cmd.Flags().GetBool("ignoreCase")
			ignoreDotFiles, _ := cmd.Flags().GetBool("ignoreDotFiles")
			respectGitignore, _ := cmd.Flags().GetBool("respectGitignore")
			readFromGitFolder, _ := cmd.Flags().GetBool("readFromGitFolder")

			logics.Search(
				logics.SearchArgs{
					Root:           cmd.Flag("searchSpaceRoot").Value.String(),
					SearchStr:      args[0],
					IngoreCase:     ignoreCase,
					IgnoreDotfile:  ignoreDotFiles,
					ReadFromGitIgnore: respectGitignore,
					ReadFromGit:    readFromGitFolder,
				},
			)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	println("Starting Search ...")

	searchRoot := ""

	rootCmd.Flags().BoolP("ignoreCase", "i", false, "To ignore cases")
	rootCmd.Flags().BoolP("respectGitignore", "r", false, "To skip files in .gitignore")
	rootCmd.Flags().BoolP("ignoreDotFiles", "d", false, "To skip files/folder starting with .")
	rootCmd.Flags().BoolP("readFromGitFolder", "g", false, "To force read from .git folder")
	rootCmd.Flags().StringP("searchSpaceRoot", "w", searchRoot, "To force read from search folder")
}
