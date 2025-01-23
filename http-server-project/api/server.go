package api

import (
	"net/http"
	"encoding/json"
	
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router

	ShoppingItems []Item
}

func NewServer() *Server{
	s := &Server{
		Router: mux.NewRouter(),
		ShoppingItems: []Item{},
	}
	s.Routes()
	return s
}

func (s *Server) Routes(){
	s.HandleFunc("/shopping-items", (s.listShoppingItems())).Methods("GET")
	s.HandleFunc("/shopping-items", s.createShoppingItems()).Methods("POST")
	s.HandleFunc("/shopping-items/{id}", s.removeShoppingItem()).Methods("DELETE")
}

func (s *Server) createShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		
		item.ID = uuid.New()
		s.ShoppingItems = append(s.ShoppingItems, item)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingItems() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("content-type", "application/json")
		if err := json.NewEncoder(w).Encode(s.ShoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		id := vars["id"]

		for i, item := range s.ShoppingItems {
			if item.ID.String() == id {
				s.ShoppingItems = append(s.ShoppingItems[:i], s.ShoppingItems[i+1:]...)
				break
			}
		}
		// w.WriteHeader(http.StatusNoContent)
	}
}
