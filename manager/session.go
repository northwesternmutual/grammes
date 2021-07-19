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
	id *uuid.UUID
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

func (s *sessionManager) NewNoopSession() Session {
	return &session{
		e: s.e,
	}
}

func (s *sessionManager) NewSession() Session {
	id := uuid.New()
	return &session{
		e: s.e,
		id: &id,
	}
}

func (s *sessionManager) GetSession(sessionId uuid.UUID) Session {
	return &session{
		e: s.e,
		id: &sessionId,
	}
}

func (s *sessionManager) WithNewSession(f func(Session) error) error {
	return s.WithSession(s.NewSession(), f)
}

func (s *sessionManager) WithSession(ss Session, f func(Session) error) error {
	var err error
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
	return s.e(stringQuery, nil, map[string]string{}, map[string]string{}, s.id)
}

func (s *session) ExecuteQuery(queryObj query.Query) (res [][]byte, err error) {
	return s.ExecuteStringQuery(queryObj.String())
}

func (s *session) Close() error {
	if s.id == nil {
		return nil
	}

	_, err := s.e("", nil, map[string]string{}, map[string]string{}, s.id)
	return err
}

func (s *session) Commit() error {
	if s.id == nil {
		return nil
	}

	commit := traversal.NewTraversal()
	commit.AddStep("tx")
	commit.AddStep("commit")

	_, err := s.ExecuteQuery(commit)
	return err
}

func (s *session) Rollback() error {
	if s.id == nil {
		return nil
	}

	rollback := traversal.NewTraversal()
	rollback.AddStep("tx")
	rollback.AddStep("rollback")

	_, err := s.ExecuteQuery(rollback)
	return err
}
