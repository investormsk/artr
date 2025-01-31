package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
)

// Parameter store keys
var ()

// ParamKeyTable for storage module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for storage at genesis
type Params struct {
	// KeyParamName string `json:"key_param_name"`
}

// NewParams creates a new Params object
func NewParams() Params {
	return Params{}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf(`
	// TODO: Return all the params as a string
	`)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		// params.NewParamSetPair(KeyParamName, &p.ParamName),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams()
}
