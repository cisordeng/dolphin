package reply

import (
	"time"
	
	"github.com/cisordeng/beego/orm"
)

type Reply struct {
	Id int
	UserId int

	ResourceType string  // 资源类型
	ResourceId int // 资源ID

	ReplyId int // 回复ID
	Content string `orm:"type(text)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (o *Reply) TableName() string {
	return "reply_reply"
}

func init() {
	orm.RegisterModel(new(Reply))
}
