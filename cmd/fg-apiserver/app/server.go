package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/setcreed/fastgo/cmd/fg-apiserver/app/options"
	"github.com/setcreed/fastgo/pkg/version"
)

var configFile string // 配置文件路径

func NewFastGOCommand() *cobra.Command {
	// 创建默认的应用命令行选项
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		// 指定命令的名字，出现在帮助信息中
		Use: "fg-apiserver",
		// 命令的简短描述
		Short: "A very lightweight full go project",
		Long: `A very lightweight full go project, designed to help beginners quickly
		learn Go project development.`,
		// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：./fg-apiserver param1 param2
		Args: cobra.NoArgs,
	}
	// 初始化配置函数，在每个命令运行时调用
	cobra.OnInitialize(onInitialize)
	// cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the fg-apiserver configuration file.")

	// 添加 --version 标志
	version.AddFlags(cmd.PersistentFlags())
	return cmd
}

// run 是主运行逻辑，负责初始化日志、解析配置、校验选项并启动服务器。
func run(opts *options.ServerOptions) error {
	// 如果传入 --version，则打印版本信息并退出
	version.PrintAndExitIfRequested()

	// 将 viper 中的配置解析到选项 opts 变量中.
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}
	// 对命令行选项值进行校验.
	if err := opts.Validate(); err != nil {
		return err
	}
	// 获取应用配置.
	// 将命令行选项和应用配置分开，可以更加灵活的处理 2 种不同类型的配置.
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	// 创建服务器实例.
	server, err := cfg.NewServer()
	if err != nil {
		return err
	}
	// 启动服务器
	return server.Run()
}
