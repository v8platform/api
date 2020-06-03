package agent

import agent "github.com/khorevaa/go-v8platform/agent/client"

type ClientID string

type Pool struct {
	pool        map[ClientID]*Client
	count       int
	maxPoolSize int
}

func (p *Pool) process() {

	p.pool[id] = &Client{
		ID: id,
	}

	p.count++

}

func (p *Pool) newClient(id ClientID) {

	p.pool[id] = &Client{
		ID: id,
	}

	p.count++

}

type Client struct {
	ID     ClientID // идентификатор строка подключения к базе
	ready  bool
	client *agent.AgentClient
}

func (c *Client) Start() error {

	c.ready = true
	// TODO
	return nil
}

func (c *Client) Stop() error {

	c.client.Stop()
	c.ready = false
	return nil
}
