package binary

import (
	"encoding/binary"
	"errors"
	"io"
)

// Byte

func WriteByte(w io.Writer, b byte, n *int64, err *error) {
	WriteTo(w, []byte{b}, n, err)
}

func ReadByte(r io.Reader, n *int64, err *error) byte {
	buf := make([]byte, 1)
	ReadFull(r, buf, n, err)
	return buf[0]
}

// Int8

func WriteInt8(w io.Writer, i int8, n *int64, err *error) {
	WriteByte(w, byte(i), n, err)
}

func ReadInt8(r io.Reader, n *int64, err *error) int8 {
	return int8(ReadByte(r, n, err))
}

// UInt8

func WriteUInt8(w io.Writer, i uint8, n *int64, err *error) {
	WriteByte(w, byte(i), n, err)
}

func ReadUInt8(r io.Reader, n *int64, err *error) uint8 {
	return uint8(ReadByte(r, n, err))
}

// Int16

func WriteInt16(w io.Writer, i int16, n *int64, err *error) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(i))
	*n += 2
	WriteTo(w, buf, n, err)
}

func ReadInt16(r io.Reader, n *int64, err *error) int16 {
	buf := make([]byte, 2)
	ReadFull(r, buf, n, err)
	return int16(binary.LittleEndian.Uint16(buf))
}

// UInt16

func WriteUInt16(w io.Writer, i uint16, n *int64, err *error) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(i))
	*n += 2
	WriteTo(w, buf, n, err)
}

func ReadUInt16(r io.Reader, n *int64, err *error) uint16 {
	buf := make([]byte, 2)
	ReadFull(r, buf, n, err)
	return uint16(binary.LittleEndian.Uint16(buf))
}

// Int32

func WriteInt32(w io.Writer, i int32, n *int64, err *error) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	*n += 4
	WriteTo(w, buf, n, err)
}

func ReadInt32(r io.Reader, n *int64, err *error) int32 {
	buf := make([]byte, 4)
	ReadFull(r, buf, n, err)
	return int32(binary.LittleEndian.Uint32(buf))
}

// UInt32

func WriteUInt32(w io.Writer, i uint32, n *int64, err *error) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	*n += 4
	WriteTo(w, buf, n, err)
}

func ReadUInt32(r io.Reader, n *int64, err *error) uint32 {
	buf := make([]byte, 4)
	ReadFull(r, buf, n, err)
	return uint32(binary.LittleEndian.Uint32(buf))
}

// Int64

func WriteInt64(w io.Writer, i int64, n *int64, err *error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	*n += 8
	WriteTo(w, buf, n, err)
}

func ReadInt64(r io.Reader, n *int64, err *error) int64 {
	buf := make([]byte, 8)
	ReadFull(r, buf, n, err)
	return int64(binary.LittleEndian.Uint64(buf))
}

// UInt64

func WriteUInt64(w io.Writer, i uint64, n *int64, err *error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	*n += 8
	WriteTo(w, buf, n, err)
}

func ReadUInt64(r io.Reader, n *int64, err *error) uint64 {
	buf := make([]byte, 8)
	ReadFull(r, buf, n, err)
	return uint64(binary.LittleEndian.Uint64(buf))
}

// VarInt

func WriteVarInt(w io.Writer, i int64, n *int64, err *error) {
	buf := make([]byte, 9)
	*n += int64(binary.PutVarint(buf, int64(i)))
	WriteTo(w, buf, n, err)
}

func ReadVarInt(r io.Reader, n *int64, err *error) int64 {
	res, n_, err_ := readVarint(r)
	*n += n_
	*err = err_
	return res
}

// UVarInt

func WriteUVarInt(w io.Writer, i uint64, n *int64, err *error) {
	buf := make([]byte, 9)
	*n += int64(binary.PutUvarint(buf, uint64(i)))
	WriteTo(w, buf, n, err)
}

func ReadUVarInt(r io.Reader, n *int64, err *error) uint64 {
	res, n_, err_ := readUvarint(r)
	*n += n_
	*err = err_
	return res
}

//-----------------------------------------------------------------------------

var overflow = errors.New("binary: varint overflows a 64-bit integer")

// Modified to return number of bytes read, from
// http://golang.org/src/pkg/encoding/binary/varint.go?s=3652:3699#L116
func readUvarint(r io.Reader) (uint64, int64, error) {
	var x uint64
	var s uint
	var buf = make([]byte, 1)
	for i := 0; ; i++ {
		for {
			n, err := r.Read(buf)
			if err != nil {
				return x, int64(i), err
			}
			if n > 0 {
				break
			}
		}
		b := buf[0]
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return x, int64(i), overflow
			}
			return x | uint64(b)<<s, int64(i), nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
}

// Modified to return number of bytes read, from
// http://golang.org/src/pkg/encoding/binary/varint.go?s=3652:3699#L116
func readVarint(r io.Reader) (int64, int64, error) {
	ux, n, err := readUvarint(r) // ok to continue in presence of error
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x, n, err
}
