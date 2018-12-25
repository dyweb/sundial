package descriptors

import (
	"github.com/dyweb/sundial/pkg/handlers/users"

	def "github.com/caicloud/nirvana/definition"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/users/{user}",
		Definitions: []def.Definition{getUser},
	}, {
		Path:        "/users/current",
		Definitions: []def.Definition{getCurrentUser},
	},
	}...)
}

var getUser = def.Definition{
	Method:      def.List,
	Summary:     "Get a user",
	Description: "Get the user information for the given username",
	Function:    users.GetUser,
	Parameters: []def.Parameter{
		def.PathParameterFor("user", "username"),
	},
	Results: def.DataErrorResults("An user"),
}

var getCurrentUser = def.Definition{
	Method:      def.List,
	Summary:     "Get current user",
	Description: "Get the current user information",
	Function:    users.GetCurrentUser,
	Results:     def.DataErrorResults("An user"),
}
