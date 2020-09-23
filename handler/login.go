package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/auth/token"
	"github.com/micro/go-micro/v2/auth/token/jwt"
	"github.com/micro/go-micro/v2/config"
	log "github.com/micro/go-micro/v2/logger"

	proto "gitlab.com/bimoyong/go-user/proto/user"
)

// Login function
func (u *User) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginRsp) (err error) {
	log.Debugf("[User] Request: req=[%+v]", req)

	acc := &auth.Account{ID: req.Id}
	privKey := config.Get("auth", "private_key").String("")
	pubKey := config.Get("auth", "public_key").String("")
	exp := config.Get("auth", "expiry").Int(0)

	prov := jwt.NewTokenProvider(
		token.WithPrivateKey(privKey),
		token.WithPublicKey(pubKey),
	)

	var tk *token.Token
	if tk, err = prov.Generate(
		acc,
		token.WithExpiry(time.Second*time.Duration(exp)),
	); err != nil {
		log.Errorf("[User] Login failure!: err=[%s], id=[%s]", err.Error(), req.Id)

		return
	}
	log.Infof("[User] Login success: id=[%s] scope=[%+v]", acc.ID, acc.Scopes)

	rsp.Token = tk.Token
	// rsp.Refresh = tk.RefreshToken
	rsp.Created = tk.Created.Unix()
	rsp.Expiry = tk.Expiry.Unix()

	return
}
