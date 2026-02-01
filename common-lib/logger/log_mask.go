package logger

type LogMap struct {
	Level       string   `json:"level"`
	ReqId       string   `json:"req_id"`
	EventTime   string   `json:"event_time"`
	StatusCode  any      `json:"status_code"`
	Method      string   `json:"method,omitempty"`
	Path        string   `json:"path"`
	Payload     any      `json:"payload"`
	ConnType    string   `json:"conn_type,omitempty"`
	Sequence    string   `json:"sequence,omitempty"`
	ErrorSource string   `json:"error_source,omitempty"`
	ErrorValue  string   `json:"error_value,omitempty"`
	ErrorStack  []string `json:"error_stack,omitempty"`
}
