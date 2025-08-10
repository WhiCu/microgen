/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "microgen",
	Short: "Generate Go microservice project structure with best practices",
	Long: `MicroGen is a powerful CLI tool for generating Go microservice projects with 
industry best practices and modern architecture patterns.

Features:
  • Pre-configured project structure following Go project layout standards
  • Built-in templates for HTTP handlers, services, and middleware
  • Automatic dependency management with go.mod
  • Configurable destination directories
  • Clean and maintainable code generation

MicroGen helps developers quickly bootstrap production-ready microservices
by providing a solid foundation that follows Go community conventions.`,
	Example: `  microgen gen -d ./my-service
  microgen gen --destation ./api-service --tidy
  microgen gen -d ./user-service -t`,
	Aliases: []string{"mg", "micro"},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.microgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
