package admin

import (
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminLoginLogic {
	return AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req types.AdminLoginReq) (*types.AdminReply, error) {
	// todo: add your logic here and delete this line

	return &types.AdminReply{}, nil
}

//func (l *AdminLoginLogic) getAdminJwtToken(secretKey string, iat, seconds int64, userInfo model.SystemUser) (string, error) {
//	claims := make(jwt.MapClaims)
//	claims["exp"] = iat + seconds
//	claims["iat"] = iat
//	claims["userId"] = userInfo.Id
//	claims["roleId"] = userInfo.RoleId
//	token := jwt.New(jwt.SigningMethodHS256)
//	token.Claims = claims
//	return token.SignedString([]byte(secretKey))
//}
