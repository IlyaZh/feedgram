package message_sender

import "errors"

var ErrFormatterNotFound = errors.New("formatter not found")
var ErrMessageFormatFailed = errors.New("message format has failed")
