package entities

type MessageType int

const (
	MESSAGE_TYPE_RAW MessageType = iota
	MESSAGE_TYPE_LINK
)

type Message struct {
	Type    MessageType
	Content *string
	Link    *Link
}

func NewMessageLink(link Link) Message {
	return Message{
		Type: MESSAGE_TYPE_LINK,
		Link: &link,
	}
}
