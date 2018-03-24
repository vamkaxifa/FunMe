package controllers

import (
	"github.com/revel/revel"
	"FunMe/app/model/code"
	"fmt"
)

type App struct {
	*revel.Controller
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
	path := "/Users/vivian/Documents/workfile/gopath/src/FunMe/app/controllers/app.go"
	fmt.Println(path)
	res := code.StatisticsCommentLine(path)
	return c.RenderJSON(res)
}


