package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// 参数, 之前定义的都是我们的标志啊，子命令什么的，但是没搞参数
// 这里来搞参数，参数是什么  go run ./main.go 中， ./main.go就是参数
var argsCheckCmd = &cobra.Command{
	Use:   "cusargs",
	Short: "XxXXXXX",
	Long:  "XXXXXXXXXX",

	//命令行的参数检查现在Args参数里写
	//当然 cobra也给我们提供了内置的教研逻辑 cobra.NoArgs.....

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("最多输入两个参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cusargs cmd begin .......")

		fmt.Println(
			args,
		)
		fmt.Println("cusargs cmd done .......")
	},
}

func init() {
	rootCmd.AddCommand(argsCheckCmd)
}
