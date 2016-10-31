package idgen

import (
	"encoding/base64"
)

func dumpUint64(val uint64) []byte {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	buf[0] = byte((val & 0xff00000000000000) >> 56)
	buf[1] = byte((val & 0x00ff000000000000) >> 48)
	buf[2] = byte((val & 0x0000ff0000000000) >> 40)
	buf[3] = byte((val & 0x000000ff00000000) >> 32)
	buf[4] = byte((val & 0x00000000ff000000) >> 24)
	buf[5] = byte((val & 0x0000000000ff0000) >> 16)
	buf[6] = byte((val & 0x000000000000ff00) >> 8)
	buf[7] = byte((val & 0x00000000000000ff))

	return buf
}

func dumpUint32(val uint32) []byte {
	buf := []byte{0, 0, 0, 0}

	buf[0] = byte((val & 0xff000000) >> 24)
	buf[1] = byte((val & 0x00ff0000) >> 16)
	buf[2] = byte((val & 0x0000ff00) >> 8)
	buf[3] = byte((val & 0x000000ff))

	return buf
}

func encode(id uint16, time uint32, seq uint16, suffix uint32) string {
	var prefix uint64

	prefix = (uint64(id) << 48)
	prefix |= (uint64(seq) << 32)
	prefix |= uint64(time)

	return base64.URLEncoding.EncodeToString(append(dumpUint64(prefix), dumpUint32(suffix)...))
}
