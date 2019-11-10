package interfaces

type EncoderFactory interface {
	GetNewEncoder() Encoder
}
