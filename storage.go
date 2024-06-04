package main

import (
	"errors"
	"sync"
)

type Storage struct {
	exoplanets map[string]Exoplanet
	mu         sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		exoplanets: make(map[string]Exoplanet),
	}
}

func (s *Storage) AddExoplanet(planet Exoplanet) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.exoplanets[planet.ID] = planet
}

func (s *Storage) GetExoplanets() []Exoplanet {
	s.mu.Lock()
	defer s.mu.Unlock()
	planets := make([]Exoplanet, 0, len(s.exoplanets))
	for _, planet := range s.exoplanets {
		planets = append(planets, planet)
	}
	return planets
}

func (s *Storage) GetExoplanetByID(id string) (Exoplanet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	planet, exists := s.exoplanets[id]
	if !exists {
		return Exoplanet{}, errors.New("exoplanet not found")
	}
	return planet, nil
}

func (s *Storage) UpdateExoplanet(id string, updated Exoplanet) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.exoplanets[id]
	if !exists {
		return errors.New("exoplanet not found")
	}
	s.exoplanets[id] = updated
	return nil
}

func (s *Storage) DeleteExoplanet(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.exoplanets[id]
	if !exists {
		return errors.New("exoplanet not found")
	}
	delete(s.exoplanets, id)
	return nil
}
