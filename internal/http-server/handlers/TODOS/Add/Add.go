package add

import (
	"TODO_App/todo"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Event string `json:"event" validate:"required"`
	Day   string `json:"day" validate:"required"`
	Time  string `json:"time" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type EventAddInt interface {
	AddTODO(event todo.TODO) (int64, error)
}

func New(log *slog.Logger, Adder EventAddInt) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.TODOS.Add.New"
		log = log.With(
			slog.String("op", op),
			slog.String("Request_ID", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request", slog.String("error: %s", err.Error()))

			render.JSON(w, r, Response{
				Status: "Error",
				Error:  err.Error(),
			})

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("Invalid request", slog.String("error: %s", err.Error()))

			render.JSON(w, r, Response{
				Status: "Error",
				Error:  err.Error(),
			})

			return
		}

		id, err := Adder.AddTODO(todo.NewTODO(req.Event, req.Day, req.Time))

		if err != nil {
			log.Error("Failed to add event", slog.String("error: %s", err.Error()))

			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Failed to add new event",
			})

			return
		}

		log.Info("Event added", slog.Int64("id: ", id))

		render.JSON(w, r, Response{
			Status: "OK",
		})
	}
}
