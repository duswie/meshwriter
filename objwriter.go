package meshwriter

import (
	"fmt"
	"io"
)

func writeObjVertex(w io.Writer, v [3]float64) {
	line := fmt.Sprintf("v %f %f %f\n", v[0], v[1], v[2])
	w.Write([]byte(line))
}

func writeObjFace(w io.Writer, f [3]uint32) {
	line := fmt.Sprintf("f %d %d %d\n", f[0]+1, f[1]+1, f[2]+1)
	w.Write([]byte(line))
}

//Write a waveform obj to a io.Writer from slices of vertices and faces.
//Vertices is a slice with all 3D points (X,Y,Z)
//Each face in the faces slice defines on triangle by referencing 3 vertices
//the three vertices of a triangle should be ordered counterclockwise from the outter view
func WriteObj(writer io.Writer, vertices [][3]float64, faces [][][3]uint32) error {

	head := "#Generated with duswie/meshwriter\n"
	writer.Write([]byte(head))
	for _, v := range vertices {
		writeObjVertex(writer, v)
	}
	for i, o := range faces {
		writer.Write([]byte(fmt.Sprintf("o Group%d\n", i+1)))
		for _, face := range o {
			writeObjFace(writer, face)
		}
	}

	return nil
}
