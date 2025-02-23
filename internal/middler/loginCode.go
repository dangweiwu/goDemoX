package middler

/*校验code是否有效 无效则退出登陆
1. 放在token中间件之后
2. 必须有redis
*/
import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goDemoX/internal/app/my/mymodel"
	"goDemoX/internal/ctx"
	"goDemoX/internal/pkg/api/hd"
	"goDemoX/internal/pkg/jwtx"
)

func LoginCodeErrResponse(c *gin.Context, data string) {
	c.JSON(401, hd.ErrMsg(data, "请重新登陆"))
	c.Abort()
}

func LoginCode(appctx *ctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, err := jwtx.GetCode(c)
		if err != nil {
			LoginCodeErrResponse(c, err.Error()+":code")
			return
		}
		uid, err := jwtx.GetUid(c)
		if err != nil {
			LoginCodeErrResponse(c, err.Error()+":jwt_get_id")
			return
		}
		//fmt.Println("@@", adminmodel.GetAdminRedisLoginId(int(uid)))
		logincode, err := appctx.Redis.Get(context.Background(), mymodel.GetAdminRedisLoginId(int(uid))).Result()
		if err != nil {
			if err == redis.Nil {
				LoginCodeErrResponse(c, err.Error()+":code")
			} else {
				LoginCodeErrResponse(c, err.Error())
			}
			return
		}
		//fmt.Println("@@", logincode)

		if logincode != code {
			LoginCodeErrResponse(c, "invalid_login_code")
			return
		}
		c.Next()
	}
}
