package manager

import (
	"github.com/google/uuid"
	"github.com/northwesternmutual/grammes/logging"
	"github.com/northwesternmutual/grammes/query"
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