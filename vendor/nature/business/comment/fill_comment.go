package comment

import (
	"github.com/cisordeng/beego/xenon"

	"nature/common/leo"
)

func Fill(comments []*Comment) {
	if len(comments) == 0 || comments[0] == nil {
		return
	}

	fillUser(comments)
	fillComment(comments)
}

func fillUser(comments []*Comment) {
	userIds := make([]int, 0)
	for _, Comment := range comments {
		userIds = append(userIds, Comment.UserId)
	}

	users := leo.GetUsers(xenon.Map{
		"id__in": userIds,
	})

	id2user := make(map[int]*leo.User)
	for _, user := range users {
		id2user[user.Id] = user
	}

	for _, Comment := range comments {
		if user, ok := id2user[Comment.UserId]; ok {
			Comment.User = user
		}
	}
	return
}

func fillComment(comments []*Comment) {
	commentIds := make([]int, 0)
	for _, comment := range comments {
		commentIds = append(commentIds, comment.CommentId)
	}

	iComments := GetComments(xenon.Map{
		"id__in": commentIds,
	})

	id2comment := make(map[int]*Comment)
	for _, comment := range iComments {
		id2comment[comment.Id] = comment
	}

	for _, comment := range comments {
		if comment, ok := id2comment[comment.CommentId]; ok {
			comment.Comment = comment
		}
	}
	return
}
