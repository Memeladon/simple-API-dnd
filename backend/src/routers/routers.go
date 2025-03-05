package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// RouterFactory - фабрика для создания роутеров
type RouterFactory struct {
	corsOptions cors.Options
}

// RouterOption - функции-опции для настройки фабрики.
type RouterOption func(*RouterFactory)

// NewRouterFactory - создает новую фабрику с настраиваемыми опциями. Реализация функционального паттерна опций
func NewRouterFactory(options ...RouterOption) *RouterFactory {
	factory := &RouterFactory{
		corsOptions: cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	}

	// Применение кастомных настроек
	for _, opt := range options {
		opt(factory)
	}

	return factory
}

// TODO: Реализовать WithAllowedHeaders и WithExposedHeaders, WithMaxAge, да и в целом для других параметров..

// WithAllowedOrigins - настраивает разрешенные источники
func WithAllowedOrigins(origins []string) RouterOption {
	return func(f *RouterFactory) {
		f.corsOptions.AllowedOrigins = origins
	}
}

// WithAllowedMethods - настраивает разрешенные методы
func WithAllowedMethods(methods []string) RouterOption {
	return func(f *RouterFactory) {
		f.corsOptions.AllowedMethods = methods
	}
}

// CreateRouter - создает новый роутер с текущими настройками
func (f *RouterFactory) CreateRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(f.corsOptions))
	return router
}
