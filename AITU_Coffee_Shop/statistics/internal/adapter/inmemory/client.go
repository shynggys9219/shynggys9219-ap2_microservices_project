package inmemory

import (
	"sync"

	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
)

type Client struct {
	clients map[uint64]model.Client
	m       sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		clients: make(map[uint64]model.Client),
		m:       sync.RWMutex{},
	}
}

func (q *Client) Set(client model.Client) {
	q.m.Lock()
	defer q.m.Unlock()

	q.clients[client.ID] = client
}

func (q *Client) SetMany(clients []model.Client) {
	q.m.Lock()
	defer q.m.Unlock()

	for _, client := range clients {
		q.clients[client.ID] = client
	}
}

func (q *Client) Get(clientID uint64) (model.Client, bool) {
	q.m.RLock()
	defer q.m.RUnlock()

	client, ok := q.clients[clientID]

	return client, ok
}

func (q *Client) Delete(clientID uint64) {
	q.m.Lock()
	defer q.m.Unlock()

	delete(q.clients, clientID)
}
