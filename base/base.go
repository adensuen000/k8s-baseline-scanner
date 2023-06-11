package base

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"strings"
)

const (
	conf     = ConfigDirectory
	confType = "ini"
)

var (
	//配置文件中的模块, 如server
	module string
	//配置文件中的配置项, 如kubeconfig
	confItem string
)

// 判断目录或者文件存在
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 创建目录或者文件
func makeFileExist(fileName string) error {
	err := os.MkdirAll(fileName, 0755)
	if err != nil {
		errStr := fmt.Sprintf("创建目录%s失败: ", fileName)
		logger.Error(errors.New(errStr + err.Error()))
		return errors.New(errStr + err.Error())
	}
	return nil
}

func writeIntoFile(fileName string, data interface{}) {
	var (
		MarshalErr = "反序列化失败: "
		WIFErr     = "写入文件失败: "
		CFErr      = "创建文件失败: "
	)

	//重新创建文件，旨在清空文件内容
	_, err := os.Create(fileName)
	if err != nil {
		logger.Error(errors.New(CFErr))
		panic(errors.New(CFErr))
	}

	//打开文件并写入数据
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {
		}
	}(file)
	jsonData, err1 := json.MarshalIndent(&data, "", "    ")
	if err1 != nil {
		logger.Error(errors.New(MarshalErr + err1.Error()))
		panic(errors.New(MarshalErr + err1.Error()))
	}
	_, err = file.WriteString(string(jsonData))
	if err != nil {
		logger.Error(errors.New(WIFErr + err.Error()))
		panic(errors.New(WIFErr + err.Error()))
	}
}

// 拉取成功后输出信息
func getSuccessMsg(s string) {
	fmt.Printf("拉取%s成功. 请查看data目录.\n", s)
}

func WriteDataIntoFile(msgName, fileName string, data interface{}) {
	var dataDir = DataDirectory
	//确保数据目录存在
	if !FileExist(dataDir) {
		if err := makeFileExist(dataDir); err != nil {
			panic(err)
		}
	}
	fileName = dataDir + fileName
	writeIntoFile(fileName, data)
	//cleanNullInFile(fileName)
	getSuccessMsg(msgName)
}

// 获取配置文件中某个配置项的值
func readConfig(module, confItem string) string {
	if res := FileExist(conf); !res {
		msg := fmt.Sprintf("配置文件%s不存在，请检查.", conf)
		logger.Error(errors.New(msg))
		os.Exit(0)
	}

	viper.SetConfigType(confType)
	viper.SetConfigFile(conf)
	if err := viper.ReadInConfig(); err != nil {
		str := "无法读取配置文件: "
		logger.Error(errors.New(str + err.Error()))
		panic(errors.New(str + err.Error()))
	}
	//所有大写转换成小写
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return viper.GetString(module + "." + confItem)
}

// 获取配置项
func getConfigItem(module, confItem string) string {
	return readConfig(module, confItem)
}

// 获取Kubeconfig配置项
func GetKubeconfig() string {
	module, confItem = "server", "kubeconfig"
	return getConfigItem(module, confItem)
}

// 获取命名空间配置项
func getNamespaceItem() string {
	module, confItem = "server", "namespace"
	return getConfigItem(module, confItem)
}

func GetNamespaceList() []string {
	var nsSce []string
	nsList := getNamespaceItem()
	if nsList == "all" {
		allNsList, err := K8sInit.GetClientSet().CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New("获取命名空间失败: " + err.Error()))
			panic(errors.New("获取命名空间失败: " + err.Error()))
		}
		for _, allNs := range allNsList.Items {
			nsSce = append(nsSce, allNs.Name)
		}
		return nsSce
	}
	nsSce = strings.FieldsFunc(nsList, func(r rune) bool {
		return r == ',' || r == ' '
	})
	return nsSce
}
