package main

import (
	"fmt"
	"os"

	"go.bryk.io/miracl/core"
)

func printBinary(array []byte) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("%02x", array[i])
	}
	fmt.Printf("\n")
}

func main() {
	action := "testing"
	if len(os.Args) == 2 {
		action = os.Args[1]
	}
	switch action {
	case "testing":
		testBLS()
		testECC()
		testHPKE()
		testNHS()
		testMPIN()
	case "benchmark":
		runBenchmarks()
	case "info":
		fmt.Printf("upstream commit: %s\n", Version)
	}
}

func testBLS() {
	rng := core.NewRAND()
	var raw [100]byte
	for i := 0; i < 100; i++ {
		raw[i] = byte(i + 1)
	}
	rng.Seed(100, raw[:])

	bls_BN254(rng)
	bls_BLS12383(rng)
	bls_BLS24479(rng)
	bls_BLS48556(rng)
}

func testECC() {
	rng := core.NewRAND()
	var raw [100]byte
	for i := 0; i < 100; i++ {
		raw[i] = byte(i + 1)
	}
	rng.Seed(100, raw[:])

	ecdh_ED25519(rng)
	ecdh_NIST256(rng)
	ecdh_GOLDILOCKS(rng)
	rsa_2048(rng)
}

func testHPKE() {
	rng := core.NewRAND()
	var raw [100]byte
	for i := 0; i < 100; i++ {
		raw[i] = byte(i + 1)
	}
	rng.Seed(100, raw[:])

	fmt.Printf("\nTesting HPKE for curve C25519\n")
	hpke_C25519(rng)
	fmt.Printf("\nTesting HPKE for curve NIST521\n")
	hpke_NIST521(rng)
}

func testNHS() {
	fmt.Printf("\nTesting New Hope Key Exchange\n")
	srng := core.NewRAND()
	var sraw [100]byte
	for i := 0; i < 100; i++ {
		sraw[i] = byte(i + 1)
	}
	srng.Seed(100, sraw[:])

	crng := core.NewRAND()
	var craw [100]byte
	for i := 0; i < 100; i++ {
		craw[i] = byte(i + 2)
	}
	crng.Seed(100, craw[:])

	var S [1792]byte

	var SB [1824]byte
	core.NHS_SERVER_1(srng, SB[:], S[:])
	var UC [2176]byte
	var KEYB [32]byte
	core.NHS_CLIENT(crng, SB[:], UC[:], KEYB[:])

	fmt.Printf("Bob's Key= ")
	for i := 0; i < 32; i++ {
		fmt.Printf("%02x", KEYB[i])
	}
	fmt.Printf("\n")
	var KEYA [32]byte
	core.NHS_SERVER_2(S[:], UC[:], KEYA[:])

	fmt.Printf("Alice Key= ")
	for i := 0; i < 32; i++ {
		fmt.Printf("%02x", KEYA[i])
	}
	fmt.Printf("\n")
}

func testMPIN() {
	rng := core.NewRAND()
	var raw [100]byte
	for i := 0; i < 100; i++ {
		raw[i] = byte(i)
	}
	rng.Seed(100, raw[:])

	mpin_BN254(rng)
	mpin_BLS12383(rng)
	mpin_BLS24479(rng)
	mpin_BLS48556(rng)
}

func runBenchmarks() {
	var RAW [100]byte
	rng := core.NewRAND()
	rng.Clean()
	for i := 0; i < 100; i++ {
		RAW[i] = byte(i)
	}
	rng.Seed(100, RAW[:])

	ED_25519(rng)
	NIST_256(rng)
	GOLDI_LOCKS(rng)
	BN_254(rng)
	BLS_383(rng)
	BLS_24(rng)
	BLS_48(rng)
	RSA_2048(rng)
}
const Version="723c78cbbde58c9b0239f1f09c35524ada8128a0"
