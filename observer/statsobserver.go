package observer

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/tracker"
	"time"
)

type StatsObserver struct {
	logger    interfaces.Logger
	tracker   interfaces.Tracker
	config    *models.Config
	startTime time.Time
	stop      chan bool
}

func NewStatsObserver(l *logger.ConcurrentLogger, t *tracker.StatsTracker, p interfaces.ConfigProvider) *StatsObserver {
	c := p.GetConfig()
	start := time.Now()
	stop := make(chan bool)
	return &StatsObserver{
		logger:    l,
		tracker:   t,
		config:    c,
		startTime: start,
		stop:      stop,
	}
}

func (s *StatsObserver) Start() {
	ticker := time.NewTicker(time.Duration(s.config.LogFrequencyInSeconds) * time.Second)
	for {
		select {
		case <-ticker.C:
			s.logStats()
		case <-s.stop:
			ticker.Stop()
			return
		}
	}
}

func (s *StatsObserver) Stop() {
	s.logger.LogMessage("Stopping stats observer...")
	s.stop <- true
}

func (s *StatsObserver) logStats() {
	s.logPasswordStats()
	s.logHashesComputedStats()
	s.logVerifierStats()
}

func (s *StatsObserver) logPasswordStats() {
	s.logPasswordsSentTotal()
	s.logPasswordsSentPerMinute()
}

func (s *StatsObserver) logHashesComputedStats() {
	s.logHashesComputedTotal()
	s.logHashesComputedPerMinute()
}

func (s *StatsObserver) logVerifierStats() {
	s.logHashMatchAttemptsTotal()
	s.logHashMatchAttemptsPerMinute()
	s.logHashesCrackedTotal()
	s.logHashesCrackedPerMinute()
}

func (s *StatsObserver) logPasswordsSentTotal() {
	passwordsSent := s.tracker.GetPasswordsSent()
	logMessage := fmt.Sprintf("%d passwords sent to clients in total.", passwordsSent)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesComputedTotal() {
	hashesComputed := s.tracker.GetHashesComputed()
	logMessage := fmt.Sprintf("%d hashes computed in total.", hashesComputed)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashMatchAttemptsTotal() {
	hashMatchAttempts := s.tracker.GetHashMatchAttempts()
	logMessage := fmt.Sprintf("%d hash match attempts in total.", hashMatchAttempts)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesCrackedTotal() {
	hashesCracked := s.tracker.GetHashesCracked()
	logMessage := fmt.Sprintf("%d hashes cracked in total.", hashesCracked)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logPasswordsSentPerMinute() {
	passwordsSent := s.tracker.GetPasswordsSent()
	passwordsSentPerMinute := s.getActionsPerMinute(passwordsSent)
	logMessage := fmt.Sprintf("%f passwords sent per minute.", passwordsSentPerMinute)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesComputedPerMinute() {
	hashesComputed := s.tracker.GetHashesComputed()
	hashesComputedPerMinute := s.getActionsPerMinute(hashesComputed)
	logMessage := fmt.Sprintf("%f hashes computed per minute.", hashesComputedPerMinute)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashMatchAttemptsPerMinute() {
	hashMatchAttempts := s.tracker.GetHashMatchAttempts()
	hashMatchAttemptsPerMinute := s.getActionsPerMinute(hashMatchAttempts)
	logMessage := fmt.Sprintf("%f hash match attempts per minute.", hashMatchAttemptsPerMinute)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesCrackedPerMinute() {
	hashesCracked := s.tracker.GetHashesCracked()
	hashesCrackedPerMinute := s.getActionsPerMinute(hashesCracked)
	logMessage := fmt.Sprintf("%f hashes cracked per minute.", hashesCrackedPerMinute)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) getActionsPerMinute(numActions uint64) float64 {
	duration := time.Now().Sub(s.startTime)
	minutes := duration.Round(time.Minute).Minutes()
	passwordsSentPerMinute := float64(numActions) / minutes

	return passwordsSentPerMinute
}
