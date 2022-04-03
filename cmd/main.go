package main

import (
	"fmt"
	"math"

	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/hitable"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/vec3"
)

func color(r ray.Ray, world *hitable.HitableSlice) *vec3.Vec3Impl {
	if rec, ok := world.Hit(r, 0.0, math.MaxFloat64); ok {
		return vec3.ScalarMul(&vec3.Vec3Impl{X: rec.Normal().X + 1, Y: rec.Normal().Y + 1, Z: rec.Normal().Z + 1}, 0.5)
	}
	unitDirection := vec3.UnitVector(r.Direction())
	t := 0.5*unitDirection.Y + 1.0
	return vec3.Add(vec3.ScalarMul(&vec3.Vec3Impl{X: 1.0, Y: 1.0, Z: 1.0}, (1.0-t)),
		vec3.ScalarMul(&vec3.Vec3Impl{X: 0.5, Y: 0.7, Z: 1.0}, t))
}

func main() {
	nx := 200
	ny := 100
	fmt.Printf("P3\n%v %v\n255\n", nx, ny)

	lowerLeftCorner := &vec3.Vec3Impl{X: -2.0, Y: -1.0, Z: -1.0}
	horizontal := &vec3.Vec3Impl{X: 4.0}
	vertical := &vec3.Vec3Impl{Y: 2.0}
	origin := &vec3.Vec3Impl{}

	world := hitable.NewSlice([]hitable.Hitable{
		hitable.NewSphere(&vec3.Vec3Impl{Z: -1}, 0.5),
		hitable.NewSphere(&vec3.Vec3Impl{Y: -100.5, Z: -1}, 100),
	})
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			r := ray.New(origin, vec3.Add(lowerLeftCorner, vec3.Add(vec3.ScalarMul(horizontal, u), vec3.ScalarMul(vertical, v))))
			col := color(r, world)

			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)

			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
