package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager(t *testing.T) {
	m := NewManager(&Config{
		Driver: "bcrypt",
	})

	assert.Equal(t, "bcrypt", m.getDefaultDriver())

	assert.True(t, NewBcryptHasher().Check("123456", m.Driver().MustMake("123456")))
	assert.False(t, NewMd5Hasher().Check("123456", m.Driver().MustMake("123456")))

	assert.True(t, NewBcryptHasher().Check("123456", m.Driver("bcrypt").MustMake("123456")))
	assert.False(t, NewMd5Hasher().Check("123456", m.Driver("bcrypt").MustMake("123456")))

	assert.True(t, NewMd5Hasher().Check("123456", m.Driver("md5").MustMake("123456")))
	assert.False(t, NewBcryptHasher().Check("123456", m.Driver("md5").MustMake("123456")))

	assert.Panicsf(t, func() { m.Driver("unknown") }, "unknown driver: unknown")
}
