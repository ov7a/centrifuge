package prepared

import (
	"sync"

	"github.com/centrifugal/protocol"
)

// Reply is structure for encoding reply only once.
type Reply struct {
	ProtoType protocol.Type
	Reply     *protocol.Reply
	data      []byte
	pushData  []byte
	once      sync.Once
	pushOnce  sync.Once
}

// NewReply initializes Reply.
func NewReply(reply *protocol.Reply, protoType protocol.Type) *Reply {
	return &Reply{
		Reply:     reply,
		ProtoType: protoType,
	}
}

// Data returns data associated with reply which is only calculated once.
func (r *Reply) Data() []byte {
	r.once.Do(func() {
		encoder := protocol.GetReplyEncoder(r.ProtoType)
		data, _ := encoder.Encode(r.Reply)
		r.data = data
	})
	return r.data
}

// Data returns data associated with reply which is only calculated once.
func (r *Reply) PushData() []byte {
	r.pushOnce.Do(func() {
		encoder := protocol.GetPushEncoder(r.ProtoType)
		pushData, _ := encoder.Encode(r.Reply.Push)
		r.pushData = pushData
	})
	return r.pushData
}
