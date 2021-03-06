package material

import (
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/hitrecord"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/ray"
	"github.com/flynn-nrg/raytracing-in-one-weekend/pkg/vec3"
)

// Ensure interface compliance.
var _ Material = (*Metal)(nil)

// Metal represents metallic materials.
type Metal struct {
	albedo *vec3.Vec3Impl
	fuzz   float64
}

// NewMetal returns an instance of the metal material.
func NewMetal(albedo *vec3.Vec3Impl, fuzz float64) *Metal {
	return &Metal{
		albedo: albedo,
		fuzz:   fuzz,
	}
}

// Scatter computes how the ray bounces off the surface of a metallic object.
func (m *Metal) Scatter(r ray.Ray, hr *hitrecord.HitRecord) (*ray.RayImpl, *vec3.Vec3Impl, bool) {
	reflected := reflect(vec3.UnitVector(r.Direction()), hr.Normal())
	scattered := ray.New(hr.P(), vec3.Add(reflected, vec3.ScalarMul(randomInUnitSphere(), m.fuzz)))
	attenuation := m.albedo
	return scattered, attenuation, (vec3.Dot(scattered.Direction(), hr.Normal()) > 0)
}
