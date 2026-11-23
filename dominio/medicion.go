package dominio

import "time"

type Medicion struct {
	FechaHora time.Time
	Valores   map[Contaminante]float64
}

type Historial struct {
	Mediciones []Medicion
}
