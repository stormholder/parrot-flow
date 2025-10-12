/**
 * Health Monitor Port
 *
 * Abstraction for reporting agent health and status.
 * Can be implemented with heartbeat messages, health check endpoints,
 * metrics systems (Prometheus), or logging systems.
 */

export interface AgentStatus {
  agentId: string;
  status: 'idle' | 'running' | 'error';
  currentRunId?: string;
  version: string;
  metadata?: Record<string, any>;
}

export interface IHealthMonitor {
  /**
   * Report current agent status
   * @param status - Current agent status information
   */
  reportStatus(status: AgentStatus): Promise<void>;

  /**
   * Start periodic health reporting
   * @param intervalMs - Interval in milliseconds
   */
  startMonitoring(intervalMs: number): void;

  /**
   * Stop health reporting
   */
  stopMonitoring(): void;
}
