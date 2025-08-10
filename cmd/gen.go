/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/WhiCu/microgen/render"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AddDependency() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new Go microservice project",
	Long: `Generate a complete Go microservice project structure with all necessary 
components including handlers, services, configuration, and tests.

The generated project follows the standard Go project layout and includes:
  • HTTP server setup with Gin framework
  • Structured configuration management
  • Service layer architecture
  • HTTP handlers with proper error handling
  • Test files with authentication helpers
  • Client utilities for testing
  • Task automation with Taskfile

This command creates a production-ready foundation that you can immediately
start building your business logic upon.`,
	Example: `  microgen gen -d ./my-service
  microgen gen --destation ./api-service --tidy
  microgen gen -d ./user-service -t
  microgen gen -d ./payment-service --tidy`,
	Aliases: []string{"generate", "create", "new"},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := render.Render(viper.GetString("destation"))
		if err != nil {
			return err
		}
		if viper.GetBool("tidy") {
			err = AddDependency()
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genCmd.Flags().StringP("destation", "d", "", "Destination directory for the generated project (default: current directory)")
	err := viper.BindPFlag("destation", genCmd.Flags().Lookup("destation"))
	if err != nil {
		panic(err)
	}

	genCmd.Flags().BoolP("tidy", "t", false, "Automatically run 'go mod tidy' after generation to manage dependencies")
	err = viper.BindPFlag("tidy", genCmd.Flags().Lookup("tidy"))
	if err != nil {
		panic(err)
	}
}
