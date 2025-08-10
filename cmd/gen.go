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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	genCmd.Flags().StringP("destation", "d", "", "Destination directory")
	err := viper.BindPFlag("destation", genCmd.Flags().Lookup("destation"))
	if err != nil {
		panic(err)
	}

	genCmd.Flags().BoolP("tidy", "t", false, "Add dependencies to go.mod")
	err = viper.BindPFlag("tidy", genCmd.Flags().Lookup("tidy"))
	if err != nil {
		panic(err)
	}
}
