package shortener

type RedirectSerialzer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(redirect *Redirect) ([]byte, error)
}
