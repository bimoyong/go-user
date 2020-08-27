package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/config"

	"github.com/micro/go-micro/v2/auth"
	log "github.com/micro/go-micro/v2/logger"

	proto "gitlab.com/bimoyong/go-user/proto/user"
)

// Login function
func (u *User) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginRsp) (err error) {
	log.Debugf("[User] Request: req=[%+v]", req)

	var acc *auth.Account
	if acc, err = auth.DefaultAuth.Generate(req.Id); err != nil {
		log.Errorf("[User] Login failure!: err=[%s], id=[%s]", err.Error(), req.Id)

		return
	}
	log.Infof("[User] Login success: id=[%s] scope=[%+v]", acc.ID, acc.Scopes)
	exp := config.Get("auth", "expiry").Int(0)
	token, _ := auth.DefaultAuth.Token(
		auth.WithToken(acc.Secret),
		auth.WithExpiry(time.Second*time.Duration(exp)),
	)

	rsp.Token = token.AccessToken
	rsp.Refresh = token.RefreshToken
	rsp.Created = token.Created.Unix()
	rsp.Expiry = token.Expiry.Unix()

	return
}
