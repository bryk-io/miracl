/*
   Copyright (C) 2019 MIRACL UK Ltd.

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.


    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

     https://www.gnu.org/licenses/agpl-3.0.en.html

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   You can be released from the requirements of the license by purchasing
   a commercial license. Buying such a license is mandatory as soon as you
   develop commercial activities involving the MIRACL Core Crypto SDK
   without disclosing the source code of your own applications, or shipping
   the MIRACL Core Crypto SDK with a closed source product.
*/

/* test driver and function exerciser for ECDH/ECIES/ECDSA API Functions */

package main

import (
	"fmt"

	"go.bryk.io/miracl/core"
	"go.bryk.io/miracl/core/ED25519"
	"go.bryk.io/miracl/core/GOLDILOCKS"
	"go.bryk.io/miracl/core/NIST256"
	"go.bryk.io/miracl/core/RSA2048"
)

func ecdh_ED25519(rng *core.RAND) {
	//	j:=0
	pp := "M0ng00se"
	res := 0

	var sha = ED25519.HASH_TYPE

	var S1 [ED25519.EGS]byte
	var W0 [2*ED25519.EFS + 1]byte
	var W1 [2*ED25519.EFS + 1]byte
	var Z0 [ED25519.EFS]byte
	var Z1 [ED25519.EFS]byte
	var SALT [8]byte
	var P1 [3]byte
	var P2 [4]byte
	var V [2*ED25519.EFS + 1]byte
	var M [17]byte
	var T [12]byte
	var CS [ED25519.EGS]byte
	var DS [ED25519.EGS]byte

	for i := 0; i < 8; i++ {
		SALT[i] = byte(i + 1)
	} // set Salt

	fmt.Printf("\nTesting ECDH/ECDSA/ECIES for curve ED25519\n")
	fmt.Printf("Alice's Passphrase= " + pp)
	fmt.Printf("\n")
	PW := []byte(pp)

	/* private key S0 of size MGS bytes derived from Password and Salt */

	S0 := core.PBKDF2(core.MC_SHA2, sha, PW, SALT[:], 1000, ED25519.EGS)

	fmt.Printf("Alice's private key= 0x")
	printBinary(S0)

	/* Generate Key pair S/W */
	ED25519.ECDH_KEY_PAIR_GENERATE(nil, S0, W0[:])

	fmt.Printf("Alice's public key= 0x")
	printBinary(W0[:])

	res = ED25519.ECDH_PUBLIC_KEY_VALIDATE(W0[:])
	if res != 0 {
		fmt.Printf("ECP Public Key is invalid!\n")
		return
	}

	/* Random private key for other party */
	ED25519.ECDH_KEY_PAIR_GENERATE(rng, S1[:], W1[:])

	fmt.Printf("Servers private key= 0x")
	printBinary(S1[:])

	fmt.Printf("Servers public key= 0x")
	printBinary(W1[:])

	res = ED25519.ECDH_PUBLIC_KEY_VALIDATE(W1[:])
	if res != 0 {
		fmt.Printf("ECP Public Key is invalid!\n")
		return
	}
	/* Calculate common key using DH - IEEE 1363 method */

	ED25519.ECDH_ECPSVDP_DH(S0, W1[:], Z0[:])
	ED25519.ECDH_ECPSVDP_DH(S1[:], W0[:], Z1[:])

	same := true
	for i := 0; i < ED25519.EFS; i++ {
		if Z0[i] != Z1[i] {
			same = false
		}
	}

	if !same {
		fmt.Printf("*** ECPSVDP-DH Failed\n")
		return
	}

	KEY := core.KDF2(core.MC_SHA2, sha, Z0[:], nil, ED25519.AESKEY)

	fmt.Printf("Alice's DH Key=  0x")
	printBinary(KEY)
	fmt.Printf("Servers DH Key=  0x")
	printBinary(KEY)

	if ED25519.CURVETYPE != ED25519.MONTGOMERY {
		fmt.Printf("Testing ECIES\n")

		P1[0] = 0x0
		P1[1] = 0x1
		P1[2] = 0x2
		P2[0] = 0x0
		P2[1] = 0x1
		P2[2] = 0x2
		P2[3] = 0x3

		for i := 0; i <= 16; i++ {
			M[i] = byte(i)
		}

		C := ED25519.ECDH_ECIES_ENCRYPT(sha, P1[:], P2[:], rng, W1[:], M[:], V[:], T[:])

		fmt.Printf("Ciphertext= \n")
		fmt.Printf("V= 0x")
		printBinary(V[:])
		fmt.Printf("C= 0x")
		printBinary(C)
		fmt.Printf("T= 0x")
		printBinary(T[:])

		RM := ED25519.ECDH_ECIES_DECRYPT(sha, P1[:], P2[:], V[:], C, T[:], S1[:])
		if RM == nil {
			fmt.Printf("*** ECIES Decryption Failed\n")
			return
		} else {
			fmt.Printf("Decryption succeeded\n")
		}

		fmt.Printf("Message is 0x")
		printBinary(RM)

		fmt.Printf("Testing ECDSA\n")

		if ED25519.ECDH_ECPSP_DSA(sha, rng, S0, M[:], CS[:], DS[:]) != 0 {
			fmt.Printf("***ECDSA Signature Failed\n")
			return
		}
		fmt.Printf("Signature= \n")
		fmt.Printf("C= 0x")
		printBinary(CS[:])
		fmt.Printf("D= 0x")
		printBinary(DS[:])

		if ED25519.ECDH_ECPVP_DSA(sha, W0[:], M[:], CS[:], DS[:]) != 0 {
			fmt.Printf("***ECDSA Verification Failed\n")
			return
		} else {
			fmt.Printf("ECDSA Signature/Verification succeeded \n")
		}
	}
}

