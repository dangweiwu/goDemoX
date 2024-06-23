package api

import (
	"DEMOX_ADMINAUTH/internal/app/admin/adminmodel"
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"DEMOX_ADMINAUTH/internal/ctx"
	"DEMOX_ADMINAUTH/internal/pkg"
	"DEMOX_ADMINAUTH/internal/pkg/api/hd"
	"DEMOX_ADMINAUTH/internal/router"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

/*
重置账号密码
*/

type ResetPassword struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewResetPassword(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &ResetPassword{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| admin | 3 | 修改密码
// @path 	| /api/admin/resetpwd/:id
// @method 	| PUT
// @urlparam |n id |d 用户ID |v required |t int    |e 1
// @header   |n Authorization |d token  |t string |c 鉴权
// @tbtitle  | 200 Response
// @tbrow    |n data |d 新密码 |c 数字与字母组合的随机6位密码 |t string
func (this *ResetPassword) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &adminmodel.AdminPo{}
	po.ID = id
	pwd, err := this.ResetPassword(po)
	if err != nil {
		return err
	}
	this.Rep(hd.Response{pwd})
	return nil
}

func (this *ResetPassword) ResetPassword(rawPo *adminmodel.AdminPo) (string, error) {
	id, err := this.appctx.GetUid(this.ctx)
	if err != nil {
		return "", err
	}
	if id == rawPo.ID {
		return "", errors.New("不能重置自己密码")
	}

	po := &adminmodel.AdminPo{}
	if r := this.appctx.Db.Model(po).Where("id=?", rawPo.ID).Take(po); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return "", errors.New("记录不存在")
		} else {
			return "", r.Error
		}
	}

	_pwd := this.newPwd()
	pwd := pkg.GetPassword(_pwd)
	r := this.appctx.Db.Model(po).Update("password", pwd)

	//踢出在线
	this.appctx.Redis.Del(context.Background(), mymodel.GetAdminRedisLoginId(int(po.ID)))

	return _pwd, r.Error
}

// 生成三位字符，三位数字的密码
var ltes = "abcdefghijkmnpqrstuvwxyz"
var nums = "0123456789"

func (this *ResetPassword) newPwd() string {

	rt := ""
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
		rt += string(ltes[rand.Intn(len(ltes))])
	}
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
		rt += string(nums[rand.Intn(len(nums))])
	}
	return rt
}
