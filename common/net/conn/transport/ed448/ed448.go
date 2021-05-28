package ed448

import (
	"github.com/pastelnetwork/gonode/common/net/conn/transport"
	"io"
	"math/rand"
	"net"
)

type Ed448 struct {
	cryptos                map[string]transport.Crypto
	isHandshakeEstablished bool
	chosenEncryptionScheme string
}

func New(cryptos ...transport.Crypto) transport.Transport {
	cryptosMap := make(map[string]transport.Crypto)
	for _, crypto := range cryptos {
		cryptosMap[crypto.About()] = crypto
	}
	return &Ed448{
		cryptos: cryptosMap,
	}
}

// IsHandshakeEstablished - returns flag that says is handshake process finished
func (transport *Ed448) IsHandshakeEstablished() bool {
	return transport.isHandshakeEstablished
}

func (transport *Ed448) readRecord(conn net.Conn) (interface{}, error) {
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small buffer

	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		buf = append(buf, tmp[:n]...)

	}
	if _, err := conn.Read(buf); err == nil {
		return nil, err
	}

	// trying to decrypt message
	switch buf[0] {
	case typeClientHello:
		return DecodeClientMsg(buf)
	case typeServerHello:
		return DecodeServerMsg(buf)
	case typeClientHandshakeMsg:
		return DecodeClientHandshakeMessage(buf)
	case typeServerHandshakeMsg:
		return DecodeServerHandshakeMsg(buf)
	default:
		return nil, ErrWrongFormat
	}
}

func (transport *Ed448) writeRecord(msg message, conn net.Conn) error {
	data, err := msg.marshall()
	if err != nil {
		return err
	}

	if _, err := conn.Write(data); err != nil {
		return err
	}

	return nil
}

func (transport *Ed448) initEncryptedConnection(conn net.Conn, cryptoAlias string, params string) (net.Conn, error) {
	crypto := transport.cryptos[cryptoAlias]
	if crypto == nil {
		return nil, ErrUnsupportedEncryption
	}

	if err := crypto.Configure(params); err != nil {
		return nil, err
	}

	return NewConn(conn, crypto), nil
}

// ToDo: update with appropriate implementation
func (transport *Ed448) getPastelID() []byte {
	pastelID := make([]byte, 4)
	rand.Read(pastelID)
	return pastelID
}

// ToDo: update with appropriate implementation
func (transport *Ed448) getSignedPastelId() *signedPastelID {
	var pastelID = transport.getPastelID()

	signPastelID := make([]byte, 4)
	rand.Read(pastelID)

	pubKey := make([]byte, 4)
	rand.Read(pubKey)

	ctx := make([]byte, 4)
	rand.Read(ctx)

	return &signedPastelID{
		pastelID:       pastelID,
		signedPastelID: signPastelID,
		pubKey:         pubKey,
		ctx:            ctx,
	}
}

// ToDo: update with appropriate implementation
func (transport *Ed448) verifyPastelID(_ []byte) bool {
	return true
}

// ToDo: update with external check through cNode
func (transport *Ed448) verifySignature(_, _, _, _ []byte) bool {
	return true
}
