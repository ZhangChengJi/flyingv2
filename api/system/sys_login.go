package system

import (
	"flyingv2/conf"
	"flyingv2/internal/core"
	"flyingv2/internal/core/model"
	"flyingv2/internal/core/resp"
	"flyingv2/logs"
	"flyingv2/pkg/etcd"
	"flyingv2/pkg/gin/plugin"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

type Handler struct {
	Store core.Interface
}

func NewHandler() core.RouteRegister {
	return &Handler{
		etcd.New(conf.AppPrefix),
	}
}
func (h *Handler) Router(r *gin.RouterGroup) {
	g := r.Group("sys")
	g.POST("login", h.Login)
	g.POST("register", h.Register)

}

// @Tags system
// @Summary 用户登录
// @Produce  application/json
// @Param data body model.Login true "用户名, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /sys/login [post]
func (h *Handler) Login(c *gin.Context) {
	var login model.Login
	_ = c.ShouldBind(&login)
	if value, err := h.Store.Get(context.Background(), login.Username); err != nil {
		logs.L.Error("login error: ", zap.Error(err))
		resp.FailWithMessage(fmt.Sprintf("login failed:%v", zap.Error(err)), c)
	} else {
		if len(value) > 0 {
			if ok := login.Verify(value); ok {
				h.loginNext(c, login.User)
			} else {
				logs.L.Error("username or password failed")
				resp.FailWithMessage("username or password failed", c)
			}
		}
	}
}

func (h *Handler) loginNext(c *gin.Context, user *model.User) {
	j := plugin.NewAuth()
	claims := jwt.StandardClaims{
		Subject:   user.Username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(j.ExpireTime)).Unix(),
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		logs.L.Error("create token error: ", zap.Error(err))
		resp.FailWithMessage("create token error", c)
		return
	} else {
		resp.OkWithDetailed(model.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.ExpiresAt * 1000,
		}, "login success", c)
		return
	}

}

// @Tags system
// @Summary 用户注册
// @Produce  application/json
// @Param data body model.Login true "用户名, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /sys/register [post]
func (h *Handler) Register(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	user.Salt = time.Now().Unix()
	u, _ := model.MarshalJSON(user)
	if err := h.Store.Set(context.Background(), user.Username, string(u)); err != nil {
		logs.L.Error("Register failed", zap.Error(err))
		resp.FailWithMessage(fmt.Sprintf(err.Error(), "username", user.Username), c)
	} else {
		resp.OkWithMessage("Register success", c)
	}

}
