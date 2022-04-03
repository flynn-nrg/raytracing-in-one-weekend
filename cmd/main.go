package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/camera"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/hitable"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
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

func color(r ray.Ray, world *hitable.HitableSlice) *vec3.Vec3Impl {
	if rec, ok := world.Hit(r, 0.001, math.MaxFloat64); ok {
		target := vec3.Add(vec3.Add(rec.P(), rec.Normal()), randomInUnitSphere())
		return vec3.ScalarMul(color(ray.New(rec.P(), vec3.Sub(target, rec.P())), world), 0.5)
	}
	unitDirection := vec3.UnitVector(r.Direction())
	t := 0.5*unitDirection.Y + 1.0
	return vec3.Add(vec3.ScalarMul(&vec3.Vec3Impl{X: 1.0, Y: 1.0, Z: 1.0}, (1.0-t)),
		vec3.ScalarMul(&vec3.Vec3Impl{X: 0.5, Y: 0.7, Z: 1.0}, t))
}

func main() {
	nx := 200
	ny := 100
	ns := 100

	fmt.Printf("P3\n%v %v\n255\n", nx, ny)

	world := hitable.NewSlice([]hitable.Hitable{
		hitable.NewSphere(&vec3.Vec3Impl{Z: -1}, 0.5),
		hitable.NewSphere(&vec3.Vec3Impl{Y: -100.5, Z: -1}, 100),
	})

	cam := camera.New()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := &vec3.Vec3Impl{}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				col = vec3.Add(col, color(r, world))
			}

			col = vec3.ScalarDiv(col, float64(ns))
			// gamma 2
			col = &vec3.Vec3Impl{X: math.Sqrt(col.X), Y: math.Sqrt(col.Y), Z: math.Sqrt(col.Z)}
			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)

			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
