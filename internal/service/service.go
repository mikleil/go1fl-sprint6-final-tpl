package service

import (
	"fmt"
	"log"
	"os"

	//	"path/filepath"
	"strings"
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

	var result string
	if mrs {
		result = s.converter.ToText(input)
	} else {
		result = s.converter.ToMorse(input)
	}

	timeStr := time.Now().UTC().String()

	timeStr = strings.ReplaceAll(timeStr, ":", "_")
	timeStr = strings.ReplaceAll(timeStr, " ", "_")

	fileName := timeStr + ".txt"

	/*
		ext := filepath.Ext(fileName)
		fmt.Printf("Создаем файл с расширением: %s\n", ext)
	*/

	err := os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		log.Printf("Failed to write file %s: %v", fileName, err)
		return "", fmt.Errorf("Failed to write file: %v", err)
	}

	return result, nil
}
