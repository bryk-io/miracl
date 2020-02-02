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

/* MiotCL double length DBIG number class */

package C41417

import "strconv"

//import "fmt"

func NewDBIG() *DBIG {
	b := new(DBIG)
	for i := 0; i < DNLEN; i++ {
		b.w[i] = 0
	}
	return b
}

func NewDBIGcopy(x *DBIG) *DBIG {
	b := new(DBIG)
	for i := 0; i < DNLEN; i++ {
		b.w[i] = x.w[i]
	}
	return b
}

func NewDBIGscopy(x *BIG) *DBIG {
	b := new(DBIG)
	for i := 0; i < NLEN-1; i++ {
		b.w[i] = x.w[i]
	}
	b.w[NLEN-1] = x.get(NLEN-1) & BMASK /* top word normalized */
	b.w[NLEN] = x.get(NLEN-1) >> BASEBITS

	for i := NLEN + 1; i < DNLEN; i++ {
		b.w[i] = 0
	}
	return b
}

/* normalise this */
func (r *DBIG) norm() {
	carry := Chunk(0)
	for i := 0; i < DNLEN-1; i++ {
		d := r.w[i] + carry
		r.w[i] = d & BMASK
		carry = d >> BASEBITS
	}
	r.w[DNLEN-1] = (r.w[DNLEN-1] + carry)
}

/* split DBIG at position n, return higher half, keep lower half */
func (r *DBIG) split(n uint) *BIG {
	t := NewBIG()
	m := n % BASEBITS
	carry := r.w[DNLEN-1] << (BASEBITS - m)

	for i := DNLEN - 2; i >= NLEN-1; i-- {
		nw := (r.w[i] >> m) | carry
		carry = (r.w[i] << (BASEBITS - m)) & BMASK
		t.set(i-NLEN+1, nw)
	}
	r.w[NLEN-1] &= ((Chunk(1) << m) - 1)
	return t
}

func (r *DBIG) cmove(g *DBIG, d int) {
	var b = Chunk(-d)

	for i := 0; i < DNLEN; i++ {
		r.w[i] ^= (r.w[i] ^ g.w[i]) & b
	}
}

/* Compare a and b, return 0 if a==b, -1 if a<b, +1 if a>b. Inputs must be normalised */
func dcomp(a *DBIG, b *DBIG) int {
	gt:=Chunk(0)
	eq:=Chunk(1)
	for i := DNLEN - 1; i >= 0; i-- {
		gt |= ((b.w[i]-a.w[i]) >> BASEBITS) & eq
		eq &= ((b.w[i]^a.w[i])-1) >> BASEBITS
	}
	return int(gt+gt+eq-1)
}

/* Copy from another DBIG */
func (r *DBIG) copy(x *DBIG) {
	for i := 0; i < DNLEN; i++ {
		r.w[i] = x.w[i]
	}
}

/* Copy from another BIG to upper half */
func (r *DBIG) ucopy(x *BIG) {
	for i := 0; i < NLEN; i++ {
		r.w[i] = 0
	}
	for i := NLEN; i < DNLEN; i++ {
		r.w[i] = x.w[i-NLEN]
	}
}

func (r *DBIG) add(x *DBIG) {
	for i := 0; i < DNLEN; i++ {
		r.w[i] = r.w[i] + x.w[i]
	}
}

/* this-=x */
func (r *DBIG) sub(x *DBIG) {
	for i := 0; i < DNLEN; i++ {
		r.w[i] = r.w[i] - x.w[i]
	}
}

/* this-=x */
func (r *DBIG) rsub(x *DBIG) {
	for i := 0; i < DNLEN; i++ {
		r.w[i] = x.w[i] - r.w[i]
	}
}

/* general shift left */
func (r *DBIG) shl(k uint) {
	n := k % BASEBITS
	m := int(k / BASEBITS)

	r.w[DNLEN-1] = (r.w[DNLEN-1-m] << n) | (r.w[DNLEN-m-2] >> (BASEBITS - n))
	for i := DNLEN - 2; i > m; i-- {
		r.w[i] = ((r.w[i-m] << n) & BMASK) | (r.w[i-m-1] >> (BASEBITS - n))
	}
	r.w[m] = (r.w[0] << n) & BMASK
	for i := 0; i < m; i++ {
		r.w[i] = 0
	}
}

/* general shift right */
func (r *DBIG) shr(k uint) {
	n := (k % BASEBITS)
	m := int(k / BASEBITS)
	for i := 0; i < DNLEN-m-1; i++ {
		r.w[i] = (r.w[m+i] >> n) | ((r.w[m+i+1] << (BASEBITS - n)) & BMASK)
	}
	r.w[DNLEN-m-1] = r.w[DNLEN-1] >> n
	for i := DNLEN - m; i < DNLEN; i++ {
		r.w[i] = 0
	}
}

/* reduces this DBIG mod a BIG, and returns the BIG */
func (r *DBIG) mod(c *BIG) *BIG {
	r.norm()
	m := NewDBIGscopy(c)
	dr := NewDBIG()

	if dcomp(r, m) < 0 {
		return NewBIGdcopy(r)
	}

	m.shl(1)
	k := 1

	for dcomp(r, m) >= 0 {
		m.shl(1)
		k++
	}

	for k > 0 {
		m.shr(1)

		dr.copy(r)
		dr.sub(m)
		dr.norm()
		r.cmove(dr, int(1-((dr.w[DNLEN-1]>>uint(CHUNK-1))&1)))
		k--
	}
	return NewBIGdcopy(r)
}

/* return this/c */
func (r *DBIG) div(c *BIG) *BIG {
	var d int
	k := 0
	m := NewDBIGscopy(c)
	a := NewBIGint(0)
	e := NewBIGint(1)
	sr := NewBIG()
	dr := NewDBIG()
	r.norm()

	for dcomp(r, m) >= 0 {
		e.fshl(1)
		m.shl(1)
		k++
	}

	for k > 0 {
		m.shr(1)
		e.shr(1)

		dr.copy(r)
		dr.sub(m)
		dr.norm()
		d = int(1 - ((dr.w[DNLEN-1] >> uint(CHUNK-1)) & 1))
		r.cmove(dr, d)
		sr.copy(a)
		sr.add(e)
		sr.norm()
		a.cmove(sr, d)

		k--
	}
	return a
}

/* Convert to Hex String */
func (r *DBIG) toString() string {
	s := ""
	len := r.nbits()

	if len%4 == 0 {
		len /= 4
	} else {
		len /= 4
		len++

	}

	for i := len - 1; i >= 0; i-- {
		b := NewDBIGcopy(r)

		b.shr(uint(i * 4))
		s += strconv.FormatInt(int64(b.w[0]&15), 16)
	}
	return s
}

/* return number of bits */
func (r *DBIG) nbits() int {
	k := DNLEN - 1
	t := NewDBIGcopy(r)
	t.norm()
	for k >= 0 && t.w[k] == 0 {
		k--
	}
	if k < 0 {
		return 0
	}
	bts := int(BASEBITS) * k
	c := t.w[k]
	for c != 0 {
		c /= 2
		bts++
	}
	return bts
}

/* convert from byte array to BIG */
func DBIG_fromBytes(b []byte) *DBIG {
	m := NewDBIG()
	for i := 0; i < len(b); i++ {
		m.shl(8)
		m.w[0] += Chunk(int(b[i] & 0xff))
	}
	return m
}
