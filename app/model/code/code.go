package code

import (
	"FunMe/app/model/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"
)

type StatisticsResult struct {
	TotalNum   int64 `json:"total_num"`   //代码总行数
	CodeNum    int64 `json:"code_num"`    //代码行数
	CommentNum int64 `json:"comment_num"` //注释行数
}

/**
统计结果chan
*/
var statisticsResultChan chan StatisticsResult

/**
Statistics Number of comment lines
*/
func StatisticsCommentLine(path string) {
	if utils.CheckFileIsExist(path) == false {
		fmt.Println("path doesn't exist!")
		statisticsResultChan <- StatisticsResult{TotalNum: 0, CodeNum: 0, CommentNum: 0}
	}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		statisticsResultChan <- StatisticsResult{TotalNum: 0, CodeNum: 0, CommentNum: 0}
	}
	rd := bufio.NewReader(f)
	rs := StatisticsResult{TotalNum: 0, CodeNum: 0, CommentNum: 0}
	ifMulty := false
	for {
		line, err := rd.ReadString('\n')
		if io.EOF == err {
			break
		}
		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
			continue //当前行失败，忽略，go on  //FIXME think about 多行注释 /**/
		}

		line = strings.Trim(line, " ") //去掉首尾空格

		if len(line) == 0 || line == "\n" || line == "\r" { //空行不统计
			continue
		}

		if ifMulty { //是否处于多行注释统计过程中
			rs.TotalNum += 1
			rs.CommentNum += 1
			if strings.Contains(line, "*/") {
				ifMulty = false
			}
			continue
		}

		if strings.HasPrefix(line, "//") { // //类型注释
			rs.TotalNum += 1
			rs.CommentNum += 1
			continue
		}

		if strings.HasPrefix(line, "/*") { // /* 类型注释
			rs.TotalNum += 1
			rs.CommentNum += 1
			ifMulty = true
			continue
		}

		rs.TotalNum += 1
		rs.CodeNum += 1
	}
	jsonbyte, err := json.Marshal(rs)
	fmt.Println(string(jsonbyte))
	statisticsResultChan <- rs
}

/*
统计给定文件夹下指定文件类型里代码行数、注释行数
*/
func Statistics(path string, suffixAry []string) (*StatisticsResult, error) {
	var fiAry []string
	if utils.IsDir(path) {
		ary, err := utils.WalkDir(path, suffixAry)
		if err != nil {
			return nil, err
		}
		fiAry = ary
	} else {
		fiAry = []string{path}
	}
	fmt.Println(len(fiAry))
	statisticsResultChan = make(chan StatisticsResult, len(fiAry)) //FIXME length of cache doesn't need that too much
	for _, fi := range fiAry {
		go StatisticsCommentLine(fi)
	}
	statisticsResult := StatisticsResult{}

	for i := 0; i < len(fiAry); i++ {
		sr := <-statisticsResultChan
		statisticsResult.TotalNum += sr.TotalNum
		statisticsResult.CodeNum += sr.CodeNum
		statisticsResult.CommentNum += sr.CommentNum
	}

	return &statisticsResult, nil
}
