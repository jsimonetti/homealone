// Code generated by "stringer -type=EventType"; DO NOT EDIT

package message

import "fmt"

const _EventType_name = "EventDriverStateChangeEventDeviceStateChangeEventComponentStateChangeEventComponentValueChange"

var _EventType_index = [...]uint8{0, 22, 44, 69, 94}

func (i EventType) String() string {
	if i >= EventType(len(_EventType_index)-1) {
		return fmt.Sprintf("EventType(%d)", i)
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}