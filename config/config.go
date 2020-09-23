package config

import (
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/auth/jwt"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/config/reader"
	"github.com/micro/go-micro/v2/server"

	"github.com/bimoyong/go-util/config"
	"github.com/bimoyong/go-util/config/database"
)

// Config struct
type Config struct{}

// Init function
func Init() (err error) {
	if err = newAuth(); err != nil {
		return
	}

	if err = newBroker(); err != nil {
		return
	}

	config.Init(&Config{})

	return
}

// Auth_PrivateKey function
func (s *Config) Auth_PrivateKey(value reader.Value) {
	privateKey := auth.PrivateKey(value.String(""))
	auth.DefaultAuth.Init(privateKey)
}

// Auth_PublicKey function
func (s *Config) Auth_PublicKey(value reader.Value) {
	publicKey := auth.PublicKey(value.String(""))
	auth.DefaultAuth.Init(publicKey)
}

// Database function
func (s *Config) Database(value reader.Value) {
	database.Refresh(value)
}

func newAuth() error {
	auth.DefaultAuth = jwt.NewAuth()

	return nil
}

func newBroker() error {
	broker.DefaultBroker = server.DefaultServer.Options().Broker

	return broker.DefaultBroker.Connect()
}
