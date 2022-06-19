package srcparser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewVarSpec(t *testing.T) {
	actual := NewVarSpec("mypkg", "varname", false)
	require.NotNil(t, actual)
}

func TestVarSpec_FullName(t *testing.T) {
	tests := []struct {
		name     string
		varspec  *VarSpec
		expected string
	}{
		{
			name:     "all filled",
			varspec:  NewVarSpec("mypkg", "myvar", false),
			expected: "mypkg.myvar",
		},
		{
			name:     "empty pkg",
			varspec:  NewVarSpec("", "myvar", false),
			expected: "myvar",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, test.varspec.FullName())
	}

}

func TestVarSpec_Header(t *testing.T) {
	tests := []struct {
		name     string
		varspec  *VarSpec
		expected string
	}{
		{
			name:     "all filled",
			varspec:  NewVarSpec("mypkg", "myvar", false),
			expected: "mypkg.myvar",
		},
		{
			name:     "all filled with pointer",
			varspec:  NewVarSpec("mypkg", "myvar", true),
			expected: "*mypkg.myvar",
		},
		{
			name:     "empty pkg",
			varspec:  NewVarSpec("", "myvar", false),
			expected: "myvar",
		},
		{
			name:     "empty pkg pointer",
			varspec:  NewVarSpec("", "myvar", true),
			expected: "*myvar",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, test.varspec.Header())
	}
}
