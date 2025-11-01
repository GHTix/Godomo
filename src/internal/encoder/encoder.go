package encoder

type Encoder interface {
	Write(data interface{}) []byte
}
