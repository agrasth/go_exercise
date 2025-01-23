package api

import (
    "encoding/json"
    "net/http"
    "item-manager/internal/data"
    "github.com/gorilla/mux"
)

type Server struct {
    mux.Router
    storage *data.Storage
}

func NewServer(storage *data.Storage) *Server {
    s := &Server{
        Router:  *mux.NewRouter(),
        storage: storage,
    }
    s.routes()
    return s
}

func (s *Server) routes() {
    s.HandleFunc("/shopping-items", s.listShoppingItems()).Methods("GET")
    s.HandleFunc("/shopping-items", s.createShoppingItems()).Methods("POST")
    s.HandleFunc("/shopping-items/{id}", s.removeShoppingItem()).Methods("DELETE")
}

func (s *Server) createShoppingItems() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var item data.Item
        if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        item.ID = data.NewUUID()
        s.storage.AddItem(item)
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(item); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func (s *Server) listShoppingItems() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "application/json")
        if err := json.NewEncoder(w).Encode(s.storage.GetItems()); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func (s *Server) removeShoppingItem() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        s.storage.RemoveItem(id)
        w.WriteHeader(http.StatusNoContent)
    }
}