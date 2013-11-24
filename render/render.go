package render

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"net/http"
)

const (
	ContentType = "Content-Type"
	ContentJSON = "application/json"
)

type Render interface {
	JSON(status int, v interface{})
}

func Renderer() martini.Handler {
	return func(res http.ResponseWriter, c martini.Context) {
		c.MapTo(&renderer{res}, (*Render)(nil))
	}
}

type renderer struct {
	http.ResponseWriter
}

func (r *renderer) JSON(status int, v interface{}) {
	r.Header().Set(ContentType, ContentJSON)
	r.WriteHeader(status)

	result, err := json.Marshal(v)
	if err != nil {
		http.Error(r, err.Error(), 500)
	}

	r.Write(result)
}
