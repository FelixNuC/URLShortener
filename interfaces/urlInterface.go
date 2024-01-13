package interfaces

type URLInterface interface {
	GenerateShortURL() string
	Validate() error
	Save() error
	Delete() error
	Get() (string, error)
}
