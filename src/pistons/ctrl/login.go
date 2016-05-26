package ctrl

import (
	"html/template"
	"net/http"

	"github.com/golang-webapp-gorilla/src/pistons/vm"
)

type loginController struct {
	loginTemplate *template.Template
}

func (lc *loginController) GetLogin(w http.ResponseWriter, r *http.Request) {
	vmodel := vm.Base{}
	lc.loginTemplate.Execute(w, vmodel)
}
