/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// 注意：当子命令集成服命令中的全局标志的时候，用cmd.Flags()访问
// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called.....")
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("author").Value,
			viper.GetString("author"),
			cmd.Flags().Lookup("config").Value,
			cmd.Flags().Lookup("license").Value,
			//访问父级标识
			cmd.Parent().Flags().Lookup("source").Value,
		)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
