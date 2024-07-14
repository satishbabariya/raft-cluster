package node

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb/v2"
)

func NewRaftNode(nodeID, bindAddr, dataDir string, peerAddrs []string) (*raft.Raft, error) {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(nodeID)

	addr, err := net.ResolveTCPAddr("tcp", bindAddr)
	if err != nil {
		return nil, err
	}

	transport, err := raft.NewTCPTransport(bindAddr, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}

	snapshots, err := raft.NewFileSnapshotStore(dataDir, 1, os.Stderr)
	if err != nil {
		return nil, err
	}

	logStore, err := raftboltdb.NewBoltStore(filepath.Join(dataDir, "raft-log.bolt"))
	if err != nil {
		return nil, err
	}

	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(dataDir, "raft-stable.bolt"))
	if err != nil {
		return nil, err
	}

	fsm := &FSM{}
	raftNode, err := raft.NewRaft(config, fsm, logStore, stableStore, snapshots, transport)
	if err != nil {
		return nil, err
	}

	servers := make([]raft.Server, len(peerAddrs))
	for i, peer := range peerAddrs {
		servers[i] = raft.Server{
			ID:      raft.ServerID(fmt.Sprintf("node%d", i+1)),
			Address: raft.ServerAddress(peer),
		}
	}

	raftConfig := raft.Configuration{Servers: servers}
	raftNode.BootstrapCluster(raftConfig)

	return raftNode, nil
}
