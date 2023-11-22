package blockchain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAccount(t *testing.T) {
	mnemoinic, err := GenerateMnemonic12()
	fmt.Println(mnemoinic)
	if err != nil {
		t.Error(err)
	}
	mnemoinicWords := strings.Split(mnemoinic, " ")
	assert.Equal(t, 12, len(mnemoinicWords))

	mnemoinic, err = GenerateMnemonic24()
	fmt.Println(mnemoinic)
	if err != nil {
		t.Error(err)
	}
	mnemoinicWords = strings.Split(mnemoinic, " ")
	assert.Equal(t, 24, len(mnemoinicWords))
}
