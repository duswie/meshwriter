package meshwriter

import (
	"encoding/binary"
	"io"
	"math"
)

func writeStlHeader(w io.Writer, n_faces uint32, header_text ...string) {
	headerBuf := make([]byte, 84)
	text := "Generated with duswie/meshwriter"
	if len(header_text) > 0 {
		text = header_text[0]
	}

	copy(headerBuf, text)
	binary.LittleEndian.PutUint32(headerBuf[80:84], uint32(n_faces))
	w.Write(headerBuf)

}

func encodePoint(buf []byte, offset *int, pt [3]float32) {
	encodeFloat32(buf, offset, pt[0])
	encodeFloat32(buf, offset, pt[1])
	encodeFloat32(buf, offset, pt[2])
}

func encodeFloat32(buf []byte, offset *int, f float32) {
	bits := math.Float32bits(f)
	binary.LittleEndian.PutUint32(buf[*offset:(*offset)+4], bits)
	*offset += 4
}

func encodeUint16(buf []byte, offset *int, u uint16) {
	binary.LittleEndian.PutUint16(buf[*offset:(*offset)+2], u)
	*offset += 2
}

func arraySub(a [3]float32, b [3]float32) [3]float32 {
	return [3]float32{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func getNormal(p [3][3]float32) [3]float32 {
	var n [3]float32
	u := arraySub(p[1], p[0])
	v := arraySub(p[2], p[0])
	n[0] = ((u[1] * v[2]) - (u[2] * v[1]))
	n[1] = ((u[2] * v[0]) - (u[0] * v[2]))
	n[2] = ((u[0] * v[1]) - (u[1] * v[0]))
	return n
}
