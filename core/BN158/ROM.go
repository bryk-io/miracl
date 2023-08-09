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

/* Fixed Data in ROM - Field and Curve parameters */

package BN158

//BN158 Curve

// Base Bits= 56
var Modulus = [...]Chunk{0x72280AB04E013, 0x9953CF6FD3FB95, 0x24012027628C}
var R2modp = [...]Chunk{0xA60739B545973D, 0x52616565DB288, 0x1731500B226C}
var ROI = [...]Chunk{0x72280AB04E012, 0x9953CF6FD3FB95, 0x24012027628C}
var SQRTm3 = [...]Chunk{0x7F0E4048018004, 0x95F35CA99F4CE2, 0x2401202761FC}
var CRu = [...]Chunk{0x440A203181B007, 0x1B039631A5759, 0x48}

const MConst Chunk = 0xFC609004F615E5

var Fra = [...]Chunk{0xA1674296ECE2A9, 0xE1BF21C99296F2, 0xA85ECF82A02}
var Frb = [...]Chunk{0x65BB3E1417FD6A, 0xB794ADA64164A2, 0x197B332F3889}

// *** rom curve parameters *****
// Ate Bits= 42
// G2 Table size= 49
const CURVE_Cof_I int = 1

var CURVE_Cof = [...]Chunk{0x1, 0x0, 0x0}

const CURVE_B_I int = 5

var CURVE_B = [...]Chunk{0x5, 0x0, 0x0}
var CURVE_Order = [...]Chunk{0xD59F209F04200D, 0x9953CF6F73FA14, 0x24012027628C}
var CURVE_Gx = [...]Chunk{0x72280AB04E012, 0x9953CF6FD3FB95, 0x24012027628C}
var CURVE_Gy = [...]Chunk{0x2, 0x0, 0x0}
var CURVE_HTPC = [...]Chunk{0x1, 0x0, 0x0}

var CURVE_Bnx = [...]Chunk{0x4000801001, 0x0, 0x0}
var CURVE_Pxa = [...]Chunk{0x3B2765033A5768, 0x1EECE2B3022922, 0x1EA35F882728}
var CURVE_Pxb = [...]Chunk{0x7B04ACE776A2F5, 0x5D05BA314F9D68, 0x23485611EB92}
var CURVE_Pya = [...]Chunk{0x69AB26E30CFE24, 0x1FB7A85F92C435, 0x1C952F906B6E}
var CURVE_Pyb = [...]Chunk{0x91017738E8609D, 0x8445B3BA0F3EE2, 0x23E289544ED8}
var CURVE_W = [2][3]Chunk{{0x3182600A008003, 0x600180, 0x0}, {0x8001002001, 0x0, 0x0}}
var CURVE_SB = [2][2][3]Chunk{{{0x3182E00B00A004, 0x600180, 0x0}, {0x8001002001, 0x0, 0x0}}, {{0x8001002001, 0x0, 0x0}, {0xA41CC09503A00A, 0x9953CF6F13F894, 0x24012027628C}}}
var CURVE_WB = [4][3]Chunk{{0x10806002801000, 0x200080, 0x0}, {0xF907C026815005, 0x1202642519090, 0x30}, {0x7C84001380B003, 0x90132128C848, 0x18}, {0x1080E003803001, 0x200080, 0x0}}
var CURVE_BB = [4][4][3]Chunk{{{0xD59EE09E84100D, 0x9953CF6F73FA14, 0x24012027628C}, {0xD59EE09E84100C, 0x9953CF6F73FA14, 0x24012027628C}, {0xD59EE09E84100C, 0x9953CF6F73FA14, 0x24012027628C}, {0x8001002002, 0x0, 0x0}}, {{0x8001002001, 0x0, 0x0}, {0xD59EE09E84100C, 0x9953CF6F73FA14, 0x24012027628C}, {0xD59EE09E84100D, 0x9953CF6F73FA14, 0x24012027628C}, {0xD59EE09E84100C, 0x9953CF6F73FA14, 0x24012027628C}}, {{0x8001002002, 0x0, 0x0}, {0x8001002001, 0x0, 0x0}, {0x8001002001, 0x0, 0x0}, {0x8001002001, 0x0, 0x0}}, {{0x4000801002, 0x0, 0x0}, {0x10002004002, 0x0, 0x0}, {0xD59EA09E04000A, 0x9953CF6F73FA14, 0x24012027628C}, {0x4000801002, 0x0, 0x0}}}
