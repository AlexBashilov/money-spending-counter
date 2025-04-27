package restapi

import _ "embed"

// go:embed v1/api.swagger.json
var data []byte

func SwaggerFile() []byte {
	return data
}
