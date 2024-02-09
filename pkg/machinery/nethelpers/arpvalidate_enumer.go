// Code generated by "enumer -type=ARPValidate -linecomment -text"; DO NOT EDIT.

package nethelpers

import (
	"fmt"
	"strings"
)

const _ARPValidateName = "noneactivebackupall"

var _ARPValidateIndex = [...]uint8{0, 4, 10, 16, 19}

const _ARPValidateLowerName = "noneactivebackupall"

func (i ARPValidate) String() string {
	if i >= ARPValidate(len(_ARPValidateIndex)-1) {
		return fmt.Sprintf("ARPValidate(%d)", i)
	}
	return _ARPValidateName[_ARPValidateIndex[i]:_ARPValidateIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ARPValidateNoOp() {
	var x [1]struct{}
	_ = x[ARPValidateNone-(0)]
	_ = x[ARPValidateActive-(1)]
	_ = x[ARPValidateBackup-(2)]
	_ = x[ARPValidateAll-(3)]
}

var _ARPValidateValues = []ARPValidate{ARPValidateNone, ARPValidateActive, ARPValidateBackup, ARPValidateAll}

var _ARPValidateNameToValueMap = map[string]ARPValidate{
	_ARPValidateName[0:4]:        ARPValidateNone,
	_ARPValidateLowerName[0:4]:   ARPValidateNone,
	_ARPValidateName[4:10]:       ARPValidateActive,
	_ARPValidateLowerName[4:10]:  ARPValidateActive,
	_ARPValidateName[10:16]:      ARPValidateBackup,
	_ARPValidateLowerName[10:16]: ARPValidateBackup,
	_ARPValidateName[16:19]:      ARPValidateAll,
	_ARPValidateLowerName[16:19]: ARPValidateAll,
}

var _ARPValidateNames = []string{
	_ARPValidateName[0:4],
	_ARPValidateName[4:10],
	_ARPValidateName[10:16],
	_ARPValidateName[16:19],
}

// ARPValidateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ARPValidateString(s string) (ARPValidate, error) {
	if val, ok := _ARPValidateNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ARPValidateNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ARPValidate values", s)
}

// ARPValidateValues returns all values of the enum
func ARPValidateValues() []ARPValidate {
	return _ARPValidateValues
}

// ARPValidateStrings returns a slice of all String values of the enum
func ARPValidateStrings() []string {
	strs := make([]string, len(_ARPValidateNames))
	copy(strs, _ARPValidateNames)
	return strs
}

// IsAARPValidate returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ARPValidate) IsAARPValidate() bool {
	for _, v := range _ARPValidateValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for ARPValidate
func (i ARPValidate) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ARPValidate
func (i *ARPValidate) UnmarshalText(text []byte) error {
	var err error
	*i, err = ARPValidateString(string(text))
	return err
}
