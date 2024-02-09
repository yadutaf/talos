// Code generated by "enumer -type=Port -text"; DO NOT EDIT.

package nethelpers

import (
	"fmt"
	"strings"
)

const (
	_PortName_0      = "TwistedPairAUIMIIFibreBNCDirectAttach"
	_PortLowerName_0 = "twistedpairauimiifibrebncdirectattach"
	_PortName_1      = "None"
	_PortLowerName_1 = "none"
	_PortName_2      = "Other"
	_PortLowerName_2 = "other"
)

var (
	_PortIndex_0 = [...]uint8{0, 11, 14, 17, 22, 25, 37}
	_PortIndex_1 = [...]uint8{0, 4}
	_PortIndex_2 = [...]uint8{0, 5}
)

func (i Port) String() string {
	switch {
	case 0 <= i && i <= 5:
		return _PortName_0[_PortIndex_0[i]:_PortIndex_0[i+1]]
	case i == 239:
		return _PortName_1
	case i == 255:
		return _PortName_2
	default:
		return fmt.Sprintf("Port(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PortNoOp() {
	var x [1]struct{}
	_ = x[TwistedPair-(0)]
	_ = x[AUI-(1)]
	_ = x[MII-(2)]
	_ = x[Fibre-(3)]
	_ = x[BNC-(4)]
	_ = x[DirectAttach-(5)]
	_ = x[None-(239)]
	_ = x[Other-(255)]
}

var _PortValues = []Port{TwistedPair, AUI, MII, Fibre, BNC, DirectAttach, None, Other}

var _PortNameToValueMap = map[string]Port{
	_PortName_0[0:11]:       TwistedPair,
	_PortLowerName_0[0:11]:  TwistedPair,
	_PortName_0[11:14]:      AUI,
	_PortLowerName_0[11:14]: AUI,
	_PortName_0[14:17]:      MII,
	_PortLowerName_0[14:17]: MII,
	_PortName_0[17:22]:      Fibre,
	_PortLowerName_0[17:22]: Fibre,
	_PortName_0[22:25]:      BNC,
	_PortLowerName_0[22:25]: BNC,
	_PortName_0[25:37]:      DirectAttach,
	_PortLowerName_0[25:37]: DirectAttach,
	_PortName_1[0:4]:        None,
	_PortLowerName_1[0:4]:   None,
	_PortName_2[0:5]:        Other,
	_PortLowerName_2[0:5]:   Other,
}

var _PortNames = []string{
	_PortName_0[0:11],
	_PortName_0[11:14],
	_PortName_0[14:17],
	_PortName_0[17:22],
	_PortName_0[22:25],
	_PortName_0[25:37],
	_PortName_1[0:4],
	_PortName_2[0:5],
}

// PortString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PortString(s string) (Port, error) {
	if val, ok := _PortNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PortNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Port values", s)
}

// PortValues returns all values of the enum
func PortValues() []Port {
	return _PortValues
}

// PortStrings returns a slice of all String values of the enum
func PortStrings() []string {
	strs := make([]string, len(_PortNames))
	copy(strs, _PortNames)
	return strs
}

// IsAPort returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Port) IsAPort() bool {
	for _, v := range _PortValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Port
func (i Port) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Port
func (i *Port) UnmarshalText(text []byte) error {
	var err error
	*i, err = PortString(string(text))
	return err
}
