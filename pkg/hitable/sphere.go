package hitable

import (
	"math"

	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/vec3"
)

// Ensure interface compliance.
var _ Hitable = (*Sphere)(nil)

// Sphere represents a sphere in the 3d world.
type Sphere struct {
	center *vec3.Vec3Impl
	radius float64
}

// NewSphere returns a new instance of Sphere.
func NewSphere(center *vec3.Vec3Impl, radius float64) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
	}
}

// Hit computes whether a ray intersects with the defined sphere.
func (s *Sphere) Hit(r ray.Ray, tMin float64, tMax float64) (*HitRecord, bool) {
	oc := vec3.Sub(r.Origin(), s.center)
	a := vec3.Dot(r.Direction(), r.Direction())
	b := vec3.Dot(oc, r.Direction())
	c := vec3.Dot(oc, oc) - (s.radius * s.radius)

	discriminant := (b * b) - (a * c)
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return &HitRecord{
				t:      temp,
				p:      r.PointAtParameter(temp),
				normal: vec3.ScalarDiv(vec3.Sub(r.PointAtParameter(temp), s.center), s.radius),
			}, true
		}

		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return &HitRecord{
				t:      temp,
				p:      r.PointAtParameter(temp),
				normal: vec3.ScalarDiv(vec3.Sub(r.PointAtParameter(temp), s.center), s.radius),
			}, true
		}
	}

	return nil, false
}
