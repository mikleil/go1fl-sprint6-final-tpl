package service

import (
	"go1fl-sprint6-final-tpl/internal/handlers"
	"go1fl-sprint6-final-tpl/pkg/morse"
)

type Service struct {
	converter morse.Converter
}

func NewService(converter morse.Converter) *Service {
	return &Service{
		converter: converter,
	}
}

func (s *Service) ConvertString(input string) (string, error) {
	var mrs bool

	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' {
			mrs = false
			break
		}
		mrs = true
	}

	if mrs {
		text := morse.ToText(input)
		handlers.ReadToFile(text)
		return text, nil
	} else {
		morse := morse.ToMorse(input)
		handlers.ReadToFile(morse)
		return morse, nil
	}
}
