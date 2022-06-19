package svc

import (
	"github.com/betalixt/eventSourceUsers/util/blerr"
	"github.com/betalixt/eventSourceUsers/util/txcontext"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (tsvc *TokenService) ValidateToken(
	tctx *txcontext.TransactionContext,
	tknStr string,
) (*jwt.MapClaims, *blerr.Error) {
	return nil, nil
}
