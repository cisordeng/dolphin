package reply

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mReply "dolphin/model/reply"
)

func GetOneReply(filters xenon.Map) *Reply {
	o := orm.NewOrm()
	qs := o.QueryTable(&mReply.Reply{})

	var model mReply.Reply
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	xenon.PanicNotNilError(err, "raise:comment:not_exits", "comment不存在")
	return InitReplyFromModel(&model)
}

func GetReplies(filters xenon.Map, orderExprs ...string ) []*Reply {
	o := orm.NewOrm()
	qs := o.QueryTable(&mReply.Reply{})

	var models []*mReply.Reply
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	replies := make([]*Reply, 0)
	for _, model := range models {
		replies = append(replies, InitReplyFromModel(model))
	}
	return replies
}

func GetPagedReplies(page *xenon.Paginator, filters xenon.Map, orderExprs ...string ) ([]*Reply, xenon.PageInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mReply.Reply{})

	var models []*mReply.Reply
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	pageInfo, err := xenon.Paginate(qs, page, &models)
	xenon.PanicNotNilError(err)

	replies := make([]*Reply, 0)
	for _, model := range models {
		replies = append(replies, InitReplyFromModel(model))
	}
	return replies, pageInfo
}

func GetReplyById(id int) *Reply {
	return GetOneReply(xenon.Map{
		"id": id,
	})
}
