// Package hitable implements the methods used to compute intersections between a ray and geometry.
package hitable

import (
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/vec3"
)

// Hitable defines the methods compute ray/geometry operations.
type Hitable interface {
	Hit(r ray.Ray, tMin float64, tMax float64) (*HitRecord, bool)
}

// HitRecord contains data related to an intersection between a ray and an object.
type HitRecord struct {
	t      float64
	p      *vec3.Vec3Impl
	normal *vec3.Vec3Impl
}

// Normal returns the normal vector at the intersection point.
func (hr *HitRecord) Normal() *vec3.Vec3Impl {
	return hr.normal
}
