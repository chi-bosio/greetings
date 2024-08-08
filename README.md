# Saludos en GO

Este paquete proporciona una forma simple de obtener saludos personalizados en Go

## Instalación
Ejecuta el siguiente comando para instalas el paquete:
``` bash
go get -u github.com/chi-bosio/greetings
```

## Uso
Aquí tienes un ejemplo de como utilizar el paquete en tu código:
``` go

package main

import (
	"fmt"
	"github.com/chi-bosio/greetings"
)

func main() {
    message, err := greetings.Hellos("Chiara")

	if err != nil {
		fmt.Println("Ocurrió un error:", err)
        return
	}

	fmt.Println(message)
}

```

Este ejemplo importa el paquete github.com/chi-bosio/greetings y llama a la función Hello, obteniendo un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.