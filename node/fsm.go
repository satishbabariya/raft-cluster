package node

import (
	"io"

	"github.com/hashicorp/raft"
)

type FSM struct{}

func (f *FSM) Apply(log *raft.Log) interface{} {
	println("Log applied:", string(log.Data))
	return nil
}

func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return &Snapshot{}, nil
}

func (f *FSM) Restore(io.ReadCloser) error {
	return nil
}
