// +nirvana:api=descriptors:"Descriptor"

package apis

import (
	"github.com/dyweb/sundial/pkg/apis/middlewares"
	v1 "github.com/dyweb/sundial/pkg/apis/v1/descriptors"

	def "github.com/caicloud/nirvana/definition"
)

// Descriptor returns a combined descriptor for APIs of all versions.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Description: "APIs",
		Path:        "/apis",
		Middlewares: middlewares.Middlewares(),
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Children: []def.Descriptor{
			v1.Descriptor(),
		},
	}
}
