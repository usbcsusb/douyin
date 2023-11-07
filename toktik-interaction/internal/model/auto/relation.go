package auto

import "strconv"

type Relation struct {
	BaseModel
	UserId   uint `json:"user_id,string" gorm:"index:idx_relation,not null"`   // 用户ID
	TargetId uint `json:"target_id,string" gorm:"index:idx_relation,not null"` // 目标ID，添加复合索引
	//IsFriend int  `json:"is_friend" gorm:"not null"` // 如果需要保证 relation_id 唯一，可以使用该字段
}

func (*Relation) TableName() string {
	return "relation"
}

func CreateFollowKey(userId int64) string {
	return "follow::" + strconv.FormatInt(userId, 10)
}

func CreateFanKey(userId int64) string {
	return "fan::" + strconv.FormatInt(userId, 10)
}

func CreateFriendKey(userId int64) string {
	return "friend::" + strconv.FormatInt(userId, 10)
}
