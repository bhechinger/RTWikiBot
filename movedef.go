package main

type MoveDef struct {
	Description struct {
		Details string `json:"Details"`
		Icon    string `json:"Icon"`
		ID      string `json:"Id"`
		Name    string `json:"Name"`
	} `json:"Description"`
	LimpVelocity          int64 `json:"LimpVelocity"`
	MaxJumpAccel          int64 `json:"MaxJumpAccel"`
	MaxJumpVel            int64 `json:"MaxJumpVel"`
	MaxJumpVerticalAccel  int64 `json:"MaxJumpVerticalAccel"`
	MaxRadialAcceleration int64 `json:"MaxRadialAcceleration"`
	MaxRadialVelocity     int64 `json:"MaxRadialVelocity"`
	MaxSprintDistance     int64 `json:"MaxSprintDistance"`
	MaxWalkDistance       int64 `json:"MaxWalkDistance"`
	RunVelocity           int64 `json:"RunVelocity"`
	SprintAcceleration    int64 `json:"SprintAcceleration"`
	SprintVelocity        int64 `json:"SprintVelocity"`
	WalkAcceleration      int64 `json:"WalkAcceleration"`
	WalkVelocity          int64 `json:"WalkVelocity"`
}
