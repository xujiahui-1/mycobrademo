/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run start ...")
		//打印 命令行标志
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.PersistentFlags().Lookup("author").Value,
			cmd.PersistentFlags().Lookup("config").Value,
			cmd.PersistentFlags().Lookup("license").Value,
			cmd.Flags().Lookup("source").Value,
		)

		fmt.Println("-------------------")
		//打印 viper
		fmt.Println(
			viper.GetString("author"),
			viper.GetString("license"),
		)
		fmt.Println("root cmd run end ...")
	},
	//配置全局标志传递给子命令
	TraverseChildren: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var configFile string
var userLicense string

func init() {
	//cobra 初始化的时候会去加载的方法
	cobra.OnInitialize(initConfig)
	// PersistentFlags() 持久化标志，可以传递给子命令
	rootCmd.PersistentFlags().BoolP("viper", "v", true, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Help message for toggle")
	// Flags() 本地化标志，只有当前命令可以使用
	rootCmd.Flags().StringP("source", "s", "", "Help message for toggle")

	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name")
	//  全局标志，也就是根命令上的持久化标志，所有的命令都可以使用

	//将 cobra 命令行参数与viper进行绑定
	//绑定viper的逻辑是，根据命令行或者配置文件改变viper中的值，不会改变cobra的值
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	viper.SetDefault("author", "defaultAuthor")
	viper.SetDefault("license", "defaultLicense")
}

// 通过viper将命令行参数绑定我们的配置文件
func initConfig() {
	if configFile != "" {
		//不为空，说明用户通过cobra传入配置文件了，那么用viper绑定该配置文件
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		//如果有报错，cobra会帮我们提醒命令行错误
		cobra.CheckErr(err)
		//配置默认的配置文件
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	//检查环境变量，将环境变量加载到viper
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
