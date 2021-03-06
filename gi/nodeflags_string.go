// Code generated by "stringer -type=NodeFlags"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoLayout-16]
	_ = x[EventsConnected-17]
	_ = x[CanFocus-18]
	_ = x[HasFocus-19]
	_ = x[FullReRender-20]
	_ = x[ReRenderAnchor-21]
	_ = x[Invisible-22]
	_ = x[Inactive-23]
	_ = x[Selected-24]
	_ = x[MouseHasEntered-25]
	_ = x[DNDHasEntered-26]
	_ = x[NodeDragging-27]
	_ = x[InstaDrag-28]
	_ = x[NodeFlagsN-29]
	_ = x[TextFieldFocusActive-29]
}

const _NodeFlags_name = "NoLayoutEventsConnectedCanFocusHasFocusFullReRenderReRenderAnchorInvisibleInactiveSelectedMouseHasEnteredDNDHasEnteredNodeDraggingInstaDragNodeFlagsN"

var _NodeFlags_index = [...]uint8{0, 8, 23, 31, 39, 51, 65, 74, 82, 90, 105, 118, 130, 139, 149}

func (i NodeFlags) String() string {
	i -= 16
	if i < 0 || i >= NodeFlags(len(_NodeFlags_index)-1) {
		return "NodeFlags(" + strconv.FormatInt(int64(i+16), 10) + ")"
	}
	return _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]]
}

func StringToNodeFlags(s string) (NodeFlags, error) {
	for i := 0; i < len(_NodeFlags_index)-1; i++ {
		if s == _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]] {
			return NodeFlags(i + 16), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: NodeFlags")
}
