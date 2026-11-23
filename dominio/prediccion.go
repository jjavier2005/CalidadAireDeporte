package dominio

import "time"

type Prediccion struct {
	FechaHora  time.Time
	Valores    map[Contaminante]float64
	disponible bool
}
