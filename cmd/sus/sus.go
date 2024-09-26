package main

import (
	"encoding/json"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/somsubhra/susproto/pkg/sus"
	logger "log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/raft"
)

func main() {
	filename := os.Getenv("INPUT")
	if len(filename) == 0 {
		filename = filepath.Join("data", "names.txt")
	}

	names, err := sus.ReadNamesFromFile(filename)
	if err != nil {
		logger.Fatalf("failed to read names from file: %v", err)
	}

	config := raft.DefaultConfig()

	address := os.Getenv("ADDRESS")
	if len(address) == 0 {
		address = "127.0.0.1:8000"
	}

	config.LocalID = raft.ServerID(address)

	state := &sus.NameState{
		Names: names, // Names read from the file
		Sus:   []string{},
	}

	store := os.Getenv("STORE")
	if len(store) == 0 {
		store = "raft-log.bolt"
	}

	logStore, err := raftboltdb.NewBoltStore(store)
	if err != nil {
		logger.Fatalf("failed to create log store: %v", err)
	}

	snapshotStore := raft.NewInmemSnapshotStore()

	var bootstrapServers []raft.Server
	bootstrapServersString := os.Getenv("BOOTSTRAP_SERVERS")
	if len(bootstrapServersString) != 0 {
		bootstrapServerStringList := strings.Split(bootstrapServersString, ",")
		for _, bootstrapServerString := range bootstrapServerStringList {
			bootstrapServerString = strings.TrimSpace(bootstrapServerString)
			bootstrapServers = append(bootstrapServers, raft.Server{
				ID:      raft.ServerID(bootstrapServerString),
				Address: raft.ServerAddress(bootstrapServerString),
			})
		}
	}

	bootstrapServers = append(bootstrapServers, raft.Server{
		ID:      config.LocalID,
		Address: raft.ServerAddress(address),
	})

	transport, err := raft.NewTCPTransport(address, nil, 2, 10*time.Second, os.Stdout)
	if err != nil {
		logger.Fatalf("failed to create transport: %v", err)
	}

	raftNode, err := raft.NewRaft(config, state, logStore, logStore, snapshotStore, transport)
	if err != nil {
		logger.Fatalf("failed to create raft: %v", err)
	}

	raftNode.BootstrapCluster(raft.Configuration{
		Servers: bootstrapServers,
	})

	for {
		if raftNode.State() == raft.Leader {
			logger.Println("This node is the leader.")

			rand.Seed(time.Now().UnixNano())

			minimum := 1
			maximum := len(names)

			numToSelect := rand.Intn(maximum-minimum) + minimum
			susNames := state.SelectRandomSus(numToSelect)

			data, err := json.Marshal(susNames)
			if err != nil {
				logger.Fatalf("failed to marshal command: %v", err)
			}

			f := raftNode.Apply(data, 10*time.Second)
			if err := f.Error(); err != nil {
				logger.Fatalf("failed to apply log: %v", err)
			}

			logger.Printf("I am the leader and here is the sus list I declare :P - %v", susNames)
			break
		}

		logger.Printf("I am just the follower, waiting for the leader to come and declare the sus list :P")
		time.Sleep(200 * time.Millisecond)
	}
}
