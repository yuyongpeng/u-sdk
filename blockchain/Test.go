package blockchain

import (
	"fmt"
)

type KeyManager interface {
	Add(a int)
}

type Key struct {
}

func (k Key) Add(a int) {
	fmt.Println("key-a")
}

type Logger interface {
	Bdd(b int)
}

type BaseClient interface {
	Logger
	KeyManager
}

type Client struct {
	BaseClient
	a int
	k Key
}

func (c Client) test() {
	c.Add(2)
}

func NewClient() Client {
	return Client{a: 4}
}

type integration struct {
	Client
}

func (i *integration) Abc() {
	i.Client = NewClient()
	i.Add(4)
}
