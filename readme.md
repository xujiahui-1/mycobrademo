命令行三要素
1.命令 git add
2.参数  go.mod
3.标志 -b



/Users/xu/go/bin/cobra-cli init   命令行初始化项目

cobra-cli add 添加子命令


- PersistentFlags() 持久化标志，可以传递给子命令
- Flags() 本地化标志，只有当前命令可以使用
  ```rootCmd.Flags().BoolP("viper", "v", true, "Help message for toggle")```
- 全局标志，也就是根命令上的持久化标志，所有的命令都可以使用