package meshwriter

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

const testfile = "meshwriter_stlout_testfile"

func getTestFile(filetype string) *os.File {
	filepath := os.TempDir() + testfile + "." + filetype
	w, _ := os.Create(filepath)
	return w
}

func delTestFile(filetype string) {
	filepath := os.TempDir() + testfile + "." + filetype
	os.Remove(filepath)
}

func TestWriteStlHeader(t *testing.T) {

	w := getTestFile("stl")
	writeStlHeader(w, 777777, "testtext")

	file := make([]byte, 100)
	w.Seek(0, 0)
	n, _ := w.Read(file)

	//test header size
	if n != 84 {
		t.Errorf("Wrong stl header size: %d, want 84", n)
	}

	//test face count
	want_faces := []byte{0x31, 0xDE, 0x0B, 0x00}
	if !bytes.Equal(file[80:84], want_faces) {
		t.Errorf("Wrong face count value: %b, want %b", file[80:84], want_faces)
	}

	//test face count
	want_text := "testtext"
	if !bytes.Equal(file[0:8], []byte(want_text)) {
		t.Errorf("Wrong header text: %s, want %s", file[0:8], want_text)
	}

	w.Close()
	delTestFile("stl")

}

func TestWriteStlFace(t *testing.T) {

	vertices := [][3]float64{{1, 1, 1}, {1, 2, 1}, {2, 1, 1}}
	faces := [][3]uint32{{0, 1, 2}}

	w := getTestFile("stl")

	writeStlFace(0, w, vertices, faces)

	file := make([]byte, 100)
	w.Seek(0, 0)
	n, _ := w.Read(file)

	//test header size
	if n != 50 {
		t.Errorf("Wrong stl header size: %d, want 50", n)
	}

}

func BenchmarkEncodeStlFace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeStlFace([3][3]float32{{1, 1, 1}, {1, 2, 1}, {2, 1, 1}})
	}
}

func ExampleWriteBinaryStl() {

	//defining a simple tetraeder geometry
	vertices := [][3]float64{{0, 0, 0}, {0, 3, 0}, {3, 0, 0}, {1.5, 1.5, 3}}
	faces := [][3]uint32{{0, 2, 1}, {0, 1, 3}, {1, 2, 3}, {0, 3, 2}}

	//creating a new file at temp dir
	file, _ := os.Create(os.TempDir() + "stl_output_example.stl")

	//using a write buffer for increased performance
	file_buf := bufio.NewWriter(file)

	//write the binary stl
	WriteBinaryStl(file_buf, vertices, faces)

	//flush buffer and close file
	file_buf.Flush()
	file.Close()

	println("Example STL file:" + os.TempDir() + "stl_output_example.stl")

	// Output:

}
