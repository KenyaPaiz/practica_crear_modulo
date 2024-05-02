package practicacrearmodulo

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("El campo nombre esta vacio")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// recibiendo varios saludos
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// va devolver saludos random
func randomFormat() string {
	formats := []string{
		"Hola, %v! Bienvenido",
		"Que bueno verte! %v",
		"Saludo mi querido %v",
	}

	return formats[rand.Intn(len(formats))]

}
