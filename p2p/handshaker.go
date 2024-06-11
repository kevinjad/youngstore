package p2p

type HandShaker func(peer Peer) error

var NOPHandShake = func(peer Peer) error {
	return nil
}
