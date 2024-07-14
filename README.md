# Raft Cluster

This project demonstrates a Raft cluster with three nodes using the `github.com/hashicorp/raft` and `github.com/hashicorp/raft-boltdb/v2` packages in Go.

## Project Structure

```plaintext
raft-cluster/
├── config/
│   └── config.yaml
├── main.go
├── node/
│   ├── fsm.go
│   ├── node.go
│   └── snapshot.go
├── raftcluster/
│   ├── config.go
│   └── node.go
└── go.mod
```

## Running the Nodes

Run each node with the appropriate configuration file:

```
go run main.go --config=config/config-node1.yaml
go run main.go --config=config/config-node2.yaml
```

## Raft Cluster Diagram

```mermaid
graph TD;
    A[Node 1] -->|Replicates logs| B[Node 2];
    B -->|Replicates logs| C[Node 3];
    C -->|Replicates logs| A;

    subgraph Node1
        direction TB
        A[Node 1]
        AID[ID: node1]
        ABindAddr[BindAddr: 127.0.0.1:5001]
        ADataDir[DataDir: /tmp/node1]
    end

    subgraph Node2
        direction TB
        B[Node 2]
        BID[ID: node2]
        BBindAddr[BindAddr: 127.0.0.1:5002]
        BDataDir[DataDir: /tmp/node2]
    end

    subgraph Node3
        direction TB
        C[Node 3]
        CID[ID: node3]
        CBindAddr[BindAddr: 127.0.0.1:5003]
        CDataDir[DataDir: /tmp/node3]
    end
```
