package utils

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// CustomDuration - пользовательский тип для работы с duration в JSON
type CustomDuration struct {
	time.Duration
}

// UnmarshalJSON реализует интерфейс json.Marshaler
func (d *CustomDuration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return fmt.Errorf("неверный формат длительности: %q", value)
		}
		return nil
	default:
		return errors.New("неподдерживаемый тип для длительности")
	}
}

// MarshalJSON реализует интерфейс json.Marshaler
func (d CustomDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// String возвращает строковое представление длительности
func (d CustomDuration) String() string {
	return d.Duration.String()
}

// Configuration представляет конфигурацию приложения
type Configuration struct {
	Env         string         `json:"env"`
	StoragePath string         `json:"storage_path"`
	TokenTTL    CustomDuration `json:"token_ttl"`
	GRPC        GRPCConfig     `json:"grpc"`
}

// GRPCConfig представляет конфигурацию gRPC
type GRPCConfig struct {
	Port    int            `json:"port"`
	Timeout CustomDuration `json:"timeout"`
}

func MustLoad() *Configuration {
	path := fetchConfigPath()
	if path == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist")
	}

	var cfg Configuration

	// Чтение файла конфигурации
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("failed to read config file: " + err.Error())
	}

	// Парсинг JSON
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic("failed to parse config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path из командной строки или env.
// Priority: flag > env > default.
// Default value is empty string
func fetchConfigPath() string {
	res := ""

	// go run cmd/main.go --config="./config/local.json"
	// CONFIG_PATH="./config/local.json" go run cmd/main.go
	flag.StringVar(&res, "config", "./local.json", "config file path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
