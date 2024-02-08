package version

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	params.VersionMajor = 1     // Major version component of the current release
	params.VersionMinor = 1     // Minor version component of the current release
	params.VersionPatch = 5     // Patch version component of the current release
	params.VersionMeta = "rc.2" // Version metadata to append to the version string
}

func BigToString(b *big.Int) string {
	if len(b.Bytes()) > 8 {
		return "_malformed_version_"
	}
	return U64ToString(b.Uint64())
}

func AsString() string {
	return ToString(uint16(params.VersionMajor), uint16(params.VersionMinor), uint16(params.VersionPatch))
}

func AsU64() uint64 {
	return ToU64(uint16(params.VersionMajor), uint16(params.VersionMinor), uint16(params.VersionPatch))
}

func AsBigInt() *big.Int {
	return new(big.Int).SetUint64(AsU64())
}

func ToU64(vMajor, vMinor, vPatch uint16) uint64 {
	return uint64(vMajor)*1e12 + uint64(vMinor)*1e6 + uint64(vPatch)
}

func ToString(major, minor, patch uint16) string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func U64ToString(v uint64) string {
	return ToString(uint16((v/1e12)%1e6), uint16((v/1e6)%1e6), uint16(v%1e6))
}

func ParseVersionString(clientName string) (major, minor, patch uint16) {
	re := regexp.MustCompile(`v([0-9]+)\.([0-9]+)\.([0-9]+)`)
	groups := re.FindStringSubmatch(strings.ToLower(clientName))

	if len(groups) < 4 {
		return 0, 0, 0
	}

	value, _ := strconv.ParseUint(groups[1], 10, 16)
	major = uint16(value)
	value, _ = strconv.ParseUint(groups[2], 10, 16)
	minor = uint16(value)
	value, _ = strconv.ParseUint(groups[3], 10, 16)
	patch = uint16(value)

	return major, minor, patch
}

func ParseVersionStringIntoU64(version string) uint64 {
	major, miner, patch := ParseVersionString(version)
	return ToU64(major, miner, patch)
}
