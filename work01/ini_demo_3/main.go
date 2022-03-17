package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0.参数的校验
	// 传进来的data必须是指针类型
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be pointer") //格式化输出之后返回一个error类型
		return
	}
	// 传进来的data参数必须是一个结构体类型指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer") //格式化输出之后返回一个error类型
		return
	}
	// 1.读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//将字节类型的文件内容转化为字符串
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	// 2.一行一行的读
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[]表示是节 section
		if strings.HasPrefix(line, "[") {
			if line[0] == '[' && line[len(line)-1] == ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//把这一行首尾的[]去掉，取到中间的内容并把首尾的空格去掉
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[]，就是=分割的键值对

		}

	}

	return

}

func main() {
	var mc MysqlConfig
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed,err:%v", err)
		return
	}
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
