package dominio

import "time"

type Medicion struct {
	FechaHora time.Time
	Valores   map[Contaminante]DatoContaminante
}
type Historial struct {
	Mediciones []Medicion
}
