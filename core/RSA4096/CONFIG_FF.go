/*
 * Copyright (c) 2012-2020 MIRACL UK Ltd.
 *
 * This file is part of MIRACL Core
 * (see https://github.com/miracl/core).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package RSA4096

// Modulus as number of BIGs
const FFLEN int = 8

// Modulus length in bits
const FF_BITS int = (BIGBITS * FFLEN) /* Finite Field Size in bits - must be 256.2^n */
const HFLEN int = (FFLEN / 2)         /* Useful for half-size RSA private key operations */

const P_MBITS uint = MODBYTES * 8
const P_OMASK Chunk = (Chunk(-1) << (P_MBITS % BASEBITS))
const P_FEXCESS Chunk = (Chunk(1) << (BASEBITS*uint(NLEN) - P_MBITS - 1))
const P_TBITS uint = (P_MBITS % BASEBITS)

