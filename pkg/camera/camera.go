// Package camera implements a set of functions to work with cameras.
package camera

import (
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
func New() *Camera {
	return &Camera{
		lowerLeftCorner: &vec3.Vec3Impl{X: -2.0, Y: -1.0, Z: -1.0},
		horizontal:      &vec3.Vec3Impl{X: 4.0},
		vertical:        &vec3.Vec3Impl{Y: 2.0},
		origin:          &vec3.Vec3Impl{},
	}
}

// GetRay returns the ray associated for the supplied u and v.
func (c *Camera) GetRay(u float64, v float64) *ray.RayImpl {
	return ray.New(c.origin,
		// lowerLeftCorner + u*horizontal + v*vertical - origin
		vec3.Sub(vec3.Add(c.lowerLeftCorner, vec3.Add(vec3.ScalarMul(c.horizontal, u),
			vec3.ScalarMul(c.vertical, v))), c.origin))
}
