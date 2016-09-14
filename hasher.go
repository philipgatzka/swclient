package swclient

type hasher interface {
	Reset()
	Write([]byte) (int, error)
	Sum([]byte) []byte
}
