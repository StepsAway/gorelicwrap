package gorelicwrap

type MetricData struct {
	MetricData struct {
		From         string    `json:"from"`
		MetricsFound []string  `json:"metrics_found"`
		Metrics      []Metrics `json:"metrics"`
		To           string    `json:"to"`
	} `json:"metric_data"`
}

type Metrics struct {
	Name       string `json:"name"`
	Timeslices []struct {
		From   string       `json:"from"`
		To     string       `json:"to"`
		Values MetricValues `json:"values"`
	} `json:"timeslices"`
}

type MetricValues struct {
	AverageCallTime         float32 `json:"average_call_time,omitempty"`
	AverageExclusiveTime    float32 `json:"average_exclusive_time,omitempty"`
	AverageResponseTime     float32 `json:"average_response_time,omitempty"`
	AverageTime             float32 `json:"average_time,omitempty"`
	AverageValue            float32 `json:"average_value,omitempty"`
	AverageValuePerInstance float32 `json:"average_value_per_instance,omitempty"`
	BusyPercent             int     `json:"busy_percent,omitempty"`
	CallCount               int     `json:"call_count,omitempty"`
	CallsPerMinute          int     `json:"calls_per_minute,omitempty"`
	Count                   int     `json:"count,omitempty"`
	F                       int     `json:"f,omitempty"`
	InstanceCount           int     `json:"instance_count,omitempty"`
	MinResponseTime         float32 `json:"min_response_time,omitempty"`
	MaxResponseTime         float32 `json:"max_response_time,omitempty"`
	Percent                 float32 `json:"percent,omitempty"`
	RequestsPerMinute       float32 `json:"requests_per_minute,omitempty"`
	S                       int     `json:"s,omitempty"`
	Score                   float32 `json:"score,omitempty"`
	SessionsActive          int     `json:"sessions_active,omitempty"`
	StandardDeviation       float32 `json:"standard_deviation,omitempty"`
	T                       int     `json:"t,omitempty"`
	Threshold               float32 `json:"threshold,omitempty"`
	ThresholdMin            float32 `json:"threshold_min,omitempty"`
	TotalCallTimePerMinute  float32 `json:"total_call_time_per_minute,omitempty"`
	TotalTime               int     `json:"total_time,omitempty"`
	TotalUsedMb             float32 `json:"total_used_mb,omitempty"`
	UsedBytesByHost         float32 `json:"used_bytes_by_host,omitempty"`
	UsedMbByHost            float32 `json:"used_mb_by_host,omitempty"`
	Value                   float32 `json:"value,omitempty"`
}

type MetricParms struct {
	From      string   `url:"from,omitempty"`
	Names     []string `url:"names[],omitempty"`
	To        string   `url:"to,omitempty"`
	Raw       string   `url:"raw,omitempty"`
	Summarize string   `url:"summarize,omitempty"`
	Values    []string `url:"values[],omitempty"`
}
