package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"lmp/dao/mysql"
	"lmp/logic"
	"lmp/models"
)

// SignUPHandler 函数处理注册请求
func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil { // 只能检测请求的格式、类型对不对
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExit) {
			zap.L().Error("SignUp with User Exists", zap.Error(err))
			ResponseError(c, CodeUserExist)
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExit) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, token)
}
