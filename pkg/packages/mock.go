package packages

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// Mock LangManager interface
type Mock struct {
	mock.Mock
}

// Install mock
func (m *Mock) Install(_ context.Context, dir string, requirements LanguageRequirements, commands []string) error {
	args := m.Called(dir, requirements, commands)
	return args.Error(0)
}

// FindExec mock
func (m *Mock) FindExec(_ context.Context, requirements LanguageRequirements, cmdExec string) ([]string, error) {
	args := m.Called(requirements, cmdExec)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}
