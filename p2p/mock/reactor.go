package mock

import (
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/p2p/conn"
)

type Reactor struct {
	p2p.BaseReactor

	Channels []*conn.ChannelDescriptor
}

func NewReactor() *Reactor {
	r := &Reactor{}
	r.BaseReactor = *p2p.NewBaseReactor("Mock-PEX", r)
	r.SetLogger(log.TestingLogger())
	return r
}

<<<<<<< HEAD
func (r *Reactor) GetChannels() []*conn.ChannelDescriptor            { return r.Channels }
func (r *Reactor) AddPeer(peer p2p.Peer)                             {}
func (r *Reactor) RemovePeer(peer p2p.Peer, reason interface{})      {}
func (r *Reactor) ReceiveEnvelope(e p2p.Envelope)                    {}
func (r *Reactor) Receive(chID byte, peer p2p.Peer, msgBytes []byte) {}
=======
func (r *Reactor) GetChannels() []*conn.ChannelDescriptor { return r.Channels }
func (r *Reactor) AddPeer(_ p2p.Peer)                     {}
func (r *Reactor) RemovePeer(_ p2p.Peer, _ interface{})   {}
func (r *Reactor) Receive(_ p2p.Envelope)                 {}
>>>>>>> 111d252d7 (Fix lints (#625))
