package main

import (
    "log"
    "time"

    "raft-cluster/raftcluster"
    "raft-cluster/node"
)

func main() {
    config := raftcluster.MustLoadConfig()

    raftNode, err := node.NewRaftNode(config.NodeID, config.BindAddr, config.DataDir, config.Peers)
    if err != nil {
        log.Fatalf("Failed to create Raft node: %v", err)
    }

    for {
        // Simulate log entries
        time.Sleep(5 * time.Second)
        logEntry := []byte("Event from " + config.NodeID + " at " + time.Now().String())
        raftNode.Apply(logEntry, 10*time.Second)
    }
}
