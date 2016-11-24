package swclient

// hasher will probably removed soon (digest.go line 164)
type hasher interface {
	Reset()
	Write([]byte) (int, error)
	Sum([]byte) []byte
}
