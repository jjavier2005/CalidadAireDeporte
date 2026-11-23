package dominio

type Contaminante string

const (
	PM10 Contaminante = "PM10"
	PM25 Contaminante = "PM25"
	SO2  Contaminante = "SO2"
	NO2  Contaminante = "NO2"
)

type DatoContaminante struct {
	Valor      float64
	Disponible bool
}
