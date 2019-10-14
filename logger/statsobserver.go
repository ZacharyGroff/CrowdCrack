package logger

import (
	"fmt"
	"time"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type StatsObserver struct {
	logger Logger
	tracker Tracker
	config *models.ServerConfig
	startTime time.Time
	stop chan bool
}

func (s *StatsObserver) Start() {
	ticker := time.NewTicker(time.Duration(s.config.LogFrequencyInSeconds) * time.Second)
	for {
		select {
		case <- ticker.C:
			s.logStats()
		case <- s.stop:
			ticker.Stop()
			s.logger.LogMessage("Stopping stats observer...")
			return
		}
	}
}

func (s *StatsObserver) logStats() {
	s.logPasswordStats()
	s.logHashesComputedStats()
	s.logHashesCrackedStats()
}

func (s *StatsObserver) logPasswordStats() {
	s.logPasswordsSentTotal()
	s.logPasswordsSentPerMinute()
}

func (s *StatsObserver) logHashesComputedStats() {
	s.logHashesComputedTotal()
	s.logHashesComputedPerMinute()
}

func (s *StatsObserver) logHashesCrackedStats() {
	s.logHashesCrackedTotal()
	s.logHashesCrackedPerMinute()
}

func (s *StatsObserver) logPasswordsSentTotal() {
	passwordsSent := s.tracker.GetPasswordsSent()
	logMessage := fmt.Sprintf("%d passwords sent to clients in total.\n", passwordsSent)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesComputedTotal() {
	hashesComputed := s.tracker.GetHashesComputed()
	logMessage := fmt.Sprintf("%d hashes computed in total.\n", hashesComputed)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesCrackedTotal() {
	hashesCracked := s.tracker.GetHashesCracked()
	logMessage := fmt.Sprintf("%d hashes cracked in total.\n", hashesCracked)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logPasswordsSentPerMinute() {
	passwordsSent := s.tracker.GetPasswordsSent()
	passwordsSentPerMinute := s.getActionsPerMinute(passwordsSent)
	logMessage := fmt.Sprintf("%f passwords sent per minute.\n", passwordsSentPerMinute)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesComputedPerMinute() {
	hashesComputed := s.tracker.GetHashesComputed()
	hashesComputedPerMinute := s.getActionsPerMinute(hashesComputed)
	logMessage := fmt.Sprintf("%f hashes computed per minute.\n", hashesComputedPerMinute)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) logHashesCrackedPerMinute() {
	hashesCracked := s.tracker.GetHashesCracked()
	hashesCrackedPerMinute := s.getActionsPerMinute(hashesCracked)
	logMessage := fmt.Sprintf("%f hashes cracked per minute.\n", hashesCrackedPerMinute)
	s.logger.LogMessage(logMessage)
}

func (s *StatsObserver) getActionsPerMinute(numActions uint64) float64 {
	duration := time.Now().Sub(s.startTime)
	minutes := duration.Round(time.Minute).Minutes()
	passwordsSentPerMinute := float64(numActions) / minutes

	return passwordsSentPerMinute
}