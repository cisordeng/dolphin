package reply

import (
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	"dolphin/common/leo"
	mReply "dolphin/model/reply"
)

type Reply struct {
	Id int
	UserId int
	ResourceType string
	ResourceId int
	ReplyId int
	Content string
	CreatedAt time.Time

	User *leo.User
	Reply *Reply
}

func init() {
}

func InitReplyFromModel(model *mReply.Reply) *Reply {
	instance := new(Reply)
	instance.Id = model.Id
	instance.UserId = model.UserId
	instance.ResourceType = model.ResourceType
	instance.ResourceId = model.ResourceId
	instance.ReplyId = model.ReplyId
	instance.Content = model.Content
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewReply(user leo.User, restResource string, resourceId int, replyId int, content string) *Reply {
	model := mReply.Reply{
		UserId: user.Id,
		ResourceType: restResource,
		ResourceId: resourceId,
		ReplyId: replyId,
		Content: content,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitReplyFromModel(&model)
}