func ecdh_NIST256(rng *core.RAND) {
	//	j:=0
	pp := "M0ng00se"
	res := 0

	var sha = NIST256.HASH_TYPE

	var S1 [NIST256.EGS]byte
	var W0 [2*NIST256.EFS + 1]byte
	var W1 [2*NIST256.EFS + 1]byte
	var Z0 [NIST256.EFS]byte
	var Z1 [NIST256.EFS]byte
	var SALT [8]byte
	var P1 [3]byte
	var P2 [4]byte
	var V [2*NIST256.EFS + 1]byte
	var M [17]byte
	var T [12]byte
	var CS [NIST256.EGS]byte
	var DS [NIST256.EGS]byte

	for i := 0; i < 8; i++ {
		SALT[i] = byte(i + 1)
	} // set Salt

	fmt.Printf("\nTesting ECDH/ECDSA/ECIES for curve NIST256\n")
	fmt.Printf("Alice's Passphrase= " + pp)
	fmt.Printf("\n")
	PW := []byte(pp)

	/* private key S0 of size MGS bytes derived from Password and Salt */

	S0 := core.PBKDF2(core.MC_SHA2, sha, PW, SALT[:], 1000, NIST256.EGS)

	fmt.Printf("Alice's private key= 0x")
	printBinary(S0)

	/* Generate Key pair S/W */
	NIST256.ECDH_KEY_PAIR_GENERATE(nil, S0, W0[:])

	fmt.Printf("Alice's public key= 0x")
	printBinary(W0[:])

	res = NIST256.ECDH_PUBLIC_KEY_VALIDATE(W0[:])
	if res != 0 {
		fmt.Printf("ECP Public Key is invalid!\n")
		return
	}

	/* Random private key for other party */
	NIST256.ECDH_KEY_PAIR_GENERATE(rng, S1[:], W1[:])

	fmt.Printf("Servers private key= 0x")
	printBinary(S1[:])

	fmt.Printf("Servers public key= 0x")
	printBinary(W1[:])

	res = NIST256.ECDH_PUBLIC_KEY_VALIDATE(W1[:])
	if res != 0 {
		fmt.Printf("ECP Public Key is invalid!\n")
		return
	}
	/* Calculate common key using DH - IEEE 1363 method */

	NIST256.ECDH_ECPSVDP_DH(S0, W1[:], Z0[:])
	NIST256.ECDH_ECPSVDP_DH(S1[:], W0[:], Z1[:])

	same := true
	for i := 0; i < NIST256.EFS; i++ {
		if Z0[i] != Z1[i] {
			same = false
		}
	}

	if !same {
		fmt.Printf("*** ECPSVDP-DH Failed\n")
		return
	}

	KEY := core.KDF2(core.MC_SHA2, sha, Z0[:], nil, NIST256.AESKEY)

	fmt.Printf("Alice's DH Key=  0x")
	printBinary(KEY)
	fmt.Printf("Servers DH Key=  0x")
	printBinary(KEY)

	if NIST256.CURVETYPE != NIST256.MONTGOMERY {
		fmt.Printf("Testing ECIES\n")

		P1[0] = 0x0
		P1[1] = 0x1
		P1[2] = 0x2
		P2[0] = 0x0
		P2[1] = 0x1
		P2[2] = 0x2
		P2[3] = 0x3

		for i := 0; i <= 16; i++ {
			M[i] = byte(i)
		}

		C := NIST256.ECDH_ECIES_ENCRYPT(sha, P1[:], P2[:], rng, W1[:], M[:], V[:], T[:])

		fmt.Printf("Ciphertext= \n")
		fmt.Printf("V= 0x")
		printBinary(V[:])
		fmt.Printf("C= 0x")
		printBinary(C)
		fmt.Printf("T= 0x")
		printBinary(T[:])

		RM := NIST256.ECDH_ECIES_DECRYPT(sha, P1[:], P2[:], V[:], C, T[:], S1[:])
		if RM == nil {
			fmt.Printf("*** ECIES Decryption Failed\n")
			return
		} else {
			fmt.Printf("Decryption succeeded\n")
		}

		fmt.Printf("Message is 0x")
		printBinary(RM)

		fmt.Printf("Testing ECDSA\n")

		if NIST256.ECDH_ECPSP_DSA(sha, rng, S0, M[:], CS[:], DS[:]) != 0 {
			fmt.Printf("***ECDSA Signature Failed\n")
			return
		}
		fmt.Printf("Signature= \n")
		fmt.Printf("C= 0x")
		printBinary(CS[:])
		fmt.Printf("D= 0x")
		printBinary(DS[:])

		if NIST256.ECDH_ECPVP_DSA(sha, W0[:], M[:], CS[:], DS[:]) != 0 {
			fmt.Printf("***ECDSA Verification Failed\n")
			return
		} else {
			fmt.Printf("ECDSA Signature/Verification succeeded \n")
		}
	}
}

