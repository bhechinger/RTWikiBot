package defs

type PathingDef struct {
	Description            PathingDef_sub1 `json:"Description"`
	GradeMultMaxAscending  float64         `json:"GradeMultMaxAscending"`
	GradeMultMaxDescending float64         `json:"GradeMultMaxDescending"`
	GradeMultiplier        float64         `json:"GradeMultiplier"`
	MaxGrade               float64         `json:"MaxGrade"`
	MaxStairSize           float64         `json:"MaxStairSize"`
	MaxSteepness           float64         `json:"MaxSteepness"`
	MinGrade               float64         `json:"MinGrade"`
	MoveCostBackward       float64         `json:"MoveCostBackward"`
	MoveCostNormal         float64         `json:"MoveCostNormal"`
	PathingAngleCost       float64         `json:"PathingAngleCost"`
}

type PathingDef_sub1 struct {
	Details string `json:"Details"`
	Icon    string `json:"Icon"`
	ID      string `json:"Id"`
	Name    string `json:"Name"`
}
