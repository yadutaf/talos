// Code generated by "enumer -type=ConfigLayer -linecomment -text"; DO NOT EDIT.

package network

import (
	"fmt"
	"strings"
)

const _ConfigLayerName = "defaultcmdlineplatformoperatorconfiguration"

var _ConfigLayerIndex = [...]uint8{0, 7, 14, 22, 30, 43}

const _ConfigLayerLowerName = "defaultcmdlineplatformoperatorconfiguration"

func (i ConfigLayer) String() string {
	if i < 0 || i >= ConfigLayer(len(_ConfigLayerIndex)-1) {
		return fmt.Sprintf("ConfigLayer(%d)", i)
	}
	return _ConfigLayerName[_ConfigLayerIndex[i]:_ConfigLayerIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ConfigLayerNoOp() {
	var x [1]struct{}
	_ = x[ConfigDefault-(0)]
	_ = x[ConfigCmdline-(1)]
	_ = x[ConfigPlatform-(2)]
	_ = x[ConfigOperator-(3)]
	_ = x[ConfigMachineConfiguration-(4)]
}

var _ConfigLayerValues = []ConfigLayer{ConfigDefault, ConfigCmdline, ConfigPlatform, ConfigOperator, ConfigMachineConfiguration}

var _ConfigLayerNameToValueMap = map[string]ConfigLayer{
	_ConfigLayerName[0:7]:        ConfigDefault,
	_ConfigLayerLowerName[0:7]:   ConfigDefault,
	_ConfigLayerName[7:14]:       ConfigCmdline,
	_ConfigLayerLowerName[7:14]:  ConfigCmdline,
	_ConfigLayerName[14:22]:      ConfigPlatform,
	_ConfigLayerLowerName[14:22]: ConfigPlatform,
	_ConfigLayerName[22:30]:      ConfigOperator,
	_ConfigLayerLowerName[22:30]: ConfigOperator,
	_ConfigLayerName[30:43]:      ConfigMachineConfiguration,
	_ConfigLayerLowerName[30:43]: ConfigMachineConfiguration,
}

var _ConfigLayerNames = []string{
	_ConfigLayerName[0:7],
	_ConfigLayerName[7:14],
	_ConfigLayerName[14:22],
	_ConfigLayerName[22:30],
	_ConfigLayerName[30:43],
}

// ConfigLayerString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ConfigLayerString(s string) (ConfigLayer, error) {
	if val, ok := _ConfigLayerNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ConfigLayerNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ConfigLayer values", s)
}

// ConfigLayerValues returns all values of the enum
func ConfigLayerValues() []ConfigLayer {
	return _ConfigLayerValues
}

// ConfigLayerStrings returns a slice of all String values of the enum
func ConfigLayerStrings() []string {
	strs := make([]string, len(_ConfigLayerNames))
	copy(strs, _ConfigLayerNames)
	return strs
}

// IsAConfigLayer returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ConfigLayer) IsAConfigLayer() bool {
	for _, v := range _ConfigLayerValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for ConfigLayer
func (i ConfigLayer) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ConfigLayer
func (i *ConfigLayer) UnmarshalText(text []byte) error {
	var err error
	*i, err = ConfigLayerString(string(text))
	return err
}
