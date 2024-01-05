package verifier

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgon2Hash(t *testing.T) {
	assert := require.New(t)

	salt := []byte("XEN10082022XEN")
	key := []byte("0000da975bd6ec3aa878dadc395943619d23407371bc15066c1505ef23203d871633c687a9e5e89f5fc7fb61f05e1ff4ec49ecee28577c5143711185afe2d5a5")[:]

	hash, err := Argon2Hash(1, 8, 1, salt, key)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	assert.Equal("$argon2id$v=19$m=8,t=1,p=1$WEVOMTAwODIwMjJYRU4$0kyTUOCcr1NOOnc8ynGWGGldI+8ssTIW5Xy3/VWztQ6YLC8p7llYzzUiEe7pgBk/aVKhjsDH6ArPSSDJng+KzA", hash)
}
