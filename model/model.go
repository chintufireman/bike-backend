package model

type Bike struct {
	Name                          string `json:"name" unique:"true"`
	Price                         string `json:"price"`
	Colour                        string `json:"colour"`
	Displacement                  string `json:"displacement"`
	Mileage                       string `json:"mileage"`
	Weight                        string `json:"weight"`
	FuelCapacity                  string `json:"fuelCapacity"`
	SeatHeight                    string `json:"seatHeight"`
	Summary                       string `json:"summary"`
	MaxPower                      string `json:"maxPower"`
	MaxTorque                     string `json:"maxTorque"`
	FrontSuspension               string `json:"frontSuspension"`
	RearSuspension                string `json:"rearSuspension"`
	BrakingSystem                 string `json:"brakingSystem"`
	FrontBrakeType                string `json:"frontBrakeType"`
	FrontBrakeSize                string `json:"frontBrakeSize"`
	RearBrakeType                 string `json:"rearBrakeType"`
	RearBrakeSize                 string `json:"rearBrakeSize"`
	CalliperType                  string `json:"calliperType"`
	WheelType                     string `json:"wheelType"`
	FrontWheelSize                string `json:"frontWheelSize"`
	RearWheelSize                 string `json:"rearWheelSize"`
	FrontTyreSize                 string `json:"frontTyreSize"`
	RearTyreSize                  string `json:"rearTyreSize"`
	TyreType                      string `json:"tyreType"`
	RadialTyres                   string `json:"radialTyres"`
	FrontTyrePressureRider        string `json:"frontTyrePressureRider"`
	RearTyrePressureRider         string `json:"rearTyrePressureRider"`
	FrontTyrePressureRiderPillion string `json:"frontTyrePressureRiderPillion"`
	RearTyrePressureRiderPillion  string `json:"rearTyrePressureRiderPillion"`
	KerbWeight                    string `json:"kerbWeight"`
	GroundClearance               string `json:"groundClearance"`
	OverallLength                 string `json:"overallLength"`
	OverallWidth                  string `json:"overallWidth"`
	OverallHeight                 string `json:"overallHeight"`
	Wheelbase                     string `json:"wheelbase"`
	ChassisType                   string `json:"chassisType"`
	Odometer                      string `json:"odometer"`
	Speedometer                   string `json:"speedometer"`
	FuelGauge                     string `json:"fuelGauge"`
	DigitalFuelGauge              string `json:"digitalFuelGauge"`
	Tachometer                    string `json:"tachometer"`
	StandAlarm                    string `json:"standAlarm"`
}
