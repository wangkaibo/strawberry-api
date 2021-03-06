package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"strawberry-wallpaper/services"
	"strconv"
	"time"
)

type NoticeController struct {
	BaseController
	NoticeService services.NoticeService
}

func (c *NoticeController) Notice(ctx *gin.Context) {
	notices, err := c.NoticeService.GetNotices()
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}

func (c *NoticeController) NoticeList(ctx *gin.Context) {
	notices, err := c.NoticeService.GetNoticeList()
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}

func (c *NoticeController) DeleteNotice(ctx *gin.Context) {
	id := ctx.Param("id")
	notices, err := c.NoticeService.DeleteNotice(id)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}


func (c *NoticeController) PublishNotice(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	status := ctx.Request.FormValue("status")
	if status == "" {
		c.error(ctx, 400, "参数错误", gin.H{})
		return
	}
	isPublish, _ := strconv.Atoi(status)
	err := c.NoticeService.PublishNotice(id, isPublish)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, gin.H{})
	}
}

func (c *NoticeController) AddNotice(ctx *gin.Context) {
	content := ctx.Request.FormValue("content")
	id := ctx.Request.FormValue("id")
	if content == "" {
		c.error(ctx, 400, "内容不能为空", gin.H{})
		return
	}

	expireTimeStr := ctx.Request.FormValue("expire_time")
	var err error
	expireTime, err := stringToTime(expireTimeStr)
	if err != nil {
		c.error(ctx, 400, "参数错误", gin.H{})
		return
	}
	if id == "" {
		err = c.NoticeService.AddNotice(content, expireTime)
	} else {
		idInt,err := strconv.Atoi(id)
		if err != nil {
			c.error(ctx, 400, "参数错误", gin.H{})
			return
		}
		err = c.NoticeService.UpdateNotice(idInt, content, expireTime)
	}
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, gin.H{})
	}
}

func stringToTime(ms string) (time.Time, error) {
	msInt,err := strconv.Atoi(ms)
	if err != nil {
		return time.Time{}, err
	}
	nsImt64 := int64(msInt) * int64(time.Millisecond)
	return time.Unix(0, nsImt64), nil
}

func (c *NoticeController) Login(ctx *gin.Context) {
	inputUsername := ctx.Request.FormValue("username")
	inputPassword := ctx.Request.FormValue("password")
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	if inputUsername == "" || inputPassword == "" {
		c.error(ctx, 401, "帐号或密码错误", gin.H{})
		return
	}
	if !(inputUsername == username && inputPassword == password) {
		c.error(ctx, 401, "帐号或密码错误", gin.H{})
		return
	}
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerByte, _ := json.Marshal(header)
	headerBase64 := base64.StdEncoding.EncodeToString(headerByte)
	payload := map[string]interface{}{
		"username": inputUsername,
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + 86400 * 7,
	}
	payloadByte, _ := json.Marshal(payload)
	payloadBase64 := base64.StdEncoding.EncodeToString(payloadByte)
	sign := GetJwtSign(headerBase64, payloadBase64)
	token := headerBase64 + "." + payloadBase64 + "." + sign

	c.success(ctx, gin.H{
		"token": token,
	})
}

func GetJwtSign(header string, payload string) string {
	secret := os.Getenv("JWT_SECRET")
	signStr := header + "." + payload
	hmacHash256 := hmac.New(sha256.New, []byte(secret))
	hmacHash256.Write([]byte(signStr))
	b := hmacHash256.Sum(nil)

	return hex.EncodeToString(b)
}