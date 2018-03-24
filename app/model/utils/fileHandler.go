package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

/**
Check file is exist
*/
func CheckFileIsExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

/**
Make dir
*/
func MakeDir(dirPath string) (err error) {
	if 0 == len(dirPath) {
		return errors.New("dirPath length is 0")
	}

	dirPath = strings.Replace(dirPath, "\\", "/", -1)
	pathAry := strings.Split(dirPath, "/")
	index := strings.Index(dirPath, ".")
	pathLen := len(pathAry)
	if index > -1 {
		pathLen -= 1
	}
	for i := 1; i < pathLen; i++ {
		curPath := strings.Join(pathAry[:i], "/")
		if CheckFileIsExist(curPath) {
			continue
		}
		os.MkdirAll(curPath, 0755)
	}
	if index > -1 {
		_, err = os.Create(dirPath)
	}
	return
}

/**
Is dir
*/
func IsDir(fileName string) bool {
	f, err := os.Stat(fileName)
	if err != nil {
		debug.PrintStack()
		return false
	}
	return f.IsDir()
}

/**
获取指定目录及所有子目录下的所有文件，可以匹配多个后缀过滤
 */
func WalkDir(dirPth string, suffixAry []string) (files []string, err error) {
	files = make([]string,0)
	for i := range suffixAry {
		suffixAry[i] = strings.ToUpper(suffixAry[i]) //忽略大小写匹配
	}
	//遍历目录
	err = filepath.Walk(dirPth, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if hasSuffix(strings.ToUpper(fi.Name()), suffixAry) {
			files = append(files, fileName)
		}
		return nil
	})
	return files, err
}


func hasSuffix(str string, strAry []string) bool {
	for i := range strAry {
		if strings.HasSuffix(str, strAry[i]) {
			return true
		}
	}
	return false
}
