package handlers

import (
	"github.com/AykoSousa/uncle-cli/cmd/config"
)

func OpenAPIResponse(responseModel string) {
	config.SetResponseColorTheme(responseModel)
}