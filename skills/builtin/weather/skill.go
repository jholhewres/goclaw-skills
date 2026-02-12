// Package weather implementa a skill de clima para o AgentGo Copilot.
// Fornece informações de clima atual e previsão via OpenWeatherMap API.
package weather

import (
	"context"
	"fmt"
)

// WeatherSkill fornece informações de clima e previsão.
type WeatherSkill struct {
	apiKey      string
	defaultCity string
}

// New cria uma nova instância da skill de clima.
func New(config map[string]any) (*WeatherSkill, error) {
	apiKey, _ := config["api_key"].(string)
	defaultCity, _ := config["default_city"].(string)
	if defaultCity == "" {
		defaultCity = "São Paulo"
	}

	return &WeatherSkill{
		apiKey:      apiKey,
		defaultCity: defaultCity,
	}, nil
}

// Init inicializa a skill com a configuração fornecida.
func (s *WeatherSkill) Init(_ context.Context, config map[string]any) error {
	if key, ok := config["api_key"].(string); ok {
		s.apiKey = key
	}
	return nil
}

// Execute executa a skill com o input fornecido.
func (s *WeatherSkill) Execute(_ context.Context, input string) (string, error) {
	// TODO: Implementar chamada real à API OpenWeatherMap.
	return fmt.Sprintf("Weather skill executada com input: %s", input), nil
}

// Shutdown libera recursos da skill.
func (s *WeatherSkill) Shutdown() error {
	return nil
}
