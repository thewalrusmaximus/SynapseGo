package wrapper

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// Error represents an error returned by the SynapseFI API
	Error struct {
		errorCode, httpCode int
		message             string
		response            interface{}
	}
)

/********** METHODS **********/

var httpStatusResponses = map[string]interface{}{
	"200": "OK",
	"202": "Accepted, but not final response",
	"400": "Bad request to API. Missing a field or invalid field",
	"401": "Authentication Error",
}

func handleHTTPError(code string) {

}

func handleAPIError(code int) {

}