func ecdh_GOLDILOCKS(rng *core.RAND) {
	//	j:=0
	pp := "M0ng00se"
	res := 0

	var sha = GOLDILOCKS.HASH_TYPE

	var S1 [GOLDILOCKS.EGS]byte
	var W0 [2*GOLDILOCKS.EFS + 1]byte
	var W1 [2*GOLDILOCKS.EFS + 1]byte
	var Z0 [GOLDILOCKS.EFS]byte
	var Z1 [GOLDILOCKS.EFS]byte
	var SALT [8]byte
	var P1 [3]byte
	var P2 [4]byte
	var V [2*GOLDILOCKS.EFS + 1]byte
	var M [17]byte
	var T [12]byte
	var CS [GOLDILOCKS.EGS]byte
	var DS [GOLDILOCKS.EGS]byte

	for i := 0; i < 8; i++ {
		SALT[i] = byte(i + 1)
	} // set Salt

	fmt.Printf("\nTesting ECDH/ECDSA/ECIES for curve GOLDILOCKS\n")
	fmt.Printf("Alice's Passphrase= " + pp)
	fmt.Printf("\n")
	PW := []byte(pp)

	/* private key S0 of size MGS bytes derived from Password and Salt */

	S0 := core.PBKDF2(core.MC_SHA2, sha, PW, SALT[:], 1000, GOLDILOCKS.EGS)

	fmt.Printf("Alice's private key= 0x")
	printBinary(S0)

	/* Generate Key pair S/W */
	GOLDILOCKS.ECDH_KEY_PAIR_GENERATE(nil, S0, W0[:])

	fmt.Printf("Alice's public key= 0x")
	printBinary(W0[:])

	res = GOLDILOCKS.ECDH_PUBLIC_KEY_VALIDATE(W0[:])
	if res != 0 {
		fmt.Printf("ECP Public Key is invalid!\n")
		return
	}

	/* Random private key for other party */
	GOLDILOCKS.ECDH_KEY_PAIR_GENERATE(rng, S1[:], W1[:])

	fmt.Printf("Servers private key= 0x")
	printBinary(S1[:])

	fmt.Printf("Servers public key= 0x")
	printBinary(W1[:])

	res = GOLDILOCKS.ECDH_PUBLIC_KEY_VALIDATE(W1[:])
	if res != 0 {
		fmt.Printf("ECP Public Key is invalid!\n")
		return
	}
	/* Calculate common key using DH - IEEE 1363 method */

	GOLDILOCKS.ECDH_ECPSVDP_DH(S0, W1[:], Z0[:])
	GOLDILOCKS.ECDH_ECPSVDP_DH(S1[:], W0[:], Z1[:])

	same := true
	for i := 0; i < GOLDILOCKS.EFS; i++ {
		if Z0[i] != Z1[i] {
			same = false
		}
	}

	if !same {
		fmt.Printf("*** ECPSVDP-DH Failed\n")
		return
	}

	KEY := core.KDF2(core.MC_SHA2, sha, Z0[:], nil, GOLDILOCKS.AESKEY)

	fmt.Printf("Alice's DH Key=  0x")
	printBinary(KEY)
	fmt.Printf("Servers DH Key=  0x")
	printBinary(KEY)

	if GOLDILOCKS.CURVETYPE != GOLDILOCKS.MONTGOMERY {
		fmt.Printf("Testing ECIES\n")

		P1[0] = 0x0
		P1[1] = 0x1
		P1[2] = 0x2
		P2[0] = 0x0
		P2[1] = 0x1
		P2[2] = 0x2
		P2[3] = 0x3

		for i := 0; i <= 16; i++ {
			M[i] = byte(i)
		}

		C := GOLDILOCKS.ECDH_ECIES_ENCRYPT(sha, P1[:], P2[:], rng, W1[:], M[:], V[:], T[:])

		fmt.Printf("Ciphertext= \n")
		fmt.Printf("V= 0x")
		printBinary(V[:])
		fmt.Printf("C= 0x")
		printBinary(C)
		fmt.Printf("T= 0x")
		printBinary(T[:])

		RM := GOLDILOCKS.ECDH_ECIES_DECRYPT(sha, P1[:], P2[:], V[:], C, T[:], S1[:])
		if RM == nil {
			fmt.Printf("*** ECIES Decryption Failed\n")
			return
		} else {
			fmt.Printf("Decryption succeeded\n")
		}

		fmt.Printf("Message is 0x")
		printBinary(RM)

		fmt.Printf("Testing ECDSA\n")

		if GOLDILOCKS.ECDH_ECPSP_DSA(sha, rng, S0, M[:], CS[:], DS[:]) != 0 {
			fmt.Printf("***ECDSA Signature Failed\n")
			return
		}
		fmt.Printf("Signature= \n")
		fmt.Printf("C= 0x")
		printBinary(CS[:])
		fmt.Printf("D= 0x")
		printBinary(DS[:])

		if GOLDILOCKS.ECDH_ECPVP_DSA(sha, W0[:], M[:], CS[:], DS[:]) != 0 {
			fmt.Printf("***ECDSA Verification Failed\n")
			return
		} else {
			fmt.Printf("ECDSA Signature/Verification succeeded \n")
		}
	}
}

