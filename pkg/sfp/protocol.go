package sfp

import "time"

type Protocol interface {
	Launch()
	Decode()
	Encode()
}

type ProtocolOptions = map[string]interface{}

type SFProtocol struct {
	Timeout   time.Duration
	AllowConn int
}

func NewSFP() *SFProtocol {
	return &SFProtocol{}
}

func (s *SFProtocol) Launch() {

}
