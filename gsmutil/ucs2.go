package gsmutil

/* The MIT License (MIT)
Copyright © 2016 Maxim Kupriianov <max@kc.vc>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// from https://github.com/xlab/at/blob/master/pdu/ucs2.go
import (
	"errors"
	"unicode/utf16"
)

// ErrUnevenNumber happens when the number of octets (bytes) in the input is uneven.
var ErrUnevenNumber = errors.New("decode ucs2: uneven number of octets")

// EncodeUcs2 encodes the given UTF-8 text into UCS2 (UTF-16) encoding and returns the produced octets.
func EncodeUcs2(str string) []byte {
	buf := utf16.Encode([]rune(str))
	octets := make([]byte, 0, len(buf)*2)
	for _, n := range buf {
		octets = append(octets, byte(n&0xFF00>>8), byte(n&0x00FF))
	}
	return octets
}

// DecodeUcs2 decodes the given UCS2 (UTF-16) octet data into a UTF-8 encoded string.
func DecodeUcs2(octets []byte) (str string, err error) {
	if len(octets)%2 != 0 {
		err = ErrUnevenNumber
		return
	}
	buf := make([]uint16, 0, len(octets)/2)
	for i := 0; i < len(octets); i += 2 {
		buf = append(buf, uint16(octets[i])<<8|uint16(octets[i+1]))
	}
	runes := utf16.Decode(buf)
	return string(runes), nil
}
