/*
 * @Description:基础控制器
 * @Author: gphper
 * @Date: 2021-05-18 23:21:06
 */

package controllers

import (
	"errors"
	"grm/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const AJAXSUCCESS int = 1
const AJAXFAIL int = 0

type BaseController struct {
}

func (Base *BaseController) Success(c *gin.Context, url string, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"msg":         message,
		"url":         url,
		"iframe_jump": false,
	})
}

func (Base *BaseController) Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": false,
		"msg":    message,
	})
}

func (Base *BaseController) AjaxReturn(c *gin.Context, code int, obj interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   obj,
	})
}

func (Base BaseController) ErrorHtml(c *gin.Context, err error) {

	c.HTML(http.StatusOK, "toast/error.html", gin.H{
		"title": "ERROR",
		"msg":   err.Error(),
	})

}

func (Base *BaseController) FormBind(c *gin.Context, obj interface{}) error {

	trans, err := common.InitTrans("zh")

	if err != nil {
		return err
	}

	if err := c.ShouldBind(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok && errs != nil {
			return errs
		}

		for _, v := range errs.Translate(trans) {
			return errors.New(v)
		}

	}
	return nil
}
