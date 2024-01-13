package interfaces

type URLInterface interface {
	GenerateShortURL(string) string
	Validate() error
}
