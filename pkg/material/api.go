// Package material implements the different materials and their properties.
package material

import (
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/vec3"
)

// Material defines the methods to handle materials.
type Material interface {
	Scatter(r ray.Ray, hr *hitrecord.HitRecord) (*ray.RayImpl, *vec3.Vec3Impl, bool)
}
