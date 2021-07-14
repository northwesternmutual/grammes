package manager

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/query"
	"github.com/northwesternmutual/grammes/query/traversal"
)

// sessionManager handles the sessions actions to the server.
type sessionManager struct {
	logger         logging.Logger
	e executor
}

type session struct {
	id uuid.UUID
	e executor
}

// newSessionManager returns a new Session Manager that
// implements the SessionManager interface.
func newSessionManager(logger logging.Logger, executor executor) *sessionManager {
	return &sessionManager{
		logger: logger,
		e: executor,
	}
}

func (s *sessionManager) NewSession() Session {
	return &session{
		e: s.e,
		id: uuid.New(),
	}
}

func (s *sessionManager) WithSession(f func(Session) error) error {
	var err error
	ss := s.NewSession()
	defer func() {
		err2 := ss.Close()
		if err == nil { // Capture close error if everything else was ok
			err = err2
		}
	}()

	err = f(ss)
	if err == nil {
		err = ss.Commit()
	} else {
		err2 := ss.Rollback()
		if err2 != nil { // Keep both errors
			err = fmt.Errorf("error rolling back: %v, original error: %w", err2, err)
		}
	}

	return err
}

func (s *session) ExecuteStringQuery(stringQuery string) (res [][]byte, err error) {
	return s.e(stringQuery, nil, map[string]string{}, map[string]string{}, &s.id)
}

func (s *session) ExecuteQuery(queryObj query.Query) (res [][]byte, err error) {
	return s.ExecuteStringQuery(queryObj.String())
}

func (s *session) Close() error {
	_, err := s.e("", nil, map[string]string{}, map[string]string{}, &s.id)
	return err
}

func (s *session) Commit() error {
	commit := traversal.NewTraversal()
	commit.AddStep("tx")
	commit.AddStep("commit")

	_, err := s.ExecuteQuery(commit)
	return err
}

func (s *session) Rollback() error {
	rollback := traversal.NewTraversal()
	rollback.AddStep("tx")
	rollback.AddStep("rollback")

	_, err := s.ExecuteQuery(rollback)
	return err
}
