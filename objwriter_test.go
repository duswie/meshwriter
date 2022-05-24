package meshwriter

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func TestWriteObjVertex(t *testing.T) {
	w := bytes.NewBufferString("")
	writeObjVertex(w, [3]float64{47.11, 123.1, 1.0 / 3.0})
	result, _ := w.ReadString('\n')
	want := "v 47.110000 123.100000 0.333333\n"
	if result != want {
		t.Errorf("Wrong vertex encoding: %s, want %s ", result, want)
	}

}

func TestWriteObFace(t *testing.T) {
	w := bytes.NewBufferString("")
	writeObjFace(w, [3]int{123, 456, 789})
	result, _ := w.ReadString('\n')
	want := "f 124 457 790\n"
	if result != want {
		t.Errorf("Wrong Face encoding: %s, want %s ", result, want)
	}

}

func ExampleWriteObj() {

	//defining 2 simple tetraeder geometry in different groups
	vertices := [][3]float64{{0, 0, 0}, {0, 3, 0}, {3, 0, 0}, {1.5, 1.5, 3}, {5, 5, 0}, {5, 8, 0}, {8, 5, 0}, {6.5, 6.5, 3}}
	faces := [][][3]int{
		{{0, 2, 1}, {0, 1, 3}, {1, 2, 3}, {0, 3, 2}},
		{{4, 6, 5}, {4, 5, 7}, {5, 6, 7}, {4, 7, 6}},
	}

	//creating a new file at temp dir
	file, _ := os.Create(os.TempDir() + "obj_output_example.obj")

	//using a write buffer for increased performance
	file_buf := bufio.NewWriter(file)

	//write the binary stl
	WriteObj(file_buf, vertices, faces)

	//flush buffer and close file
	file_buf.Flush()
	file.Close()

	println("Example OBJ file:" + os.TempDir() + "obj_output_example.obj")

	// Output:

}
