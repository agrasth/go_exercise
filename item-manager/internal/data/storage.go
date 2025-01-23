package data

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "sync"

    "github.com/google/uuid"
)

type Item struct {
    ID   uuid.UUID `json:"id"`
    Name string    `json:"name"`
}

type Storage struct {
    items []Item
    mu    sync.Mutex
    filename string
}

func NewStorage(filename string) *Storage {
    s := &Storage{
        filename: filename,
    }
    
    s.loadItems()
    return s
}

func (s *Storage) AddItem(item Item) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.items = append(s.items, item)
    s.saveItems()
}

func (s *Storage) GetItems() []Item {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.items
}

func (s *Storage) RemoveItem(id string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for i, item := range s.items {
        if item.ID.String() == id {
            s.items = append(s.items[:i], s.items[i+1:]...)
            break
        }
    }
    s.saveItems()
}

func (s *Storage) loadItems() {
    file, err := os.Open(s.filename)
    if err != nil {
        if os.IsNotExist(err) {
            return
        }
        panic(err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&s.items)
    if err != nil {
        panic(err)
    }
}

func (s *Storage) saveItems() {
    data, err := json.MarshalIndent(s.items, "", "  ")
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(s.filename, data, 0644)
    if err != nil {
        panic(err)
    }
}

func NewUUID() uuid.UUID {
    return uuid.New()
}