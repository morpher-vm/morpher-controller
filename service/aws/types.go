package service

type GetEc2PriceQuery struct {
	Region      string
	MinVCpu     int
	MinMemoryGB float64
	OS          string
	MaxResults  int32
}

type EC2Instance struct {
	InstanceType    string `json:"instanceType"`
	VCpu            string `json:"vcpu"`
	Memory          string `json:"memory"`
	Storage         string `json:"storage"`
	NetworkPerf     string `json:"networkPerformance"`
	PricePerHour    string `json:"pricePerHour"`
	Description     string `json:"description"`
	Location        string `json:"location"`
	OperatingSystem string `json:"operatingSystem"`
}
