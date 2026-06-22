package handlers

import "github.com/Zadigo/flipseven/internal/models"

type BaseHandler struct {
	app models.AppInterface
}

func (b *BaseHandler) SetApp(app models.AppInterface) {
	b.app = app
}
