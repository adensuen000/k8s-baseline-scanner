package tools

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
	"os"
	"strings"
)

// 打印拉取成功的信息
func GetSuccessMsg(s string) {
	fmt.Printf("拉取%s成功. 请查看data目录.\n", s)
}

// 获取配置文件中某个配置项的值
func ReadConfig(module, key string) string {
	var (
		conf = "../config/config.ini"
	)

	if res := FileExist(conf); !res {
		msg := fmt.Sprintf("配置文件%s不存在，请检查.", conf)
		logger.Error(errors.New(msg))
		os.Exit(0)
	}

	viper.SetConfigType("ini")
	viper.SetConfigFile(conf)
	if err := viper.ReadInConfig(); err != nil {
		str := "无法读取配置文件: "
		logger.Error(errors.New(str + err.Error()))
		panic(errors.New(str + err.Error()))
	}
	//所有大写转换成小写
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return viper.GetString(module + "." + key)

}

func GetNamespace() []string {
	var (
		server, namespace = "server", "namespace"
	)
	nsstr := ReadConfig(server, namespace)
	nsSce := strings.FieldsFunc(nsstr, func(r rune) bool {
		return r == ',' || r == ' '
	})
	return nsSce
}

func GetKubeconfig() string {
	var (
		server, kubeconfig = "server", "kubeconfig"
	)
	return ReadConfig(server, kubeconfig)
}
