// Code generated by "stringer -type=Topic"; DO NOT EDIT

package queue

import "fmt"

const _Topic_name = "InventoryCommandEvent"

var _Topic_index = [...]uint8{0, 9, 16, 21}

func (i Topic) String() string {
	if i >= Topic(len(_Topic_index)-1) {
		return fmt.Sprintf("Topic(%d)", i)
	}
	return _Topic_name[_Topic_index[i]:_Topic_index[i+1]]
}
