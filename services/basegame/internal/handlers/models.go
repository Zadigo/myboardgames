package handlers

import "github.com/Zadigo/basegame/internal/models"

type BaseHandler struct {
	app models.AppInterface
}

func (b *BaseHandler) SetApp(app models.AppInterface) {
	b.app = app
}

func (b *BaseHandler) GetApp() models.AppInterface {
	return b.app
}
