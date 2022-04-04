// Package camera implements a set of functions to work with cameras.
package camera

import (
	"math"

	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/vec3"
)

// Camera represents a camera in the world.
type Camera struct {
	origin          *vec3.Vec3Impl
	lowerLeftCorner *vec3.Vec3Impl
	horizontal      *vec3.Vec3Impl
	vertical        *vec3.Vec3Impl
}

// New returns an instance of a camera.
func New(lookFrom *vec3.Vec3Impl, lookAt *vec3.Vec3Impl, vup *vec3.Vec3Impl, vfov float64, aspect float64) *Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight
	w := vec3.UnitVector(vec3.Sub(lookFrom, lookAt))
	u := vec3.UnitVector(vec3.Cross(vup, w))
	v := vec3.Cross(w, u)

	// origin - half_width*u - half_height*v - w
	lowerLeftCorner := vec3.Sub(lookFrom, vec3.ScalarMul(u, halfWidth), vec3.ScalarMul(v, halfHeight), w)
	horizontal := vec3.ScalarMul(u, 2.0*halfWidth)
	vertical := vec3.ScalarMul(v, 2.0*halfHeight)
	origin := lookFrom

	return &Camera{
		lowerLeftCorner: lowerLeftCorner,
		horizontal:      horizontal,
		vertical:        vertical,
		origin:          origin,
	}
}

// GetRay returns the ray associated for the supplied u and v.
func (c *Camera) GetRay(u float64, v float64) *ray.RayImpl {
	return ray.New(c.origin,
		// lowerLeftCorner + u*horizontal + v*vertical - origin
		vec3.Sub(vec3.Add(c.lowerLeftCorner, vec3.ScalarMul(c.horizontal, u),
			vec3.ScalarMul(c.vertical, v)), c.origin))
}
