package env

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// 环境配置结构
type Env struct {
	HTTP_PORT string
}

var Config Env

func init() {
	pflag.String("env", "dev", "[dev]开发环境[test]测试环境[prod]生产环境")
	pflag.Parse()
	viper.BindPFlag("env", pflag.Lookup("env"))
	env := viper.GetString("env")
	fmt.Printf("当前程序运行的环境是:%s\n", env)
	// 从env文件读取环境配置
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AddConfigPath("./config/env/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %w", err))
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("配置文件解析失败: %w", err))
	}
}
