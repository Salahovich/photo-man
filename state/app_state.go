package state

import (
	"image"
	"sync"
)

type AppState struct {
	currentImage  image.Image
	originalImage image.Image
	mutex         sync.RWMutex
	format        string
	listeners     []func(image.Image)
}

func NewAppState() *AppState {
	return &AppState{
		listeners: make([]func(image.Image), 0),
	}
}

func (s *AppState) SetImage(img image.Image) {
	s.mutex.Lock()
	s.currentImage = img
	if s.originalImage == nil {
		s.originalImage = img
	}
	s.mutex.Unlock()

	s.notify()
}

func (s *AppState) SetOriginalImage(img image.Image) {
	s.originalImage = img
}

func (s *AppState) SetFormat(format string) {
	s.format = format
}

func (s *AppState) GetImage() image.Image {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.currentImage
}

func (s *AppState) GetOriginalImage() image.Image {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.originalImage
}

func (s *AppState) GetFormat() string {
	return s.format
}

func (s *AppState) RegisterListener(callback func(image.Image)) {
	s.listeners = append(s.listeners, callback)
}

func (s *AppState) notify() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, callback := range s.listeners {
		// Call the function with the new image
		callback(s.currentImage)
	}
}
