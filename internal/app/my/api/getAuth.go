package api

import (
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/app/role/rolemodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
)

//获取所有权限

type MyAuth struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewMyAuth(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &MyAuth{hd.NewHd(c), c, appctx}
}

// Do
// @api | me | 7 | 获取权限
// @path | /api/my-auth
// @method | GET
// @header |n Authorization |d token |t string |c 鉴权
// @response | hd.Response | 200 Response
// @tbtitle | 200 Response
// @tbrow |n data |d 角色 |t []string |c 角色数组
// @response | hd.ErrResponse | 400 Response
// @tbtitle  | msg 数据
// @tbrow |n msg |e 角色已被禁用
// @tbrow |n msg |e 角色不存在
func (this *MyAuth) Do() error {

	roleCode, err := this.appctx.GetRole(this.ctx)
	if err != nil {
		return err
	}
	rolePo := &rolemodel.RolePo{}
	r, err := this.appctx.Redis.Get(context.Background(), mymodel.ROLE_STATUS+roleCode).Result()
	if err == redis.Nil {
		//不存在从数据库里捞数据

		if r := this.appctx.Db.Model(rolePo).Where("code=?", roleCode).Take(rolePo); r.Error != nil {
			if r.Error == gorm.ErrRecordNotFound {
				return errors.New("角色不存在")
			}
			return r.Error
		}
		this.appctx.Redis.Set(context.Background(), mymodel.ROLE_STATUS+roleCode, rolePo.Status, time.Hour*24*7)
		if rolePo.Status == "0" {
			this.ctx.JSON(401, map[string]string{"data": "角色已被禁用"})
			this.ctx.Abort()
			return nil
		}
	} else {
		if r == "0" {
			this.ctx.JSON(401, map[string]string{"data": "角色已被禁用"})
			this.ctx.Abort()
			return nil
		}
	}
	authstring, err := this.appctx.Redis.Get(context.Background(), mymodel.ROLE_AUTH+roleCode).Result()
	if err == redis.Nil {
		if rolePo == nil {
			if r := this.appctx.Db.Model(rolePo).Where("code=?", roleCode).Take(rolePo); r.Error != nil {
				if r.Error == gorm.ErrRecordNotFound {
					return errors.New("该角色不存在")
				}
				return r.Error
			}
		}

		authstr := ""
		if r, err := json.Marshal(rolePo.Auth); err != nil {
			authstr = "[]"
		} else {
			authstr = string(r)
		}
		this.appctx.Redis.Set(context.Background(), mymodel.ROLE_AUTH+roleCode, authstr, time.Hour*24*7)
		hd.Rep(this.ctx, hd.Response{rolePo.Auth})
		return nil
	}
	rd := []string{}
	if err := json.Unmarshal([]byte(authstring), &rd); err != nil {
		return err
	}
	hd.Rep(this.ctx, hd.Response{rd})

	return nil
}
