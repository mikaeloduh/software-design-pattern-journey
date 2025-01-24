package main

import "time"

// mockTimeProvider provides a mechanism to mock time.Sleep in tests by using a channel.
type mockTimeProvider struct {
	c               chan struct{}
	accumulatedTime time.Duration
}

// newMockTimeProvider initializes and returns a new instance of mockTimeProvider.
func newMockTimeProvider() *mockTimeProvider {
	return &mockTimeProvider{
		c: make(chan struct{}, 1), // Buffered to prevent blocking on send
	}
}

// Sleep blocks until it receives a signal on the channel, simulating a sleep operation.
func (m *mockTimeProvider) Sleep(d time.Duration) {
	targetTime := m.accumulatedTime + d
	for m.accumulatedTime < targetTime {
		<-m.c
	}
}

func (m *mockTimeProvider) Advance(d time.Duration) {
	m.accumulatedTime += d
	// Signal any waiting Sleep call
	m.c <- struct{}{}
}
