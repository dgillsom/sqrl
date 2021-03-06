package sqrl

type Report struct {
	Ram     `json:"Ram"`
	Swap    `json:"Swap"`
	OS      `json:"OS"`
	CPUInfo `json:"CPUInfo"`
	inter   []Interface `json:"Interface"`
}

func MakeReport() (Report, error) {
	ram, os, swap, err := RamOSInfo()
	if err != nil {
		return Report{}, err
	}
	interfaces, err := GetInterfaces()
	if err != nil {
		return Report{}, err
	}
	cpu, err := GetCPUInfo()
	if err != nil {
		return Report{}, err
	}
	report := Report{ram, swap, os, cpu, interfaces}
	return report, err
}
