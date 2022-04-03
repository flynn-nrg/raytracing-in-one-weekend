package material

import (
	"math/rand"

	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/vec3"
)

func randomInUnitSphere() *vec3.Vec3Impl {
	for {
		p := vec3.Sub(vec3.ScalarMul(&vec3.Vec3Impl{X: rand.Float64(), Y: rand.Float64(), Z: rand.Float64()}, 2.0),
			&vec3.Vec3Impl{X: 1.0, Y: 1.0, Z: 1.0})
		if p.SquaredLength() < 1.0 {
			return p
		}
	}
}

func reflect(v *vec3.Vec3Impl, n *vec3.Vec3Impl) *vec3.Vec3Impl {
	// v - 2*dot(v,n)*n
	return vec3.Sub(v, vec3.ScalarMul(n, 2*vec3.Dot(v, n)))
}
