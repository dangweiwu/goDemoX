package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"goDemoX/internal/app/admin/adminmodel"
	"goDemoX/internal/app/my/mymodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/router"
	"gorm.io/gorm"
)

type AdminUpdate struct {
	*hd.Hd
	ctx    *gin.Context
	appctx *ctx.AppContext
}

func NewAdminUpdate(c *gin.Context, appctx *ctx.AppContext) router.IHandler {
	return &AdminUpdate{hd.NewHd(c), c, appctx}
}

// Do
// @api 	| admin | 2 | 修改用户
// @path 	| /api/admin/:id
// @method 	| PUT
// @urlparam |n id |d 用户ID |v required |t int    |e 1
// @header   |n Authorization |d token  |t string |c 鉴权
// @form     | adminmodel.AdminUpdateForm
// @tbtitle  | 200 Response
// @tbrow    |n data |e ok |c 成功 |t string
func (this *AdminUpdate) Do() error {
	var err error
	id, err := this.GetId()
	if err != nil {
		return err
	}
	po := &adminmodel.AdminUpdateForm{}
	err = this.Bind(po)
	if err != nil {
		return err
	}
	po.ID = id
	err = this.Update(po)
	if err != nil {
		return err
	}
	this.RepOk()
	return nil
}
func (this *AdminUpdate) Update(po *adminmodel.AdminUpdateForm) error {
	db := this.appctx.Db
	tmpPo := &adminmodel.AdminUpdateForm{}
	tmpPo.ID = po.ID
	if r := db.Model(tmpPo).Take(tmpPo); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return r.Error
		}
	}
	uid, err := this.appctx.GetUid(this.ctx)
	if err != nil {
		return err
	}

	if uid == po.ID {
		return errors.New("禁止修改自己")
	}

	//其他校验
	/*
		if err := this.Valid(po); err != nil {
			return err
		}
	*/

	//更新
	if r := db.Select("phone", "name", "status", "memo", "email", "is_super_admin", "updated_at").Updates(po); r.Error != nil {
		return r.Error
	}
	//修改人员下线
	if (tmpPo.Status == "1" && po.Status == "0") || tmpPo.Role != po.Role {
		this.appctx.Redis.Del(context.Background(), mymodel.GetAdminRedisLoginId(int(tmpPo.ID)))
	}

	return nil
}

//
//func (this *AdminUpdate) Valid(po *adminmodel.AdminUpdateForm) error {
//	var ct = int64(0)
//	if po.Phone != "" {
//		if r := this.appctx.Db.Model(po).Where("id != ? and phone = ? ", po.ID, po.Phone).Count(&ct); r.Error != nil {
//			return errs.WithMessage(r.Error, "校验失败")
//		} else if ct != 0 {
//			return errors.New("手机号已存在")
//		}
//	}
//
//	if po.Email != "" {
//		if r := this.appctx.Db.Model(po).Where("id != ? and email = ?", po.ID, po.Email).Count(&ct); r.Error != nil {
//			return errs.WithMessage(r.Error, "校验失败")
//		} else if ct != 0 {
//			return errors.New("Email已存在")
//		}
//	}
//
//	return nil
//}
