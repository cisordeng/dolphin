package reply

import (
	"github.com/cisordeng/beego/xenon"

	bReply "dolphin/business/reply"
	"dolphin/common/leo"
)

type Reply struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Reply))
}

func (this *Reply) Resource() string {
	return "reply.reply"
}

func (this *Reply) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{
			"id",
		},
		"PUT":  []string{
			"token",
			"resource_type",
			"resource_id",
			"reply_id",
			"content",
		},
	}
}

func (this *Reply) Get() {
	id, _ := this.GetInt("id", 0)

	reply := bReply.GetReplyById(id)
	data := bReply.EncodeReply(reply)
	this.ReturnJSON(data)
}

func (this *Reply) Put() {
	resourceType := this.GetString("resource_type")
	resourceId, _ := this.GetInt("resource_id", 0)
	replyId, _ := this.GetInt("reply_id", 0)
	content := this.GetString("content")

	user := leo.User{}
	this.GetUserFromToken(&user)

	iReply := bReply.NewReply(user, resourceType, resourceId, replyId, content)
	bReply.Fill([]*bReply.Reply{ iReply })
	data := bReply.EncodeReply(iReply)
	this.ReturnJSON(data)
}
