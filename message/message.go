package message

import "time"

type Message struct {
	author    string
	content   string
	timestamp time.Time
}

func (m *Message) Marshal() []byte {
	out := make([]byte, 0)
	/*
		authorBytes := []byte(m.author)
		contentBytes := []byte(m.content)
		tsBytes := []byte(m.timestamp.String())

	*/

	return out
}

func sliceLenToBytes(in []byte) []byte {
	//	inlen := uint32(len(in))
	return make([]byte, 0)
}
