package service

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"k8s-baseline-scanner-v2/base"
	"k8s.io/apimachinery/pkg/util/json"
	"os"
)

// 输出比对结果
func getCompareRes(msg string, ResList []string) {
	if ResList == nil {
		fmt.Printf("当前环境的%s与标准基线中的相同.\n", msg)
	} else {
		fmt.Printf("当前环境的%s与标准基线中的不相同,对应应用如下:\n", msg)
		for i, v := range ResList {
			fmt.Printf("%d: %s\n", i+1, v)
		}
	}
}

// 检查两个对比文件是否存在
func checkFilesExist(filepath1, filepath2 string) {
	if !base.FileExist(filepath1) || !base.FileExist(filepath2) {
		msg := fmt.Sprintf("%s或者%s不存在，请检查.\n", filepath1, filepath2)
		logger.Error(errors.New(msg))
		os.Exit(0)
	}
}

func json2Map(filePath string) map[string]interface{} {
	data := make(map[string]interface{})
	f, err := os.ReadFile(filePath)
	if err != nil {
		Err := fmt.Sprintf("读取文件%s失败: ", filePath)
		logger.Error(errors.New(Err + err.Error()))
		panic(errors.New(Err + err.Error()))
	}
	if err = json.Unmarshal(f, &data); err != nil {
		errStr := "反序列还失败: "
		logger.Error(errors.New(errStr + err.Error()))
		panic(errors.New(errStr + err.Error()))
	}
	return data
}
