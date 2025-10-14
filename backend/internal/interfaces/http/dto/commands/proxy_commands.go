package commands

// CreateProxyRequest is the input for creating a new proxy
type CreateProxyRequest struct {
	Body struct {
		Name     string   `json:"name" minLength:"1" maxLength:"100" doc:"Proxy name (must be unique)"`
		Host     string   `json:"host" minLength:"1" doc:"Proxy host/IP address"`
		Port     int      `json:"port" minimum:"1" maximum:"65535" doc:"Proxy port"`
		Protocol string   `json:"protocol" enum:"http,https,socks5" doc:"Proxy protocol"`
		Username string   `json:"username,omitempty" doc:"Optional username for authentication"`
		Password string   `json:"password,omitempty" doc:"Optional password for authentication"`
		Tags     []string `json:"tags,omitempty" doc:"Optional tag IDs to associate with this proxy"`
	}
}

// CreateProxyResponse is the output after creating a proxy
type CreateProxyResponse struct {
	Body struct {
		ID              string   `json:"id" doc:"Unique proxy identifier"`
		Name            string   `json:"name" doc:"Proxy name"`
		Host            string   `json:"host" doc:"Proxy host/IP address"`
		Port            int      `json:"port" doc:"Proxy port"`
		Protocol        string   `json:"protocol" doc:"Proxy protocol"`
		Status          string   `json:"status" doc:"Current proxy status"`
		HasCredentials  bool     `json:"has_credentials" doc:"Whether proxy has authentication credentials"`
		Tags            []string `json:"tags" doc:"Associated tag IDs"`
		LastCheckedAt   *string  `json:"last_checked_at,omitempty" doc:"Last health check timestamp"`
		FailureCount    int      `json:"failure_count" doc:"Number of consecutive failures"`
		SuccessCount    int      `json:"success_count" doc:"Total successful connections"`
		AverageLatency  int      `json:"average_latency_ms" doc:"Average latency in milliseconds"`
		CreatedAt       string   `json:"created_at" doc:"Creation timestamp"`
		UpdatedAt       string   `json:"updated_at" doc:"Last update timestamp"`
	}
}

// UpdateProxyRequest is the input for updating an existing proxy
type UpdateProxyRequest struct {
	ID string `path:"id" doc:"Proxy ID"`
	Body struct {
		Name     *string `json:"name,omitempty" minLength:"1" maxLength:"100" doc:"Proxy name"`
		Host     *string `json:"host,omitempty" minLength:"1" doc:"Proxy host/IP address"`
		Port     *int    `json:"port,omitempty" minimum:"1" maximum:"65535" doc:"Proxy port"`
		Protocol *string `json:"protocol,omitempty" enum:"http,https,socks5" doc:"Proxy protocol"`
		Username *string `json:"username,omitempty" doc:"Username for authentication"`
		Password *string `json:"password,omitempty" doc:"Password for authentication"`
	}
}

// UpdateProxyResponse is the output after updating a proxy
type UpdateProxyResponse struct {
	Body struct {
		ID              string   `json:"id" doc:"Unique proxy identifier"`
		Name            string   `json:"name" doc:"Proxy name"`
		Host            string   `json:"host" doc:"Proxy host/IP address"`
		Port            int      `json:"port" doc:"Proxy port"`
		Protocol        string   `json:"protocol" doc:"Proxy protocol"`
		Status          string   `json:"status" doc:"Current proxy status"`
		HasCredentials  bool     `json:"has_credentials" doc:"Whether proxy has authentication credentials"`
		Tags            []string `json:"tags" doc:"Associated tag IDs"`
		LastCheckedAt   *string  `json:"last_checked_at,omitempty" doc:"Last health check timestamp"`
		FailureCount    int      `json:"failure_count" doc:"Number of consecutive failures"`
		SuccessCount    int      `json:"success_count" doc:"Total successful connections"`
		AverageLatency  int      `json:"average_latency_ms" doc:"Average latency in milliseconds"`
		CreatedAt       string   `json:"created_at" doc:"Creation timestamp"`
		UpdatedAt       string   `json:"updated_at" doc:"Last update timestamp"`
	}
}

// DeleteProxyRequest is the input for deleting a proxy
type DeleteProxyRequest struct {
	ID string `path:"id" doc:"Proxy ID"`
}

// DeleteProxyResponse is the output after deleting a proxy
type DeleteProxyResponse struct {
	Body struct {
		Message string `json:"message" doc:"Confirmation message"`
	}
}

// RecordHealthRequest is the input for recording proxy health check results
type RecordHealthRequest struct {
	ID string `path:"id" doc:"Proxy ID"`
	Body struct {
		Success   bool   `json:"success" doc:"Whether the health check succeeded"`
		LatencyMs int    `json:"latency_ms,omitempty" minimum:"0" doc:"Latency in milliseconds (for successful checks)"`
		ErrorMsg  string `json:"error_msg,omitempty" doc:"Error message (for failed checks)"`
	}
}

// RecordHealthResponse is the output after recording health check
type RecordHealthResponse struct {
	Body struct {
		ID              string   `json:"id" doc:"Unique proxy identifier"`
		Name            string   `json:"name" doc:"Proxy name"`
		Host            string   `json:"host" doc:"Proxy host/IP address"`
		Port            int      `json:"port" doc:"Proxy port"`
		Protocol        string   `json:"protocol" doc:"Proxy protocol"`
		Status          string   `json:"status" doc:"Current proxy status"`
		HasCredentials  bool     `json:"has_credentials" doc:"Whether proxy has authentication credentials"`
		Tags            []string `json:"tags" doc:"Associated tag IDs"`
		LastCheckedAt   *string  `json:"last_checked_at,omitempty" doc:"Last health check timestamp"`
		FailureCount    int      `json:"failure_count" doc:"Number of consecutive failures"`
		SuccessCount    int      `json:"success_count" doc:"Total successful connections"`
		AverageLatency  int      `json:"average_latency_ms" doc:"Average latency in milliseconds"`
		CreatedAt       string   `json:"created_at" doc:"Creation timestamp"`
		UpdatedAt       string   `json:"updated_at" doc:"Last update timestamp"`
	}
}

// ActivateProxyRequest is the input for activating a proxy
type ActivateProxyRequest struct {
	ID string `path:"id" doc:"Proxy ID"`
}

// ActivateProxyResponse is the output after activating a proxy
type ActivateProxyResponse struct {
	Body struct {
		ID     string `json:"id" doc:"Unique proxy identifier"`
		Name   string `json:"name" doc:"Proxy name"`
		Status string `json:"status" doc:"Current proxy status (should be 'active')"`
	}
}

// DeactivateProxyRequest is the input for deactivating a proxy
type DeactivateProxyRequest struct {
	ID string `path:"id" doc:"Proxy ID"`
}

// DeactivateProxyResponse is the output after deactivating a proxy
type DeactivateProxyResponse struct {
	Body struct {
		ID     string `json:"id" doc:"Unique proxy identifier"`
		Name   string `json:"name" doc:"Proxy name"`
		Status string `json:"status" doc:"Current proxy status (should be 'inactive')"`
	}
}
