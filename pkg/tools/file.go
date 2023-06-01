package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"k8s-baseline-scanner/config"
	"os"
	"regexp"
)

var Writer writer

type writer struct {
}

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

func (*writer) WriteIntoFile(fileName string, data interface{}) {
	var (
		MarshalErr = "反序列化失败: "
		WIFErr     = "写入文件失败: "
		CFErr      = "创建文件失败: "
	)
	//确保数据目录存在，
	dirName := config.DataDirectory
	if !FileExist(dirName) {
		if err := makeFileExist(dirName); err != nil {
			panic(err)
		}
	}
	fileName = dirName + fileName
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
	//bytes, err1 := yaml.Marshal(&data)
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

// 清理文件中的null
func CleanNullInFile(filePath string) {
	oldStr := ":\\snull"
	newStr := ": \"\""
	// 读取文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		str := "读取文件失败: "
		logger.Error(errors.New(str + err.Error()))
		panic(errors.New(str + err.Error()))
	}

	// 将匹配正则表达式的字符串进行替换
	content := string(bytes)
	pattern := regexp.MustCompile(oldStr)
	newContent := pattern.ReplaceAllString(content, newStr)

	// 将替换后的内容写回到文件中
	err = ioutil.WriteFile(filePath, []byte(newContent), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
