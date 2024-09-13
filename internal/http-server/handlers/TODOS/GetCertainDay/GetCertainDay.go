package getcertainday

import (
	gettoday "TODO_App/internal/http-server/handlers/TODOS/GetToday"
	"TODO_App/todo"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Response struct {
	Todos []todo.TODO `json:"todos"`
	Error string      `json:"status,omitempty"`
}

func GetCertain(log *slog.Logger, getter gettoday.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		daystr := chi.URLParam(r, "day")
		const op = "handlers.TODOS.GetCertain"
		log = log.With(
			slog.String("op", op),
			slog.String("Request_ID", middleware.GetReqID(r.Context())),
		)

		todos, err := getter.GetTODOS(daystr)
		if err != nil {
			log.Error("Some error in %s: %s", op, err)

			render.JSON(w, r, Response{
				Todos: nil,
				Error: err.Error(),
			})
			return
		}

		log.Info("Requested for todos", slog.Int("Count of todos", len(todos)), slog.String("day", daystr))

		render.JSON(w, r, Response{
			Todos: todos,
		})
	}
}
