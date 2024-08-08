package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacío")
	}
	// Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos retorna un map con todos los nombres
// con un greeting msg
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

// randomFormat devuelve uno de varios formatos de mensajes de sa
// se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		" Hola %v!! Bienvenido/a!",
		" Qué bueno verte, %v!!",
		" Saludo %v, encantada de conocerte!",
	}
	return formats[rand.Intn(len((formats)))]
}
