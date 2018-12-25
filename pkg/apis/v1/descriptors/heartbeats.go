package descriptors

import (
	def "github.com/caicloud/nirvana/definition"
	"github.com/dyweb/sundial/pkg/handlers/heartbeats"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/users/{user}/heartbeats",
		Definitions: []def.Definition{postHeartBeats},
	}, {
		Path:        "/users/current/heartbeats",
		Definitions: []def.Definition{postCurrentHeartBeats},
	},
	}...)
}

var postHeartBeats = def.Definition{
	Method:      def.Create,
	Summary:     "Create a heartbeat for the given user",
	Description: "Create a heartbeat from the plugins for the given user",
	Function:    heartbeats.POSTHeartBeat,
	Parameters: []def.Parameter{
		def.PathParameterFor("user", "username"),
		def.BodyParameterFor("heartbeat"),
	},
	Results: def.DataErrorResults("A Heartbeat"),
}

var postCurrentHeartBeats = def.Definition{
	Method:      def.Create,
	Summary:     "Create a heartbeat for the current user",
	Description: "Create a heartbeat from the plugins for the current user",
	Function:    heartbeats.POSTCurrentHeartBeat,
	Parameters: []def.Parameter{
		def.BodyParameterFor("heartbeat"),
	},
	Results: def.DataErrorResults("A Heartbeat"),
}
