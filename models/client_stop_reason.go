package models

type ClientStopReason struct {
	Requester string
	Encoder   string
	Submitter string
}

func NewClientStopReason(r string, e string, s string) ClientStopReason {
	return ClientStopReason{
		Requester: r,
		Encoder:   e,
		Submitter: s,
	}
}
