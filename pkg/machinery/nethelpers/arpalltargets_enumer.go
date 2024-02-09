// Code generated by "enumer -type=ARPAllTargets -linecomment -text"; DO NOT EDIT.

package nethelpers

import (
	"fmt"
	"strings"
)

const _ARPAllTargetsName = "anyall"

var _ARPAllTargetsIndex = [...]uint8{0, 3, 6}

const _ARPAllTargetsLowerName = "anyall"

func (i ARPAllTargets) String() string {
	if i >= ARPAllTargets(len(_ARPAllTargetsIndex)-1) {
		return fmt.Sprintf("ARPAllTargets(%d)", i)
	}
	return _ARPAllTargetsName[_ARPAllTargetsIndex[i]:_ARPAllTargetsIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ARPAllTargetsNoOp() {
	var x [1]struct{}
	_ = x[ARPAllTargetsAny-(0)]
	_ = x[ARPAllTargetsAll-(1)]
}

var _ARPAllTargetsValues = []ARPAllTargets{ARPAllTargetsAny, ARPAllTargetsAll}

var _ARPAllTargetsNameToValueMap = map[string]ARPAllTargets{
	_ARPAllTargetsName[0:3]:      ARPAllTargetsAny,
	_ARPAllTargetsLowerName[0:3]: ARPAllTargetsAny,
	_ARPAllTargetsName[3:6]:      ARPAllTargetsAll,
	_ARPAllTargetsLowerName[3:6]: ARPAllTargetsAll,
}

var _ARPAllTargetsNames = []string{
	_ARPAllTargetsName[0:3],
	_ARPAllTargetsName[3:6],
}

// ARPAllTargetsString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ARPAllTargetsString(s string) (ARPAllTargets, error) {
	if val, ok := _ARPAllTargetsNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ARPAllTargetsNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ARPAllTargets values", s)
}

// ARPAllTargetsValues returns all values of the enum
func ARPAllTargetsValues() []ARPAllTargets {
	return _ARPAllTargetsValues
}

// ARPAllTargetsStrings returns a slice of all String values of the enum
func ARPAllTargetsStrings() []string {
	strs := make([]string, len(_ARPAllTargetsNames))
	copy(strs, _ARPAllTargetsNames)
	return strs
}

// IsAARPAllTargets returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ARPAllTargets) IsAARPAllTargets() bool {
	for _, v := range _ARPAllTargetsValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for ARPAllTargets
func (i ARPAllTargets) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ARPAllTargets
func (i *ARPAllTargets) UnmarshalText(text []byte) error {
	var err error
	*i, err = ARPAllTargetsString(string(text))
	return err
}
