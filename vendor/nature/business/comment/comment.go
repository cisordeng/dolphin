package comment

import (
	"time"

	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	"nature/common/leo"
	mComment "nature/model/comment"
)

type Comment struct {
	Id int
	UserId int
	ResourceType string
	ResourceId int
	CommentId int
	Content string
	CreatedAt time.Time

	User *leo.User
	Comment *Comment
}

func init() {
}

func InitCommentFromModel(model *mComment.Comment) *Comment {
	instance := new(Comment)
	instance.Id = model.Id
	instance.UserId = model.UserId
	instance.ResourceType = model.ResourceType
	instance.ResourceId = model.ResourceId
	instance.CommentId = model.CommentId
	instance.Content = model.Content
	instance.CreatedAt = model.CreatedAt

	return instance
}

func NewComment(user leo.User, restResource string, resourceId int, commentId int, content string) *Comment {
	model := mComment.Comment{
		UserId: user.Id,
		ResourceType: restResource,
		ResourceId: resourceId,
		CommentId: commentId,
		Content: content,
	}
	_, err := orm.NewOrm().Insert(&model)
	xenon.PanicNotNilError(err)
	return InitCommentFromModel(&model)
}
