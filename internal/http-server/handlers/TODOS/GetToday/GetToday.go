package gettoday

import (
	"TODO_App/todo"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	Todos []todo.TODO `json:"todos"`
	Error string      `json:"error,omitempty"`
}

type Getter interface {
	GetTODOS(day string) ([]todo.TODO, error)
}

func Get(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.TODOS.GetToday.Get"
		log = log.With(
			slog.String("op", op),
			slog.String("Request_ID", middleware.GetReqID(r.Context())),
		)
		day := time.Now().Format("2006-01-02")
		fmt.Println(day)
		todos, err := getter.GetTODOS(day)
		if err != nil {
			log.Error("Some error in %s: %s", op, err)

			render.JSON(w, r, Response{
				Todos: nil,
				Error: err.Error(),
			})
			return
		}

		log.Info("Requested for todays", slog.Int("Count of todos", len(todos)))

		render.JSON(w, r, Response{
			Todos: todos,
		})

	}
}
