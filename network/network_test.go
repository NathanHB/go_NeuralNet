package network

import (
	"testing"
)

func TestInitNewNetwork(t *testing.T) {
	net := InitNewNetwork([]uint{2, 3, 1})

	PrintNetwork(net)
}

