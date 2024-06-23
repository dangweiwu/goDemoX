package myserver

import (
	"DEMOX_ADMINAUTH/internal/app/my/mymodel"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"strings"
)

func NewLogCode(userid int64, rd *redis.Client) (logincode string, err error) {
	//登陆处理
	//登陆code 控制唯一登陆有效及踢人
	if logincode = uuid.New().String(); logincode == "" {
		err = errors.New("logincode is empty")
		return
	} else {
		logincode = strings.Split(logincode, "-")[0]
		if r := rd.Set(context.Background(), mymodel.GetAdminRedisLoginId(int(userid)), logincode, 0); r.Err() != nil {
			err = errors.New("login code redis:" + r.Err().Error())
			return
		}
	}
	return logincode, nil

}
