package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"raft-cluster/node"
	"raft-cluster/raftcluster"

	"github.com/spf13/cobra"
)

var path string

func main() {
	command := &cobra.Command{
		Use:   "raft-cluster",
		Short: "A simple Raft cluster",
		Run: func(c *cobra.Command, args []string) {

			config := raftcluster.MustLoadConfig(path)

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
		},
	}

	command.PersistentFlags().StringVarP(&path, "config", "c", "", "Config file path")

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
