package webserver

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/usecases"
	"github.com/pkg/errors"
)

const fixedPort = ":10010"

type Server struct {
	server   *http.Server
	storage  *infra.Storage
	commands map[string]func(storage *infra.Storage, args []string) []byte
}

func NewHttpServer(storage *infra.Storage, errorChan chan error) *Server {
	mux := http.NewServeMux()

	println("Starting HTTP server on " + fixedPort)

	server := Server{
		server: &http.Server{
			Addr:    fixedPort,
			Handler: mux,
		},
		storage:  storage,
		commands: usecases.Commands(),
	}

	mux.HandleFunc("/", server.handler)

	go func(errorChan chan error) {
		if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			println("Could not start server: %v\n", err)
			errorChan <- err
		}
	}(errorChan)

	return &server
}

// StopHTTPServer gracefully shuts down the server with a timeout.
func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	println(r.Method, r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	commandQueue, err := enqueueCommand(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response []byte

	if r.Method == http.MethodGet {
		command := commandQueue[1]
		commandHandler, ok := s.commands[command]
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		response = commandHandler(s.storage, commandQueue)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func enqueueCommand(req *http.Request) ([]string, error) {
	command := req.RequestURI[1:]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}
	defer req.Body.Close()
	return []string{"", command, string(body)}, nil
}
