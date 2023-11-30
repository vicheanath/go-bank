package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	str := RandomString(6)
	require.Equal(t, 6, len(str))
}

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	require.NotEmpty(t, owner)
}

func TestRandomMoney(t *testing.T) {
	money := RandomMoney()
	require.NotEmpty(t, money)
	require.True(t, money > 0)
}

func TestRandomCurrency(t *testing.T) {
	currency := RandomCurrency()
	require.NotEmpty(t, currency)
	require.Equal(t, 3, len(currency))
	currencies := []string{USD, KHR ,THB}
	require.Contains(t, currencies, currency)
}

func TestRandomInt(t *testing.T) {
	n := RandomInt(0, 1000)
	require.NotEmpty(t, n)
	require.True(t, n >= 0 && n <= 1000)
}
