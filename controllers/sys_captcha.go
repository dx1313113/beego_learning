package controllers

import (
	"beego_learning/global"
	"beego_learning/models"
	"context"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/mojocn/base64Captcha"
	"time"
)

type CaptchaController struct {
	beego.Controller
}

func (this *CaptchaController) URLMapping() {
	this.Mapping("Get", this.Get)
	this.Mapping("Post", this.Post)
}

// Get @Title Captcha
// @Description Get captcha
// @Success 200 {interface} interface
// @router /get [Get]
func (this *CaptchaController) Get() {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = models.Captcha{id, b64s}
	}
	this.ServeJSON()

}

// Post  @Title Captcha
// @Description verify captcha
// @Param data body models.Captcha true "verify captcha"
// @Success 200 {bool} bool
// @router /verify [post]
func (this *CaptchaController) Post() {
	var captcha models.Captcha
	json.Unmarshal(this.Ctx.Input.RequestBody, &captcha)
	if store.Verify(captcha.ID, captcha.B64s, true) {
		this.Data["json"] = true
	} else {
		this.Data["json"] = false
	}
	this.ServeJSON()
}

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// 配置RedisStore RedisStore实现base64Captcha.Store的接口
var store base64Captcha.Store = RedisStore{}

func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	//time.Minute*2：有效时间2分钟
	err := global.RedisDb.Set(ctx, key, value, time.Minute*2).Err()

	return err
}

func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := global.RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		//clear为true，验证通过，删除这个验证码
		err := global.RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}
