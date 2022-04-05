# meshwriter
A golang package for exporting triangle meshes. Currently, binary stl and waveform obj files are supported.

## Usage
### STL files
Write a stl file from slices of vertices and faces.
`vertices` is a slice with all 3D points (X,Y,Z) of the geometry 
Each face in the `faces` slice defines on triangle by referencing 3 vertices by there indices.
The vertices of a triangle should be ordered counterclockwise from the outer view.
```go
import (
	"bufio"
	"os"
	"github.com/duswie/meshwriter"
)

//defining a simple tetraeder geometry
vertices := [][3]float64{{0, 0, 0}, {0, 3, 0}, {3, 0, 0}, {1.5, 1.5, 3}}
faces := [][3]uint32{{0, 2, 1}, {0, 1, 3}, {1, 2, 3}, {0, 3, 2}}

//creating a new file at temp dir
file, _ := os.Create(os.TempDir() + "stl_output_example.stl")

//using a write buffer for increased performance
file_buf := bufio.NewWriter(file)

//write the binary stl
meshwriter.WriteBinaryStl(file_buf, vertices, faces)

//flush buffer and close file
file_buf.Flush()
file.Close()
```

### OBJ files
Write a stl file from slices of vertices and faces.
`vertices` is a slice with all 3D points (X,Y,Z) of the geometry 
Each face in the `faces` slice defines on triangle by referencing 3 vertices by there indices.
The `faces` can be grouped to multiple slices, wich results in an obj file with multiple objects.
The vertices of a triangle should be ordered counterclockwise from the outer view
```go
//defining 2 simple tetraeder geometrys in different groups
vertices := [][3]float64{{0, 0, 0}, {0, 3, 0}, {3, 0, 0}, {1.5, 1.5, 3}, {5, 5, 0}, {5, 8, 0}, {8, 5, 0}, {6.5, 6.5, 3}}
faces := [][][3]uint32{
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
```