package utils

import (
	"os"
	"errors"
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
func MakeDir(dirPath string) (err error){
	if 0 == len(dirPath){
		return errors.New("dirPath length is 0")
	}

	dirPath = strings.Replace(dirPath,"\\","/",-1)
	pathAry := strings.Split(dirPath,"/")
	index := strings.Index(dirPath,".")
	pathLen := len(pathAry)
	if index > -1{
		pathLen -= 1
	}
	for i:=1; i < pathLen; i++{
		curPath := strings.Join(pathAry[:i],"/")
		if CheckFileIsExist(curPath){
			continue
		}
		os.MkdirAll(curPath,0755)
	}
	if index >-1{
		_,err = os.Create(dirPath)
	}
	return
}