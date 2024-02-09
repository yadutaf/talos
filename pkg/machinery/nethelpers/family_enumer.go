// Code generated by "enumer -type=Family -linecomment -text"; DO NOT EDIT.

package nethelpers

import (
	"fmt"
	"strings"
)

const (
	_FamilyName_0      = "inet4"
	_FamilyLowerName_0 = "inet4"
	_FamilyName_1      = "inet6"
	_FamilyLowerName_1 = "inet6"
)

var (
	_FamilyIndex_0 = [...]uint8{0, 5}
	_FamilyIndex_1 = [...]uint8{0, 5}
)

func (i Family) String() string {
	switch {
	case i == 2:
		return _FamilyName_0
	case i == 10:
		return _FamilyName_1
	default:
		return fmt.Sprintf("Family(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FamilyNoOp() {
	var x [1]struct{}
	_ = x[FamilyInet4-(2)]
	_ = x[FamilyInet6-(10)]
}

var _FamilyValues = []Family{FamilyInet4, FamilyInet6}

var _FamilyNameToValueMap = map[string]Family{
	_FamilyName_0[0:5]:      FamilyInet4,
	_FamilyLowerName_0[0:5]: FamilyInet4,
	_FamilyName_1[0:5]:      FamilyInet6,
	_FamilyLowerName_1[0:5]: FamilyInet6,
}

var _FamilyNames = []string{
	_FamilyName_0[0:5],
	_FamilyName_1[0:5],
}

// FamilyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FamilyString(s string) (Family, error) {
	if val, ok := _FamilyNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FamilyNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Family values", s)
}

// FamilyValues returns all values of the enum
func FamilyValues() []Family {
	return _FamilyValues
}

// FamilyStrings returns a slice of all String values of the enum
func FamilyStrings() []string {
	strs := make([]string, len(_FamilyNames))
	copy(strs, _FamilyNames)
	return strs
}

// IsAFamily returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Family) IsAFamily() bool {
	for _, v := range _FamilyValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Family
func (i Family) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Family
func (i *Family) UnmarshalText(text []byte) error {
	var err error
	*i, err = FamilyString(string(text))
	return err
}
