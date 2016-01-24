
package math3d
import (
//"fmt"
	"fmt"
)

type Triangle struct {

	Vertices 	[]*Vertex
	uv1	Vector2
	uv2 Vector2
	uv3 Vector2

	Id 			int
	Normal		Vector3
	Removed		bool
}



func (t Triangle) SetUV(uv1 Vector2, uv2 Vector2, uv3 Vector2) {
	t.uv1 = uv1
	t.uv2 = uv2
	t.uv3 = uv3
}

// Pointer based comparison (will it work?)

func (t Triangle) HasVertex(tv *Vertex) (bool) {

	for i:=0; i< len(t.Vertices); i++ {
		if t.Vertices[i].Id == tv.Id {
			return true
		}
	}

	return false
}




func (this Triangle) ReplaceVertex(vOld *Vertex, vNew *Vertex, debug bool)() {

	worked := false

	for t:=0; t<len(this.Vertices); t++ {
		if (debug) {
			fmt.Printf("(%d) OLD:NEW = %d:%d. .... Now checking %d\r\n", t, vOld.Id, vNew.Id, this.Vertices[t].Id)
		}
		if this.Vertices[t].Id == vOld.Id {
			fmt.Println("**YES** found a match and replace the vertex")
			*this.Vertices[t] = *vNew
			worked = true
			break
		}
	}

	if debug {
		worked = worked
//		fmt.Printf("Vertex. %d  ->  %d\r\n was stuff replaced: %t\r\n", vOld.Id, vNew.Id, worked)
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

