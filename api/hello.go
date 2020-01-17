package api

import (
	"github.com/borderstech/requiem"
)

// HelloController provides a Hello World API
type HelloController struct {
}

// HelloMessage contains a message
type HelloMessage struct {
	Message string `json:"message"`
}

func (c HelloController) hello(ctx requiem.HTTPContext) {
	m := &HelloMessage{Message: "Hello, world!"}
	ctx.SendJSON(m)
}

// Load attaches the HelloController to the given router
func (c HelloController) Load(router *requiem.Router) {
	r := router.NewAPIRouter("/hello")
	r.Get("", c.hello)
}
