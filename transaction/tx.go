package transaction

import (
	"fmt"

	"github.com/discoviking/fsm"

	"github.com/ghettovoice/gosip/log"
	"github.com/ghettovoice/gosip/sip"
	"github.com/ghettovoice/gosip/transport"
)

type TxKey string

func (key TxKey) String() string {
	return string(key)
}

// Tx is an common SIP transaction
type Tx interface {
	log.Loggable

	Init() error
	Key() TxKey
	Origin() sip.Request
	// Receive receives message from transport layer.
	Receive(msg sip.Message) error
	String() string
	Transport() transport.Layer
	Terminate()
	Errors() <-chan error
	Done() <-chan bool
}

type commonTx struct {
	key      TxKey
	fsm      *fsm.FSM
	origin   sip.Request
	tpl      transport.Layer
	lastResp sip.Response

	errs    chan error
	lastErr error
	done    chan bool

	log log.Logger
}

func (tx *commonTx) String() string {
	if tx == nil {
		return "<nil>"
	}

	fields := tx.Log().Fields().WithFields(log.Fields{
		"key": tx.key,
	})

	return fmt.Sprintf("%s<%s>", tx.Log().Prefix(), fields)
}

func (tx *commonTx) Log() log.Logger {
	return tx.log
}

func (tx *commonTx) Origin() sip.Request {
	return tx.origin
}

func (tx *commonTx) Key() TxKey {
	return tx.key
}

func (tx *commonTx) Transport() transport.Layer {
	return tx.tpl
}

func (tx *commonTx) Errors() <-chan error {
	return tx.errs
}

func (tx *commonTx) Done() <-chan bool {
	return tx.done
}
