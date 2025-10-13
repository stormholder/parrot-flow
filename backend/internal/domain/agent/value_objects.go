package agent

import (
	"errors"
	"fmt"
)

// BrowserType represents the type of browser
type BrowserType struct {
	value string
}

func NewBrowserType(value string) (BrowserType, error) {
	switch value {
	case "chromium", "firefox", "webkit":
		return BrowserType{value: value}, nil
	default:
		return BrowserType{}, fmt.Errorf("invalid browser type: %s (must be chromium, firefox, or webkit)", value)
	}
}

func (bt BrowserType) String() string {
	return bt.value
}

// Common browser types
var (
	BrowserChromium = BrowserType{value: "chromium"}
	BrowserFirefox  = BrowserType{value: "firefox"}
	BrowserWebkit   = BrowserType{value: "webkit"}
)

// BrowserCapability represents a single browser capability
type BrowserCapability struct {
	Type            BrowserType
	Version         string
	Headless        bool
	SupportedExtensions []string
}

func NewBrowserCapability(browserType BrowserType, version string, headless bool) (BrowserCapability, error) {
	if version == "" {
		return BrowserCapability{}, errors.New("browser version cannot be empty")
	}
	return BrowserCapability{
		Type:            browserType,
		Version:         version,
		Headless:        headless,
		SupportedExtensions: make([]string, 0),
	}, nil
}

// Platform represents the operating system platform
type Platform struct {
	value string
}

func NewPlatform(value string) (Platform, error) {
	switch value {
	case "linux", "darwin", "windows":
		return Platform{value: value}, nil
	default:
		return Platform{}, fmt.Errorf("invalid platform: %s (must be linux, darwin, or windows)", value)
	}
}

func (p Platform) String() string {
	return p.value
}

// Common platforms
var (
	PlatformLinux   = Platform{value: "linux"}
	PlatformDarwin  = Platform{value: "darwin"}
	PlatformWindows = Platform{value: "windows"}
)

// Architecture represents the CPU architecture
type Architecture struct {
	value string
}

func NewArchitecture(value string) (Architecture, error) {
	switch value {
	case "amd64", "arm64", "386", "arm":
		return Architecture{value: value}, nil
	default:
		return Architecture{}, fmt.Errorf("invalid architecture: %s", value)
	}
}

func (a Architecture) String() string {
	return a.value
}

// Common architectures
var (
	ArchAMD64 = Architecture{value: "amd64"}
	ArchARM64 = Architecture{value: "arm64"}
	Arch386   = Architecture{value: "386"}
	ArchARM   = Architecture{value: "arm"}
)

// OSInfo represents operating system information
type OSInfo struct {
	Platform     Platform
	Architecture Architecture
	Version      string
}

func NewOSInfo(platform Platform, arch Architecture, version string) (OSInfo, error) {
	if version == "" {
		return OSInfo{}, errors.New("os version cannot be empty")
	}
	return OSInfo{
		Platform:     platform,
		Architecture: arch,
		Version:      version,
	}, nil
}

// ProxyCapability represents proxy support capabilities
type ProxyCapability struct {
	SupportsProxy       bool
	SupportedProtocols  []string // e.g., ["http", "https", "socks5"]
}

func NewProxyCapability(supportsProxy bool, protocols []string) ProxyCapability {
	if protocols == nil {
		protocols = make([]string, 0)
	}
	return ProxyCapability{
		SupportsProxy:      supportsProxy,
		SupportedProtocols: protocols,
	}
}

// ResourceLimits represents agent resource limits
type ResourceLimits struct {
	MaxConcurrentRuns int
	MaxMemoryMB       int
	MaxCPUCores       int
}

func NewResourceLimits(maxConcurrentRuns, maxMemoryMB, maxCPUCores int) (ResourceLimits, error) {
	if maxConcurrentRuns <= 0 {
		return ResourceLimits{}, errors.New("max concurrent runs must be positive")
	}
	if maxMemoryMB <= 0 {
		return ResourceLimits{}, errors.New("max memory must be positive")
	}
	if maxCPUCores <= 0 {
		return ResourceLimits{}, errors.New("max CPU cores must be positive")
	}
	return ResourceLimits{
		MaxConcurrentRuns: maxConcurrentRuns,
		MaxMemoryMB:       maxMemoryMB,
		MaxCPUCores:       maxCPUCores,
	}, nil
}

// Capabilities represents all capabilities of an agent
type Capabilities struct {
	Browsers       []BrowserCapability
	OS             OSInfo
	Proxy          ProxyCapability
	ResourceLimits ResourceLimits
	Features       []string // Additional features: "screenshots", "video-recording", "network-interception", etc.
}

func NewCapabilities(
	browsers []BrowserCapability,
	osInfo OSInfo,
	proxyCapability ProxyCapability,
	resourceLimits ResourceLimits,
	features []string,
) (Capabilities, error) {
	if len(browsers) == 0 {
		return Capabilities{}, errors.New("agent must support at least one browser")
	}
	if features == nil {
		features = make([]string, 0)
	}
	return Capabilities{
		Browsers:       browsers,
		OS:             osInfo,
		Proxy:          proxyCapability,
		ResourceLimits: resourceLimits,
		Features:       features,
	}, nil
}

// HasBrowser checks if the agent supports a specific browser type
func (c Capabilities) HasBrowser(browserType BrowserType) bool {
	for _, browser := range c.Browsers {
		if browser.Type.String() == browserType.String() {
			return true
		}
	}
	return false
}

// HasFeature checks if the agent supports a specific feature
func (c Capabilities) HasFeature(feature string) bool {
	for _, f := range c.Features {
		if f == feature {
			return true
		}
	}
	return false
}

// SupportsProxyProtocol checks if the agent supports a specific proxy protocol
func (c Capabilities) SupportsProxyProtocol(protocol string) bool {
	if !c.Proxy.SupportsProxy {
		return false
	}
	for _, p := range c.Proxy.SupportedProtocols {
		if p == protocol {
			return true
		}
	}
	return false
}

// ConnectionInfo represents agent connection information
type ConnectionInfo struct {
	IPAddress string
	Hostname  string
	QueueName string // RabbitMQ queue name for this agent
}

func NewConnectionInfo(ipAddress, hostname, queueName string) (ConnectionInfo, error) {
	if queueName == "" {
		return ConnectionInfo{}, errors.New("queue name cannot be empty")
	}
	return ConnectionInfo{
		IPAddress: ipAddress,
		Hostname:  hostname,
		QueueName: queueName,
	}, nil
}
