package sus

import (
	"encoding/json"
	"github.com/hashicorp/raft"
	"io"
	logger "log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type NameState struct {
	Names []string
	Sus   []string
	mu    sync.Mutex
}

func (s *NameState) Snapshot() (raft.FSMSnapshot, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	names := make([]string, len(s.Names))
	copy(names, s.Names)

	sus := make([]string, len(s.Sus))
	copy(sus, s.Sus)

	return &NameSnapshot{Names: names, Sus: sus}, nil
}

func (s *NameState) Restore(snapshot io.ReadCloser) error {
	var restoredSnapshot NameSnapshot
	if err := json.NewDecoder(snapshot).Decode(&restoredSnapshot); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Names = restoredSnapshot.Names
	s.Sus = restoredSnapshot.Sus

	return nil
}

func (s *NameState) Apply(log *raft.Log) interface{} {
	var susList []string
	err := json.Unmarshal(log.Data, &susList)
	if err != nil {
		logger.Fatalf("failed to unmarshal log data: %v", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i, name := range susList {
		label := "sus" + strconv.Itoa(i)
		s.Sus = append(s.Sus, label+": "+name)
	}

	logger.Printf("Updated state with sus names: %v", s.Sus)
	return nil
}

func (s *NameState) SelectRandomSus(num int) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	selected := make(map[int]bool)
	var susList []string
	rand.Seed(time.Now().UnixNano())

	for len(susList) < num && len(susList) < len(s.Names) {
		index := rand.Intn(len(s.Names))
		if !selected[index] {
			selected[index] = true
			susList = append(susList, s.Names[index])
		}
	}
	return susList
}
