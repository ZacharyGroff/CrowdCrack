package encoder

type EncoderFactory interface {
	GetNewEncoder() Encoder
}
