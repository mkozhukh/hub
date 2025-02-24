package hub

import (
	"fmt"
	"sort"
	"strings"
)

type (
	// Fields is a [key]value storage for Messages values.
	Fields map[string]interface{}

	// Message represent some message/event passed into the hub
	// It also contain some helper functions to convert the fields to primitive types.
	Message struct {
		Name   string
		Body   interface{}
		Fields Fields
	}
)

func NewMessage(name string, body interface{}) *Message {
	return &Message{
		Name:   name,
		Body:   body,
		Fields: make(Fields),
	}
}

// Topic return the message topic used when the message was sended.
func (m *Message) Topic() string {
	return m.Name
}

func (m *Message) String() string {
	return fmt.Sprintf("%s(%+v) %s", m.Name, m.Body, m.Fields.String())
}

func (f Fields) String() string {
	if len(f) == 0 {
		return "Fields(<empty>)"
	}

	fields := make([]string, 0, len(f))
	for k, v := range f {
		fields = append(fields, fmt.Sprintf("[%s]%v", k, v))
	}

	sort.Strings(fields)

	return "Fields( " + strings.Join(fields, ", ") + " )"
}
