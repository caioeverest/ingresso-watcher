package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Interval   uint64
	HttpPort   string
	ApiKey     string
	ApiAddress string
}

func InitConfig() *Config {
	interval := uint64(60)
	tmpInterval, ok := os.LookupEnv("INTERVAL")
	if ok {
		interval, _ = strconv.ParseUint(tmpInterval, 10, 64)
	}

	httpPort, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		httpPort = "7000"
	}

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Panicf("Did not found API_KEY")
	}

	apiAddress, ok := os.LookupEnv("API_ADDRESS")
	if !ok {
		log.Panicf("Did not found API_ADDRESS")
	}

	confSetted := &Config{
		Interval:   interval,
		HttpPort:   httpPort,
		ApiKey:     apiKey,
		ApiAddress: apiAddress,
	}

	printConfig(confSetted)

	return confSetted
}

func printConfig(c *Config) {
	log.Printf("interval: %d", c.Interval)
	log.Printf("api-key: %s", c.ApiKey)
}
