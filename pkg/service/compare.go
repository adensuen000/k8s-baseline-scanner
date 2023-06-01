package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"k8s-baseline-scanner/config"
	"k8s-baseline-scanner/pkg/tools"
	"os"
	"reflect"
)

func Compare(resourceName string) {
	var (
		fileList             []string
		jsonFile1, jsonFile2 string
		curBaselineDir       = config.DataDirectory
		stdBaselineDir       = config.StdBaselineDir
		PvcFile              = config.PvcFile
		VolumeFile           = config.VolumeFile
		ConfigmapFile        = config.ConfigmapFile
		ReplicasFile         = config.ReplicasFile
		AffinityFile         = config.AffinityFile
		NodeSelectorFile     = config.NodeSelectorFile
		TolerationFile       = config.TolerationFile
		EnvFile              = config.EnvFile
		ImageFile            = config.ImageFile
		ProbeFile            = config.ProbeFile
		VolumeMountFile      = config.VolumeMountFile
	)
	fileList = append(fileList, PvcFile, VolumeFile, ConfigmapFile, ReplicasFile, AffinityFile, NodeSelectorFile, TolerationFile, EnvFile, ImageFile, ProbeFile, VolumeMountFile)
	switch resourceName {
	case "all":
		for _, file := range fileList {
			jsonFile1 = curBaselineDir + file
			jsonFile2 = stdBaselineDir + file
			ExecuteCompare(jsonFile1, jsonFile2)
		}
	case "pvc":
		jsonFile1 = curBaselineDir + PvcFile
		jsonFile2 = stdBaselineDir + PvcFile
	case "volume":
		jsonFile1 = curBaselineDir + VolumeFile
		jsonFile2 = stdBaselineDir + VolumeFile
	case "cm":
		jsonFile1 = curBaselineDir + ConfigmapFile
		jsonFile2 = stdBaselineDir + ConfigmapFile
	case "replicas":
		jsonFile1 = curBaselineDir + ReplicasFile
		jsonFile2 = stdBaselineDir + ReplicasFile
	case "affinity":
		jsonFile1 = curBaselineDir + AffinityFile
		jsonFile2 = stdBaselineDir + AffinityFile
	case "nodeSelector":
		jsonFile1 = curBaselineDir + NodeSelectorFile
		jsonFile2 = stdBaselineDir + NodeSelectorFile
	case "toleration":
		jsonFile1 = curBaselineDir + TolerationFile
		jsonFile2 = stdBaselineDir + TolerationFile
	case "env":
		jsonFile1 = curBaselineDir + EnvFile
		jsonFile2 = stdBaselineDir + EnvFile
	case "image":
		jsonFile1 = curBaselineDir + ImageFile
		jsonFile2 = stdBaselineDir + ImageFile
	case "probe":
		jsonFile1 = curBaselineDir + ProbeFile
		jsonFile2 = stdBaselineDir + ProbeFile
	case "volumeMount":
		jsonFile1 = curBaselineDir + VolumeMountFile
		jsonFile2 = stdBaselineDir + VolumeMountFile
	}
	ExecuteCompare(jsonFile1, jsonFile2)

}

// 执行比较的方法
func ExecuteCompare(jsonFile1, jsonFile2 string) {
	var (
		i interface{}
	)
	if !tools.FileExist(jsonFile1) || !tools.FileExist(jsonFile2) {
		Err := fmt.Sprintf("要比对的两个文件%s和%s，有文件不存在，请检查.\n", jsonFile1, jsonFile2)
		logger.Error(errors.New(Err))
		panic(errors.New(Err))
	}
	data1 := Json2Interface(jsonFile1, i)
	data2 := Json2Interface(jsonFile2, i)
	recursiveCompare("", data1, data2)
}

func recursiveCompare(prefix string, v1, v2 interface{}) {
	// 获取v1和v2的类型和值
	t1, v1 := getDefaultValue(v1)
	t2, v2 := getDefaultValue(v2)

	// 如果类型不同，直接输出错误信息
	if t1 != t2 {
		fmt.Printf("%s: 不同类型:\n\t文件1: %v (类型：%T)\n\t文件2: %v (类型：%T)\n", prefix, v1, v1, v2, v2)
		return
	}
	// 如果类型是基本类型，则比较值
	switch t1.Kind() {
	case reflect.String, reflect.Bool, reflect.Float64, reflect.Int:
		if v1 != v2 {
			fmt.Printf("%s: 值不同:\n\t文件1: %v\n\t文件2: %v\n", prefix, v1, v2)
		}
	}

	// 如果类型是map，则比较key和value
	if t1.Kind() == reflect.Map {
		m1 := v1.(map[string]interface{})
		m2 := v2.(map[string]interface{})
		for k, v := range m1 {
			if v21, ok := m2[k]; ok {
				recursiveCompare(prefix+"."+k, v, v21)
			} else {
				fmt.Printf("%s.%s: 不存在于文件2\n", prefix, k)
			}
		}
		for k := range m2 {
			if _, ok := m1[k]; !ok {
				fmt.Printf("%s.%s: 不存在于文件1\n", prefix, k)
			}
		}
	}
	// 如果类型是slice，则比较元素值
	if t1.Kind() == reflect.Slice && reflect.TypeOf(v1).Elem().Kind() == reflect.Interface {
		s1 := v1.([]interface{})
		s2 := v2.([]interface{})
		for i, v := range s1 {
			if i < len(s2) {
				recursiveCompare(fmt.Sprintf("%s[%d]", prefix, i), v, s2[i])
			} else {
				fmt.Printf("%s[%d]: 不存在于文件2\n", prefix, i)
			}
		}
		for i := len(s1); i < len(s2); i++ {
			fmt.Printf("%s[%d]: 不存在于文件1\n", prefix, i)
		}
	}

	if t1.Kind() == reflect.Slice && reflect.TypeOf(v1).Elem().Kind() == reflect.String {
		s1 := v1.([]string)
		s2 := v2.([]string)
		for i, v := range s1 {
			if i < len(s2) {
				recursiveCompare(fmt.Sprintf("%s[%d]", prefix, i), v, s2[i])
			} else {
				fmt.Printf("%s[%d]: 不存在于文件2\n", prefix, i)
			}
		}
		for i := len(s1); i < len(s2); i++ {
			fmt.Printf("%s[%d]: 不存在于文件1\n", prefix, i)
		}
	}

}

func getDefaultValue(v interface{}) (reflect.Type, interface{}) {
	if reflect.TypeOf(v).Kind() == reflect.Interface {
		i := v.(map[string]interface{})
		t := reflect.TypeOf(i["type"])
		return t, i["value"]
	}
	return reflect.TypeOf(v), v
}

func Json2Interface(jsonFile string, i interface{}) interface{} {
	f, err := os.ReadFile(jsonFile)
	if err != nil {
		Err := fmt.Sprintf("读取文件%s失败: ", jsonFile)
		logger.Error(errors.New(Err + err.Error()))
		panic(errors.New(Err + err.Error()))
	}

	if err = json.Unmarshal(f, &i); err != nil {
		Err := fmt.Sprintf("解析文件%s失败: ", jsonFile)
		logger.Error(errors.New(Err + err.Error()))
		panic(errors.New(Err + err.Error()))
	}
	return i
}
