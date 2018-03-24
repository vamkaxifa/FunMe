package controllers

import (
	"github.com/revel/revel"
	"FunMe/app/model/code"
	"fmt"
)

type App struct {
	*revel.Controller
}

type CommonResponse struct {
	Code 	int64 `json:"code"`
	Data    code.StatisticsResult `json:"data"`
	ErrorDes string		`json:"error_des"`
}

/**
index page
 */
func (c App) Index() revel.Result {
	return c.Render()
}


/**
health check
 */
func (c App) Healthcheck() revel.Result{
	return c.Render()
}

/**
Statistics Number of comment lines
 */

func (c App) Code() revel.Result{
	path := "/Users/vivian/Documents/workfile/gopath/src/FunMe"
	fmt.Println(path)
	suffixAry := []string{"go","java","py"}
	res,err := code.Statistics(path,suffixAry)
	if err!=nil{
		return c.RenderJSON(CommonResponse{Code:1,ErrorDes:err.Error()})
	}
	return c.RenderJSON(CommonResponse{Code:0,Data:*res})
}


