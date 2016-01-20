
package math3d

type Triangle struct {

	Vertices 	[3] *Vertex
	uv1	Vector2
	uv2 Vector2
	uv3 Vector2

	id 			int
	Normal		Vector3
}


func (this Triangle) SetVertices(v0 *Vertex, v1 *Vertex, v2 *Vertex) {
	this.Vertices[0] = v0
	this.Vertices[1] = v1
	this.Vertices[2] = v2
}

func (t Triangle) SetUV(uv1 Vector2, uv2 Vector2, uv3 Vector2) {
	t.uv1 = uv1
	t.uv2 = uv2
	t.uv3 = uv3
}

// Pointer based comparison (will it work?)

func (t Triangle) HasVertex(tv *Vector3) (bool) {

	for _,v:= range t.Vertices {

		if v.Position == tv {
			return true
		}
	}

	return false
}


func (this Triangle) ReplaceVertex(vOld *Vector3, vNew *Vector3, newTexUV *Vector2 )() {


	for _,v:= range this.Vertices {
		if v.Position == vOld {
			v.Position = vNew
			// update UV code here
		}
	}

	this.CalcNormal()

	return // old_uv coords see Triangle.js:111

}

func (this Triangle) CalcNormal() (Vector3) {

	v0 := this.Vertices[0].Position
	v1 := this.Vertices[1].Position
	v2 := this.Vertices[2].Position

	w1 := v1.Subtract(v0)
	w2 := v2.Subtract(v1)

	this.Normal = w1.Cross(w2)

	if this.Normal.Magnitude() == 0 {
		return this.Normal
	}

	this.Normal.Normalise()

	return this.Normal

}

/*

func (t Triangle) HasVertex(v Vector3) (bool) {

	if reflect.DeepEqual(t.p1, v) {
		return true
	}

	return true
}
*/