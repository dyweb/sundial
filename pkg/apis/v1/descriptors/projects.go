package descriptors

import (
	def "github.com/caicloud/nirvana/definition"

	"github.com/dyweb/sundial/pkg/handlers/projects"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/users/{user}/projects",
		Definitions: []def.Definition{getProjects},
	}, {
		Path:        "/users/current/projects",
		Definitions: []def.Definition{getCurrentProjects},
	},
	}...)
}

var getProjects = def.Definition{
	Method:      def.List,
	Summary:     "List projects for the given user",
	Description: "List of WakaTime projects for for the given user",
	Function:    projects.GetProjects,
	Parameters: []def.Parameter{
		def.PathParameterFor("user", "username"),
	},
	Results: def.DataErrorResults("A project"),
}

var getCurrentProjects = def.Definition{
	Method:      def.List,
	Summary:     "List projects for the current user",
	Description: "List of projects for the current user.",
	Function:    projects.GetCurrentProjects,
	Results:     def.DataErrorResults("A project"),
}
