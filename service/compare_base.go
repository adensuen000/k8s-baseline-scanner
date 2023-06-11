package service

import (
	"fmt"
	"reflect"
)

// 比较两个接口类型的数据是否相等
func compareInterface(m1, m2 interface{}) bool {
	t1 := reflect.TypeOf(m1)
	t2 := reflect.TypeOf(m2)
	if t1 != t2 {
		return false
	}
	switch t1.Kind() {
	case reflect.String, reflect.Bool, reflect.Float64, reflect.Int, reflect.Int32, reflect.Int64:
		if m1 != m2 {
			return false
		} else {
			return true
		}
	case reflect.Map:
		m11, m22 := m1.(map[string]interface{}), m2.(map[string]interface{})
		if m11 == nil && m22 == nil {
			return true
		} else if (m11 == nil && m22 != nil) || (m11 != nil && m22 == nil) {
			return false
		}

		for k11, v11 := range m11 {
			if v22, ok := m22[k11]; ok {
				return compareNotNull(v11, v22)
			}
			return false
		}
		for k22, _ := range m22 {
			if _, ok := m11[k22]; !ok {
				return false
			}
		}
	}

	if t1.Kind() == reflect.Slice && reflect.TypeOf(m1).Elem().Kind() == reflect.Interface {
		s1, s2 := m1.([]interface{}), m2.([]interface{})
		if s1 == nil && s2 == nil {
			return true
		} else if (s1 == nil && s2 != nil) || (s1 != nil && s2 == nil) {
			return false
		}

		if len(s1) > len(s2) || len(s1) < len(s2) {
			return false
		}

		for i, v := range s1 {
			if res := compareInterface(v, s2[i]); res {
				return true
			} else {
				return false
			}
		}
	}

	if t1.Kind() == reflect.Slice && reflect.TypeOf(m1).Elem().Kind() == reflect.String {
		s1, s2 := m1.([]string), m2.([]string)
		if s1 == nil && s2 == nil {
			return true
		} else if (s1 == nil && s2 != nil) || (s1 != nil && s2 == nil) {
			return false
		}

		if len(s1) > len(s2) || len(s1) < len(s2) {
			return false
		}

		for i, v := range s1 {
			return compareNotNull(v, s2[i])
		}
	}
	fmt.Println("未知类型, 请检查，先将比对结果设为false.")
	return false
}

// 比较基线和标准基线
func compareResource(curFile, stdFile string) []string {
	var resList []string
	curResource, stdResource := json2Map(curFile), json2Map(stdFile)
	checkFilesExist(curFile, stdFile)
	for k1, v1 := range curResource {
		if v2, ok := stdResource[k1]; ok {
			if res := compareInterface(v1, v2); !res {
				resList = append(resList, k1)
			}
		} else {
			resList = append(resList, k1)
		}
	}

	for k2, _ := range stdResource {
		if _, ok := curResource[k2]; !ok {
			resList = append(resList, k2)
		}
	}
	return resList
}

func compareNotNull(a, b interface{}) bool {
	if (a != nil && b == nil) || (a == nil && b != nil) {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if res := compareInterface(a, b); !res {
		return false
	}
	return true
}
