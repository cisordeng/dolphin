package reply

import (
	"github.com/cisordeng/beego/xenon"

	"dolphin/common/leo"
)

func Fill(replies []*Reply) {
	if len(replies) == 0 || replies[0] == nil {
		return
	}

	fillUser(replies)
	fillReply(replies)
}

func fillUser(replies []*Reply) {
	userIds := make([]int, 0)
	for _, Reply := range replies {
		userIds = append(userIds, Reply.UserId)
	}

	users := leo.GetUsers(xenon.Map{
		"id__in": userIds,
	})

	id2user := make(map[int]*leo.User)
	for _, user := range users {
		id2user[user.Id] = user
	}

	for _, Reply := range replies {
		if user, ok := id2user[Reply.UserId]; ok {
			Reply.User = user
		}
	}
	return
}

func fillReply(replies []*Reply) {
	replyIds := make([]int, 0)
	for _, reply := range replies {
		replyIds = append(replyIds, reply.ReplyId)
	}

	iReplies := GetReplies(xenon.Map{
		"id__in": replyIds,
	})

	id2reply := make(map[int]*Reply)
	for _, reply := range iReplies {
		id2reply[reply.Id] = reply
	}

	for _, reply := range replies {
		if reply, ok := id2reply[reply.ReplyId]; ok {
			reply.Reply = reply
		}
	}
	return
}
