package main

type Config struct {
	Port int `json:"name" env:"PORT" envDefault:"3000"`
}
