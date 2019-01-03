package descriptors

import (
	"github.com/dyweb/sundial/pkg/handlers/durations"

	def "github.com/caicloud/nirvana/definition"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/users/{user}/durations",
		Definitions: []def.Definition{getDurations},
	}, {
		Path:        "/users/current/durations",
		Definitions: []def.Definition{getCurrentUserDurations},
	},
	}...)
}

var getDurations = def.Definition{
	Method:      def.Get,
	Summary:     "Get a user's durations",
	Description: "A user's coding activity for the given day as an array of durations.",
	Function:    durations.GetDurations,
	Parameters: []def.Parameter{
		def.PathParameterFor("user", "Only shows durations for this user."),
		def.QueryParameterFor("date", "Requested day; Durations will be returned from 12am until 11:59pm in user's timezone for this day"),
		def.QueryParameterFor("project", "Only show durations for this project."),
		def.QueryParameterFor("branches", "Only show durations for these branches; comma separated list of branch names."),
	},
	Results: def.DataErrorResults("Durations"),
}

var getCurrentUserDurations = def.Definition{
	Method:      def.Get,
	Summary:     "Get current user's durations",
	Description: "Current user's coding activity for the given day as an array of durations.",
	Function:    durations.GetCurrentUserDurations,
	Parameters: []def.Parameter{
		def.QueryParameterFor("date", "Requested day; Durations will be returned from 12am until 11:59pm in user's timezone for this day"),
		def.QueryParameterFor("project", "Only show durations for this project."),
		def.QueryParameterFor("branches", "Only show durations for these branches; comma separated list of branch names."),
	},
	Results: def.DataErrorResults("Durations"),
}
