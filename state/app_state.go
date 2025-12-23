package state

import (
	"image"
	"sync"
)

type AppState struct {
	currentImage  image.Image
	scaledImage   image.Image
	originalImage image.Image
	mutex         sync.RWMutex
	format        string
	listeners     []func(image.Image) image.Image
	channel       chan image.Image
}

func NewAppState() *AppState {
	return &AppState{
		listeners: make([]func(image.Image) image.Image, 0),
		channel:   make(chan image.Image),
	}
}

func (s *AppState) UpdateSceneImage(img image.Image) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.currentImage = img
	s.channel <- s.currentImage
}

func (s *AppState) SetImage(original image.Image, scaled image.Image) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.originalImage = original
	s.scaledImage = scaled
	s.currentImage = scaled
	s.channel <- s.currentImage
}

func (s *AppState) SetFormat(format string) {
	s.format = format
}

func (s *AppState) GetScaledImage() image.Image {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.scaledImage
}

func (s *AppState) GetOriginalImage() image.Image {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.originalImage
}

func (s *AppState) GetCurrentImage() image.Image {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.currentImage
}

func (s *AppState) GetFormat() string {
	return s.format
}

func (s *AppState) GetChannel() chan image.Image {
	return s.channel
}

func (s *AppState) RegisterListener(callback func(image.Image) image.Image) {
	s.listeners = append(s.listeners, callback)
}

func (s *AppState) ApplyAllModification() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, callback := range s.listeners {
		s.originalImage = callback(s.originalImage)
	}
}
