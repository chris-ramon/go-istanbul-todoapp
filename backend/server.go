package main

import (
	"net/http"
	"strconv"

	"github.com/chris-ramon/go-istanbul-todoapp/backend/corsutil"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"github.com/unrolled/render"
)

var (
	R     *render.Render
	Todos []*Todo
)

type Todo struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Status bool   `json:"status"`
}

func (t *Todo) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.Text:   "text",
		&t.Status: "status",
	}
}

func TodosIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	R.JSON(w, http.StatusOK, Todos)
}

func TodosCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := new(Todo)
	if errs := binding.Bind(r, todo); errs.Handle(w) {
		return
	}
	Todos = append(Todos, todo)
	R.JSON(w, http.StatusOK, todo)
}

func TodosUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todoId, _ := strconv.Atoi(p.ByName("id"))
	todo := getTodoById(todoId)
	form := new(Todo)
	if errs := binding.Bind(r, form); errs.Handle(w) {
		return
	}
	todo.Status = form.Status
	R.JSON(w, http.StatusOK, todo)
}

func getTodoById(id int) *Todo {
	for _, todo := range Todos {
		if todo.Id == id {
			return todo
		}
	}
	return &Todo{}
}

func init() {
	R = render.New()
	Todos = []*Todo{
		&Todo{Id: 1000, Text: "Visit Izmir"},
		&Todo{Id: 1001, Text: "Visit Ankara"},
	}
}

func main() {
	router := httprouter.New()
	n := negroni.Classic()
	router.GET("/todos", TodosIndex)
	router.POST("/todos", TodosCreate)
	router.PUT("/todos/:id", TodosUpdate)
	n.Use(negroni.HandlerFunc(corsutil.Handler))
	n.UseHandler(router)
	n.Run(":8080")
}