func rsa_2048(rng *core.RAND) {
	var sha = RSA2048.RSA_HASH_TYPE
	message := "Hello World\n"

	pub := RSA2048.New_public_key(RSA2048.FFLEN)
	priv := RSA2048.New_private_key(RSA2048.HFLEN)

	var ML [RSA2048.RFS]byte
	var C [RSA2048.RFS]byte
	var S [RSA2048.RFS]byte

	fmt.Printf("\nTesting RSA 2048-bit\n")
	fmt.Printf("Generating public/private key pair\n")
	RSA2048.RSA_KEY_PAIR(rng, 65537, priv, pub)

	M := []byte(message)

	fmt.Printf("Encrypting test string\n")
	E := RSA2048.RSA_OAEP_ENCODE(sha, M, rng, nil) /* OAEP encode message M to E  */

	RSA2048.RSA_ENCRYPT(pub, E, C[:]) /* encrypt encoded message */
	fmt.Printf("Ciphertext= 0x")
	printBinary(C[:])

	fmt.Printf("Decrypting test string\n")
	RSA2048.RSA_DECRYPT(priv, C[:], ML[:])
	MS := RSA2048.RSA_OAEP_DECODE(sha, nil, ML[:]) /* OAEP decode message  */

	message = string(MS)
	fmt.Printf(message)

	fmt.Printf("Signing message\n")
	RSA2048.RSA_PKCS15(sha, M, C[:])

	RSA2048.RSA_DECRYPT(priv, C[:], S[:]) /* create signature in S */

	fmt.Printf("Signature= 0x")
	printBinary(S[:])

	RSA2048.RSA_ENCRYPT(pub, S[:], ML[:])

	cmp := true
	if len(C) != len(ML) {
		cmp = false
	} else {
		for j := 0; j < len(C); j++ {
			if C[j] != ML[j] {
				cmp = false
			}
		}
	}
	if cmp {
		fmt.Printf("Signature is valid\n")
	} else {
		fmt.Printf("Signature is INVALID\n")
	}

}
