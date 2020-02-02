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

/* Fixed Data in ROM - Field and Curve parameters */

package BN254

//BN254 Curve

// Base Bits= 56
var Modulus = [...]Chunk{0x13, 0x13A7, 0x80000000086121, 0x40000001BA344D, 0x25236482}
var ROI = [...]Chunk{0x12, 0x13A7, 0x80000000086121, 0x40000001BA344D, 0x25236482}
var R2modp = [...]Chunk{0x2F2A96FF5E7E39, 0x64E8642B96F13C, 0x9926F7B00C7146, 0x8321E7B4DACD24, 0x1D127A2E}

const MConst Chunk = 0x435E50D79435E5

const CURVE_Cof_I int = 1
const CURVE_A int = 0
const CURVE_B_I int = 2

var CURVE_B = [...]Chunk{0x2, 0x0, 0x0, 0x0, 0x0}
var CURVE_Order = [...]Chunk{0xD, 0x800000000010A1, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}
var CURVE_Gx = [...]Chunk{0x12, 0x13A7, 0x80000000086121, 0x40000001BA344D, 0x25236482}
var CURVE_Gy = [...]Chunk{0x1, 0x0, 0x0, 0x0, 0x0}

var Fra = [...]Chunk{0x7DE6C06F2A6DE9, 0x74924D3F77C2E1, 0x50A846953F8509, 0x212E7C8CB6499B, 0x1B377619}
var Frb = [...]Chunk{0x82193F90D5922A, 0x8B6DB2C08850C5, 0x2F57B96AC8DC17, 0x1ED1837503EAB2, 0x9EBEE69}
var CURVE_Bnx = [...]Chunk{0x80000000000001, 0x40, 0x0, 0x0, 0x0}
var CURVE_Cof = [...]Chunk{0x1, 0x0, 0x0, 0x0, 0x0}
var CURVE_Cru = [...]Chunk{0x80000000000007, 0x6CD, 0x40000000024909, 0x49B362, 0x0}
var CURVE_Pxa = [...]Chunk{0xEE4224C803FB2B, 0x8BBB4898BF0D91, 0x7E8C61EDB6A464, 0x519EB62FEB8D8C, 0x61A10BB}
var CURVE_Pxb = [...]Chunk{0x8C34C1E7D54CF3, 0x746BAE3784B70D, 0x8C5982AA5B1F4D, 0xBA737833310AA7, 0x516AAF9}
var CURVE_Pya = [...]Chunk{0xF0E07891CD2B9A, 0xAE6BDBE09BD19, 0x96698C822329BD, 0x6BAF93439A90E0, 0x21897A0}
var CURVE_Pyb = [...]Chunk{0x2D1AEC6B3ACE9B, 0x6FFD739C9578A, 0x56F5F38D37B090, 0x7C8B15268F6D44, 0xEBB2B0E}
var CURVE_W = [2][5]Chunk{{0x3, 0x80000000000204, 0x6181, 0x0, 0x0}, {0x1, 0x81, 0x0, 0x0, 0x0}}
var CURVE_SB = [2][2][5]Chunk{{{0x4, 0x80000000000285, 0x6181, 0x0, 0x0}, {0x1, 0x81, 0x0, 0x0, 0x0}}, {{0x1, 0x81, 0x0, 0x0, 0x0}, {0xA, 0xE9D, 0x80000000079E1E, 0x40000001BA344D, 0x25236482}}}
var CURVE_WB = [4][5]Chunk{{0x80000000000000, 0x80000000000040, 0x2080, 0x0, 0x0}, {0x80000000000005, 0x54A, 0x8000000001C707, 0x312241, 0x0}, {0x80000000000003, 0x800000000002C5, 0xC000000000E383, 0x189120, 0x0}, {0x80000000000001, 0x800000000000C1, 0x2080, 0x0, 0x0}}
var CURVE_BB = [4][4][5]Chunk{{{0x8000000000000D, 0x80000000001060, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}, {0x8000000000000C, 0x80000000001060, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}, {0x8000000000000C, 0x80000000001060, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}, {0x2, 0x81, 0x0, 0x0, 0x0}}, {{0x1, 0x81, 0x0, 0x0, 0x0}, {0x8000000000000C, 0x80000000001060, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}, {0x8000000000000D, 0x80000000001060, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}, {0x8000000000000C, 0x80000000001060, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}}, {{0x2, 0x81, 0x0, 0x0, 0x0}, {0x1, 0x81, 0x0, 0x0, 0x0}, {0x1, 0x81, 0x0, 0x0, 0x0}, {0x1, 0x81, 0x0, 0x0, 0x0}}, {{0x80000000000002, 0x40, 0x0, 0x0, 0x0}, {0x2, 0x102, 0x0, 0x0, 0x0}, {0xA, 0x80000000001020, 0x8000000007FF9F, 0x40000001BA344D, 0x25236482}, {0x80000000000002, 0x40, 0x0, 0x0, 0x0}}}

