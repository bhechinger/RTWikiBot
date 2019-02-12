package main

type MoveDef struct {
	Description struct {
		Details string `json:"Details"`
		Icon    string `json:"Icon"`
		ID      string `json:"Id"`
		Name    string `json:"Name"`
	} `json:"Description"`
	LimpVelocity          float64 `json:"LimpVelocity"`
	MaxJumpAccel          float64 `json:"MaxJumpAccel"`
	MaxJumpVel            float64 `json:"MaxJumpVel"`
	MaxJumpVerticalAccel  float64 `json:"MaxJumpVerticalAccel"`
	MaxRadialAcceleration float64 `json:"MaxRadialAcceleration"`
	MaxRadialVelocity     float64 `json:"MaxRadialVelocity"`
	MaxSprintDistance     float64 `json:"MaxSprintDistance"`
	MaxWalkDistance       float64 `json:"MaxWalkDistance"`
	RunVelocity           float64 `json:"RunVelocity"`
	SprintAcceleration    float64 `json:"SprintAcceleration"`
	SprintVelocity        float64 `json:"SprintVelocity"`
	WalkAcceleration      float64 `json:"WalkAcceleration"`
	WalkVelocity          float64 `json:"WalkVelocity"`
}
