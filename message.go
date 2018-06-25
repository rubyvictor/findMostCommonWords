package main

import s "strings"

type Message struct {
	Content string
	Errors  map[string]string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	if s.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Please enter some text."
	}
	return len(msg.Errors) == 0
}
