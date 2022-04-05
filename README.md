# meshwriter
A golang package for exporting triangle meshes. Currently, only binary stl is supported.

## Usage
Write a stl file from slices of vertices and faces.
`vertices` is a slice with all 3D points (X,Y,Z) of the geometry 
Each face in the `faces` slice defines on triangle by referencing 3 vertices by there indices.
The vertices of a triangle should be ordered counterclockwise from the outer view
```go
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
```