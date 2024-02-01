package version

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	vMajor, vMinor, vPatch uint16
	result                 uint64
	str                    string
}

func TestAsBigInt(t *testing.T) {
	require := require.New(t)

	prev := testcase{0, 0, 0, 0, "0.0.0"}
	for _, next := range []testcase{
		{0, 0, 1, 1, "0.0.1"},
		{0, 0, 2, 2, "0.0.2"},
		{0, 1, 0, 1000000, "0.1.0"},
		{0, 1, math.MaxUint16, 1065535, "0.1.65535"},
		{1, 0, 0, 1000000000000, "1.0.0"},
		{1, 0, math.MaxUint16, 1000000065535, "1.0.65535"},
		{1, 1, 0, 1000001000000, "1.1.0"},
		{2, 9, 9, 2000009000009, "2.9.9"},
		{3, 1, 0, 3000001000000, "3.1.0"},
		{math.MaxUint16, math.MaxUint16, math.MaxUint16, 65535065535065535, "65535.65535.65535"},
	} {
		a := ToU64(prev.vMajor, prev.vMinor, prev.vPatch)
		b := ToU64(next.vMajor, next.vMinor, next.vPatch)
		require.Equal(a, prev.result)
		require.Equal(b, next.result)
		require.Equal(U64ToString(a), prev.str)
		require.Equal(U64ToString(b), next.str)
		require.Greater(b, a)
		prev = next
	}
}

func TestParseVersion_WithValidVersion(t *testing.T) {
	major, minor, patch := ParseVersionString("go-x1/v1.2.3-asdasadas")
	require.Equal(t, uint16(1), major)
	require.Equal(t, uint16(2), minor)
	require.Equal(t, uint16(3), patch)
}

func TestParseVersion_WithInvalidVersion(t *testing.T) {
	major, minor, patch := ParseVersionString("go-x1/v1.2.asdasadas")
	require.Equal(t, uint16(0), major)
	require.Equal(t, uint16(0), minor)
	require.Equal(t, uint16(0), patch)
}

func TestParseVersion_WithNoVersion(t *testing.T) {
	major, minor, patch := ParseVersionString("go-x1/asdasadas")
	require.Equal(t, uint16(0), major)
	require.Equal(t, uint16(0), minor)
	require.Equal(t, uint16(0), patch)
}

func TestParseVersion_WithMaxUint16Version(t *testing.T) {
	major, minor, patch := ParseVersionString("go-x1/v65535.65535.65535-asdasadas")
	require.Equal(t, uint16(65535), major)
	require.Equal(t, uint16(65535), minor)
	require.Equal(t, uint16(65535), patch)
}

func TestParseVersion_WithAnotherInvalidVersion(t *testing.T) {
	major, minor, patch := ParseVersionString("go-x1/v-1.2.0.asdasadas")
	require.Equal(t, uint16(0), major)
	require.Equal(t, uint16(0), minor)
	require.Equal(t, uint16(0), patch)
}

func TestParseVersionStringIntoU64_WithValidVersion(t *testing.T) {
	result := ParseVersionStringIntoU64("go-x1/v1.2.3-asdasadas")
	require.Equal(t, uint64(1000002000003), result)
}

func TestParseVersionStringIntoU64_WithInvalidVersion(t *testing.T) {
	result := ParseVersionStringIntoU64("go-x1/v1.2.asdasadas")
	require.Equal(t, uint64(0), result)
}

func TestParseVersionStringIntoU64_WithNoVersion(t *testing.T) {
	result := ParseVersionStringIntoU64("go-x1/asdasadas")
	require.Equal(t, uint64(0), result)
}

func TestParseVersionStringIntoU64_WithMaxUint16Version(t *testing.T) {
	result := ParseVersionStringIntoU64("go-x1/v65535.65535.65535-asdasadas")
	require.Equal(t, uint64(65535065535065535), result)
}

func TestParseVersionStringIntoU64_WithMaxUint16Version2(t *testing.T) {
	result := ParseVersionStringIntoU64("go-x1/v1.1.3-asdasadas")
	require.Equal(t, uint64(1000001000003), result)
}
