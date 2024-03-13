package safeobject

import "golang-api-template/internal/domain/entity"

type PairToken struct {
	Access  string
	Refresh *entity.RefreshToken
}

func NewPairToken(access string, refresh *entity.RefreshToken) *PairToken {
	return &PairToken{Access: access, Refresh: refresh}
}
