package hitable

import (
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
)

// Ensure interface compliance.
var _ Hitable = (*HitableSlice)(nil)

// HitableSlice represents a list of hitable entities.
type HitableSlice struct {
	hitables []Hitable
}

// NewSlice returns an instance of HitableSlice.
func NewSlice(hitables []Hitable) *HitableSlice {
	return &HitableSlice{
		hitables: hitables,
	}
}

// Hit computes whether a ray intersects with any of the elements in the slice.
func (hs *HitableSlice) Hit(r ray.Ray, tMin float64, tMax float64) (*HitRecord, bool) {
	var rec *HitRecord
	var hitAnything bool
	closestSoFar := tMax

	for _, h := range hs.hitables {
		if tempRec, ok := h.Hit(r, tMin, closestSoFar); ok {
			rec = tempRec
			hitAnything = ok
			closestSoFar = rec.t
		}
	}

	return rec, hitAnything
}
