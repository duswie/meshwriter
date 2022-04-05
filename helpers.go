package meshwriter

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
