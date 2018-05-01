// Code generated by "stringer -type=WindowFlags"; DO NOT EDIT.

package oswin

import (
	"fmt"
	"strconv"
)

const _WindowFlags_name = "DialogModalToolFullScreenWindowFlagsN"

var _WindowFlags_index = [...]uint8{0, 6, 11, 15, 25, 37}

func (i WindowFlags) String() string {
	if i < 0 || i >= WindowFlags(len(_WindowFlags_index)-1) {
		return "WindowFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _WindowFlags_name[_WindowFlags_index[i]:_WindowFlags_index[i+1]]
}

func (i *WindowFlags) FromString(s string) error {
	for j := 0; j < len(_WindowFlags_index)-1; j++ {
		if s == _WindowFlags_name[_WindowFlags_index[j]:_WindowFlags_index[j+1]] {
			*i = WindowFlags(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type WindowFlags", s)
}