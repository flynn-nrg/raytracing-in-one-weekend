// Package hitable implements the methods used to compute intersections between a ray and geometry.
package hitable

import (
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/material"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
)

// Hitable defines the methods compute ray/geometry operations.
type Hitable interface {
	Hit(r ray.Ray, tMin float64, tMax float64) (*hitrecord.HitRecord, material.Material, bool)
}
