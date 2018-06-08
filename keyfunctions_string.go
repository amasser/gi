// Code generated by "stringer -type=KeyFunctions"; DO NOT EDIT.

package gi

import (
	"fmt"
	"strconv"
)

const _KeyFunctions_name = "KeyFunNilKeyFunMoveUpKeyFunMoveDownKeyFunMoveRightKeyFunMoveLeftKeyFunPageUpKeyFunPageDownKeyFunPageRightKeyFunPageLeftKeyFunHomeKeyFunEndKeyFunFocusNextKeyFunFocusPrevKeyFunSelectItemKeyFunCancelSelectKeyFunSelectModeKeyFunSelectAllKeyFunAcceptKeyFunAbortKeyFunEditItemKeyFunCopyKeyFunCutKeyFunPasteKeyFunBackspaceKeyFunDeleteKeyFunKillKeyFunDuplicateKeyFunInsertKeyFunInsertAfterKeyFunGoGiEditorKeyFunShiftKeyFunCtrlKeyFunZoomOutKeyFunZoomInKeyFunPrefsKeyFunRefreshKeyFunctionsN"

var _KeyFunctions_index = [...]uint16{0, 9, 21, 35, 50, 64, 76, 90, 105, 119, 129, 138, 153, 168, 184, 202, 218, 233, 245, 256, 270, 280, 289, 300, 315, 327, 337, 352, 364, 381, 397, 408, 418, 431, 443, 454, 467, 480}

func (i KeyFunctions) String() string {
	if i < 0 || i >= KeyFunctions(len(_KeyFunctions_index)-1) {
		return "KeyFunctions(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _KeyFunctions_name[_KeyFunctions_index[i]:_KeyFunctions_index[i+1]]
}

func (i *KeyFunctions) FromString(s string) error {
	for j := 0; j < len(_KeyFunctions_index)-1; j++ {
		if s == _KeyFunctions_name[_KeyFunctions_index[j]:_KeyFunctions_index[j+1]] {
			*i = KeyFunctions(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type KeyFunctions", s)
}
