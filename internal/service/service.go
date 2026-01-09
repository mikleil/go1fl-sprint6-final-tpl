package service

import (
	"os"
	"time"

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

		os.WriteFile(time.Now().String()+".txt", []byte(text), 0644)

		return text, nil
	} else {
		morse := morse.ToMorse(input)

		os.WriteFile(time.Now().String()+".txt", []byte(morse), 0644)

		return morse, nil
	}
}
