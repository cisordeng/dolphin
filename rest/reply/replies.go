package reply

import (
	"github.com/cisordeng/beego/xenon"

	bReply "dolphin/business/reply"
)

type Replies struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Replies))
}

func (this *Replies) Resource() string {
	return "reply.replies"
}

func (this *Replies) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

func (this *Replies) Get() {
	page := this.GetPage()
	filters := this.GetFilters()
	orders := this.GetOrders()
	replies, pageInfo := bReply.GetPagedReplies(page, filters, orders...)
	bReply.Fill(replies)
	data := bReply.EncodeManyReply(replies)
	this.ReturnJSON(xenon.Map{
		"replies": data,
		"page_info": pageInfo.ToMap(),
	})
}

