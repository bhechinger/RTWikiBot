package main

type Movement struct {
	Rating  int64
	Tonnage int64
}

func (m Movement) getMovementPoints() float64 {
	return float64(m.Rating / m.Tonnage)
}

func RoundBy5(value float64) int64 {
	return int64(value) / 5 * 5
}

func (m Movement) CalcWalkDistance() int64 {
	// numbers the result of the best fit line for the game movement
	var walkSpeedFixed = 26.05
	var walkSpeedMultiplier = 23.14

	return RoundBy5(walkSpeedFixed + m.getMovementPoints()*walkSpeedMultiplier)
}

func (m Movement) CalcSprintDistance() int64 {
	// numbers the result of the best fit line for the game movement
	var runSpeedFixed = 52.43
	var runSpeedMultiplier = 37.29

	return RoundBy5(runSpeedFixed + m.getMovementPoints()*runSpeedMultiplier)
}
