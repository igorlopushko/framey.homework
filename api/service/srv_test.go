package service

import (
	"testing"

	"github.com/igorlopushko/framey.homework/api/internal/mock"
	"github.com/igorlopushko/framey.homework/api/provider"
	"github.com/sirupsen/logrus"
)

func TestExec_NilProviders_ReturnsEmptyResults(t *testing.T) {
	s := &Service{}
	r, err := s.Exec()

	if err != nil {
		t.Errorf("srv.Exec(). Error is not nil: %v", err)
	}

	if len(r) > 0 {
		t.Errorf("srv.Exec(). Results are not empty. Have to be 0 length")
	}
}

func TestExec_EmptyProviders_ReturnsEmptyResults(t *testing.T) {
	p := make(map[string]IProvider)
	s := &Service{Providers: p}
	r, err := s.Exec()

	if err != nil {
		t.Errorf("srv.Exec(). Error is not nil: %v", err)
	}

	if len(r) > 0 {
		t.Errorf("srv.Exec(). Results are not empty. Have to be 0 length")
	}
}

func TestExec_ProviderReturnsError_ReturnsError(t *testing.T) {
	logrus.SetLevel(logrus.FatalLevel)
	p := make(map[string]IProvider)
	p["mock"] = &mock.Provider{ReturnError: true, Name: "mock"}
	s := &Service{Providers: p}
	r, err := s.Exec()

	if r != nil {
		t.Errorf("srv.Exec(). Results are not nil. Have to be nil")
	}

	if err == nil {
		t.Errorf("srv.Exec(). Error is nil. Has to be an error")
	}
}

func TestExec_CorrectProvider_ReturnsCorrectResults(t *testing.T) {
	p := make(map[string]IProvider)
	p["mock"] = &mock.Provider{ReturnError: false, Name: "mock"}
	s := &Service{Providers: p}
	r, err := s.Exec()

	if len(r) > 1 || len(r) == 0 {
		t.Errorf("srv.Exec(). Results contain %d elements. Results len has to be 1", len(r))
	}

	if err != nil {
		t.Errorf("srv.Exec(). Error occurred: %v", err)
	}
}

func BenchmarkSpeedtestNet(b *testing.B) {
	p := make(map[string]IProvider)
	p["ookla"] = &provider.SpeedTestProvider{Name: "ookla"}
	s := &Service{Providers: p}
	for i := 0; i < b.N; i++ {
		r, err := s.Exec()

		if len(r) > 1 || len(r) == 0 {
			b.Errorf("srv.Exec(). Results contain %d elements. Results len has to be 1", len(r))
		}

		if err != nil {
			b.Errorf("srv.Exec(). Error occurred: %v", err)
		}
	}
}

func BenchmarkFastCom(b *testing.B) {
	p := make(map[string]IProvider)
	p["netflix"] = &provider.FastProvider{Name: "netflix"}
	s := &Service{Providers: p}

	for i := 0; i < b.N; i++ {
		r, err := s.Exec()

		if len(r) > 1 || len(r) == 0 {
			b.Errorf("srv.Exec(). Results contain %d elements. Results len has to be 1", len(r))
		}

		if err != nil {
			b.Errorf("srv.Exec(). Error occurred: %v", err)
		}
	}
}
