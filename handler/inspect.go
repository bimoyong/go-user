package handler

import (
	"context"
	"strings"

	aproto "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/auth"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"

	proto "gitlab.com/bimoyong/go-user/proto/user"
)

// Inspect function
func (u *User) Inspect(ctx context.Context, req *aproto.Request, rsp *proto.InspectRsp) (err error) {
	authz, ok := metadata.Get(ctx, "Authorization")
	if !ok {
		h, _ := metadata.FromContext(ctx)
		log.Warnf("[User] Anonymous request: header=[%+v]", h)

		return
	}
	log.Debugf("[User] Request: authz=[%s]", authz)
	token := strings.TrimPrefix(authz, auth.BearerScheme)

	var acc *auth.Account
	if acc, err = auth.DefaultAuth.Inspect(token); err != nil {
		log.Errorf("[User] Inspect authentication failure: token=[%s]", token)

		return
	}
	log.Debugf("[User] Valid account: token=[%s] acc=[%+v]", token, acc)

	rsp.Id = acc.ID
	rsp.Scopes = acc.Scopes

	return
}
