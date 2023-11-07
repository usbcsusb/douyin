package repo

import (
	"context"
)

type RCacheRepo interface {
	KeyExist(c context.Context, key string) (bool, error)
	SAddFollow(c context.Context, key string, targetId int64) error
	SAddManyIds(c context.Context, key string, followIds []int64) error
	SAddFriend(c context.Context, key string, targetId int64) error
	DelFollow(c context.Context, key string) error
	DelFan(c context.Context, key string) error
	DelFriend(c context.Context, key string) error
	IsFollow(c context.Context, key string, targetId int64) (bool, error)
	SGetAllIds(c context.Context, key string) ([]int64, error)
	IsFriend(c context.Context, key string, targetId int64) (bool, error)
}
