// Code generated by "stringer -type=Type"; DO NOT EDIT

package message

import "fmt"

const _Type_name = "TypeDiscoverTypeDiscoverReplyTypeRegisterTypeUnregister"

var _Type_index = [...]uint8{0, 12, 29, 41, 55}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
