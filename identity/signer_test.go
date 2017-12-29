package identity

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/hex"
)

func TestSigningMessageWithUnlockedAccount(t *testing.T) {
	identity := FromAddress("0x1e35193c8cadAA15b43B05ae3D882C91F49BB0Aa")
	ks := getIdentityTestKeystore()
	ks.Unlock(identityToAccount(identity), "test_passphrase")

	signer := keystoreSigner{ks, identityToAccount(identity)}
	sig, err := signer.Sign([]byte("Boop!"))
	assert.NoError(t, err)
	assert.Equal(
		t,
		"51e8c02f544c20a5b5b92894ffdd4dad90a71d994cad608cb3157b9ed7757de2758b9c0fc51bfaf079a4d60e81cc83c14cf6900c9bc031ba3f44cb119b82a5f000",
		hex.EncodeToString(sig),
	)
}
