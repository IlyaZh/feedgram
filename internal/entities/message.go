package entities

type MessageType int

const (
	MESSAGE_TYPE_RAW MessageType = iota
	MESSAGE_TYPE_LINK
	MESSAGE_TYPE_COMMAND
	MESSAGE_TYPE_POST
)

type Message struct {
	Type    MessageType
	Content *string
	Link    *Link
	Command *Command
}

func NewMessageLink(link Link) Message {
	return Message{
		Type: MESSAGE_TYPE_LINK,
		Link: &link,
	}
}

func NewMessageCommand(command Command) Message {
	return Message{
		Type:    MESSAGE_TYPE_COMMAND,
		Command: &command,
	}
}
