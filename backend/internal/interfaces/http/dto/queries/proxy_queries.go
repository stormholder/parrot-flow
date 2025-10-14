package queries

// GetProxyRequest is the input for getting a single proxy
type GetProxyRequest struct {
	ID string `path:"id" doc:"Proxy ID"`
}

// GetProxyResponse is the output for a single proxy
type GetProxyResponse struct {
	Body ProxyDTO `json:"proxy"`
}

// ListProxiesRequest is the input for listing proxies
type ListProxiesRequest struct {
	Status *string  `query:"status" enum:"active,inactive,checking,failed" doc:"Filter by proxy status"`
	Tags   []string `query:"tags" doc:"Filter by tag IDs (proxy must have all specified tags)"`
}

// ListProxiesResponse is the output for listing proxies
type ListProxiesResponse struct {
	Body struct {
		Proxies []ProxyDTO `json:"proxies" doc:"List of proxies"`
		Total   int        `json:"total" doc:"Total number of proxies returned"`
	}
}

// GetActiveProxiesRequest is the input for getting active proxies
type GetActiveProxiesRequest struct {
	// No parameters needed
}

// GetActiveProxiesResponse is the output for active proxies
type GetActiveProxiesResponse struct {
	Body struct {
		Proxies []ProxyDTO `json:"proxies" doc:"List of active proxies"`
		Total   int        `json:"total" doc:"Total number of active proxies"`
	}
}

// ProxyDTO represents a proxy in API responses
type ProxyDTO struct {
	ID              string   `json:"id" doc:"Unique proxy identifier"`
	Name            string   `json:"name" doc:"Proxy name"`
	Host            string   `json:"host" doc:"Proxy host/IP address"`
	Port            int      `json:"port" doc:"Proxy port"`
	Protocol        string   `json:"protocol" doc:"Proxy protocol (http, https, socks5)"`
	Status          string   `json:"status" doc:"Current proxy status"`
	ConnectionURL   string   `json:"connection_url" doc:"Full proxy connection URL"`
	HasCredentials  bool     `json:"has_credentials" doc:"Whether proxy has authentication credentials"`
	Tags            []string `json:"tags" doc:"Associated tag IDs"`
	LastCheckedAt   *string  `json:"last_checked_at,omitempty" doc:"Last health check timestamp"`
	LastFailureAt   *string  `json:"last_failure_at,omitempty" doc:"Last failure timestamp"`
	FailureCount    int      `json:"failure_count" doc:"Number of consecutive failures"`
	SuccessCount    int      `json:"success_count" doc:"Total successful connections"`
	AverageLatency  int      `json:"average_latency_ms" doc:"Average latency in milliseconds"`
	CreatedAt       string   `json:"created_at" doc:"Creation timestamp"`
	UpdatedAt       string   `json:"updated_at" doc:"Last update timestamp"`
}
