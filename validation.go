package main

import (
	"errors"
)

func validateExoplanet(planet Exoplanet) error {
	if planet.Name == "" {
		return errors.New("name is required")
	}
	if planet.Description == "" {
		return errors.New("description is required")
	}
	if planet.Distance < 10 || planet.Distance > 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if planet.Radius < 0.1 || planet.Radius > 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if planet.Type == Terrestrial {
		if planet.Mass < 0.1 || planet.Mass > 10 {
			return errors.New("mass must be between 0.1 and 10 Earth-mass units")
		}
	}
	if planet.Type != GasGiant && planet.Type != Terrestrial {
		return errors.New("type must be either GasGiant or Terrestrial")
	}
	return nil
}
