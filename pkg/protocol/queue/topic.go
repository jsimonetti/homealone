package queue

import (
	"bytes"
	"strings"
)

//go:generate stringer -type=Topic

// Topic is a topic used on the hub
type Topic uint8

// Different topics used by this network.
const (
	Inventory Topic = iota
	Command
	Event
)

// Valid return true if the Topic is known
func (i Topic) Valid() bool {
	if i >= Topic(len(_Topic_index)-1) {
		return false
	}
	return true
}

// ValidTopic return true if the Topic i is known
func ValidTopic(i Topic) bool {
	if i >= Topic(len(_Topic_index)-1) {
		return false
	}
	return true
}

// AllTopics is a shorthand to return all known Topics.
// This is mostly used for application that want to see everything
func AllTopics() (t []Topic) {
	for i := 0; i < len(_Topic_index); i++ {
		t = append(t, Topic(i))
	}
	return
}

// GetTopic will return the topic for a topic string
func GetTopic(topic string) Topic {
	c := strings.Index(_Topic_name, topic)
	i := bytes.IndexByte(_Topic_index[:], uint8(c))
	return Topic(i)
}
