package settings

import "strings"

var baseURI = "http:"

func SetBaseURI(addr string) {
	parts := strings.SplitN(addr, ":", 2)

	if parts[1] == "80" || parts[1] == "http" {
		baseURI = "http://localhost/"
	} else {
		baseURI = "http://localhost:" + parts[1] + "/"
	}
}

func BaseURI() string {
	return baseURI
}
