package verifier

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFindTokensFromHashXNM(t *testing.T) {
	assert := require.New(t)

	tokens := FindTokensFromHash("$argon2id$v=19$m=99400,t=1,p=1$salt$aaaaaXEN11aaaaa", time.Unix(0, 0))

	assert.Equal(1, len(tokens))
	assert.Equal("XNM", tokens[0].name)
	assert.Equal(1, tokens[0].currencyCode)
	assert.Equal(10, tokens[0].value)
}

func TestFindTokensFromHashXUNI(t *testing.T) {
	assert := require.New(t)

	for i := 0; i < 9; i++ {
		for j := 0; j < 60; j++ {
			timestamp := time.Date(2023, 0, 0, 0, j, 0, 0, time.UTC)
			tokens := FindTokensFromHash(fmt.Sprintf("$argon2id$v=19$m=99400,t=1,p=1$salt$aaaaaXUNI%d%daaaaa", i, j), timestamp)

			if (0 <= j && j < 5) || (55 <= j && j < 60) {
				assert.Equal(1, len(tokens))
				assert.Equal("XUNI", tokens[0].name)
				assert.Equal(2, tokens[0].currencyCode)
				assert.Equal(1, tokens[0].value)
			} else {
				assert.Equal(0, len(tokens))
			}
		}
	}
}

func TestFindTokensFromHashSB(t *testing.T) {
	assert := require.New(t)

	blockTime := time.Now()

	tokens := FindTokensFromHash("$argon2id$v=19$m=99400,t=1,p=1$salt$XEN11AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", blockTime)
	assert.Equal(2, len(tokens))
	assert.Equal("XNM", tokens[0].name)
	assert.Equal(1, tokens[0].currencyCode)
	assert.Equal(10, tokens[0].value)
	assert.Equal("SB", tokens[1].name)
	assert.Equal(3, tokens[1].currencyCode)
	assert.Equal(1, tokens[1].value)
}
