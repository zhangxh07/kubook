package cmdbsvc

type SvcResult struct {
	Services []*Service
	RunningTime string
}

type Service struct {
	Service_name   string   `json:"service_name"`
	Workload_type  string   `json:"workload_type"`
	Git_url         string  `json:"git_url"`
	Service_center string   `json:"service_center"`
	Status         string   `json:"status"`
}






