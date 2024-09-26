package sus

import (
	"encoding/json"
	"github.com/hashicorp/raft"
)

type NameSnapshot struct {
	Names []string
	Sus   []string
}

func (ns *NameSnapshot) Persist(sink raft.SnapshotSink) error {
	data, err := json.Marshal(ns)
	if err != nil {
		err := sink.Cancel()
		if err != nil {
			return err
		}
		return err
	}

	if _, err := sink.Write(data); err != nil {
		err := sink.Cancel()
		if err != nil {
			return err
		}
		return err
	}

	if err := sink.Close(); err != nil {
		return err
	}

	return nil
}

func (ns *NameSnapshot) Release() {}
