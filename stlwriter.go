package meshwriter

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

func writeStlHeader(w io.Writer, n_faces uint32, header_text ...string) error {
	headerBuf := make([]byte, 84)
	text := "Generated with duswie/meshwriter"
	if len(header_text) > 0 {
		text = header_text[0]
	}

	if len(text) > 79 {
		return errors.New("header text to long, can't be longer than 79 bytes")
	}

	copy(headerBuf, text)
	binary.LittleEndian.PutUint32(headerBuf[80:84], uint32(n_faces))
	w.Write(headerBuf)

	return nil

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

func writeStlFace(i uint32, w io.Writer, vertices [][3]float64, faces [][3]uint32) {

	var points [3][3]float32
	for n := 0; n < 3; n++ {
		points[n] = [3]float32{float32(vertices[faces[i][n]][0]), float32(vertices[faces[i][n]][1]), float32(vertices[faces[i][n]][2])}
	}
	buf := encodeStlFace(points)
	w.Write(buf[:])
}

func encodeStlFace(points [3][3]float32) [50]byte {
	var buf [50]byte
	offset := 0
	encodePoint(buf[:], &offset, getNormal(points))
	encodePoint(buf[:], &offset, points[0])
	encodePoint(buf[:], &offset, points[1])
	encodePoint(buf[:], &offset, points[2])
	encodeUint16(buf[:], &offset, 0)
	return (buf)
}

//Write a binary stl to a io.Writer from slices of vertices and faces.
//Vertices is a slice with all 3D points (X,Y,Z)
//Each face in the faces slice defines on triangle by referencing 3 vertices
//the three vertices of a triangle should be ordered counterclockwise from the outter view
func WriteBinaryStl(writer io.Writer, vertices [][3]float64, faces [][3]uint32) error {

	writeStlHeader(writer, uint32(len(faces)))

	for i := uint32(0); i < uint32(len(faces)); i++ {
		writeStlFace(i, writer, vertices, faces)
	}

	return nil
}
