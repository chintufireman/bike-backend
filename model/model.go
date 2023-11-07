package model

type Bike struct {
	Name                          string  `json:"name" unique:"true"`
	Price                         float64 `json:"price"`
	Colour                        string  `json:"colour"`
	Displacement                  int     `json:"displacement"`
	Mileage                       float64 `json:"mileage"`
	Weight                        float64 `json:"weight"`
	FuelCapacity                  float64 `json:"fuelCapacity"`
	SeatHeight                    int     `json:"seatHeight"`
	Summary                       string  `json:"summary"`
	MaxPower                      float64 `json:"maxPower"`
	MaxTorque                     float64 `json:"maxTorque"`
	FrontSuspension               string  `json:"frontSuspension"`
	RearSuspension                string  `json:"rearSuspension"`
	BrakingSystem                 string  `json:"brakingSystem"`
	FrontBrakeType                string  `json:"frontBrakeType"`
	FrontBrakeSize                float64 `json:"frontBrakeSize"`
	RearBrakeType                 string  `json:"rearBrakeType"`
	RearBrakeSize                 float64 `json:"rearBrakeSize"`
	CalliperType                  string  `json:"calliperType"`
	WheelType                     string  `json:"wheelType"`
	FrontWheelSize                float64 `json:"frontWheelSize"`
	RearWheelSize                 float64 `json:"rearWheelSize"`
	FrontTyreSize                 string  `json:"frontTyreSize"`
	RearTyreSize                  string  `json:"rearTyreSize"`
	TyreType                      string  `json:"tyreType"`
	RadialTyres                   bool    `json:"radialTyres"`
	FrontTyrePressureRider        float64 `json:"frontTyrePressureRider"`
	RearTyrePressureRider         float64 `json:"rearTyrePressureRider"`
	FrontTyrePressureRiderPillion float64 `json:"frontTyrePressureRiderPillion"`
	RearTyrePressureRiderPillion  float64 `json:"rearTyrePressureRiderPillion"`
	KerbWeight                    float64 `json:"kerbWeight"`
	GroundClearance               float64 `json:"groundClearance"`
	OverallLength                 float64 `json:"overallLength"`
	OverallWidth                  float64 `json:"overallWidth"`
	OverallHeight                 float64 `json:"overallHeight"`
	Wheelbase                     float64 `json:"wheelbase"`
	ChassisType                   string  `json:"chassisType"`
	Odometer                      float64 `json:"odometer"`
	Speedometer                   string  `json:"speedometer"`
	FuelGauge                     bool    `json:"fuelGauge"`
	DigitalFuelGauge              bool    `json:"digitalFuelGauge"`
	Tachometer                    string  `json:"tachometer"`
	StandAlarm                    bool    `json:"standAlarm"`
}
