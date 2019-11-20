package reply
import (
	"github.com/cisordeng/beego/xenon"

	"dolphin/common/leo"
)

func EncodeReply(reply *Reply) xenon.Map {
	if reply == nil {
		return nil
	}

	rUser := leo.EncodeUser(reply.User)
	rReply := EncodeReply(reply.Reply)

	mapReply := xenon.Map{
		"id": reply.Id,
		"user": rUser,
		"reply": rReply,
		"resource_id": reply.ResourceId,
		"resource_type": reply.ResourceType,
		"content": reply.Content,
		"created_at": reply.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapReply
}


func EncodeManyReply(replies []*Reply) []xenon.Map {
	mapReplies := make([]xenon.Map, 0)
	for _, reply := range replies {
		mapReplies = append(mapReplies, EncodeReply(reply))
	}
	return mapReplies
}
