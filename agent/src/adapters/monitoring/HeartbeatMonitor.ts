/**
 * Heartbeat Monitor Adapter
 *
 * Implements health monitoring by sending periodic heartbeat messages
 * to a message queue. This allows the backend to track agent availability
 * and status.
 */

import type {
  IHealthMonitor,
  AgentStatus,
} from "../../ports/monitoring/IHealthMonitor.js";
import type { IMessagePublisher } from "../../ports/messaging/IMessagePublisher.js";

export interface HeartbeatConfig {
  queueName: string;
  publisher: IMessagePublisher;
}

export class HeartbeatMonitor implements IHealthMonitor {
  private readonly config: HeartbeatConfig;
  private intervalId: NodeJS.Timeout | null = null;
  private currentStatus: AgentStatus | null = null;

  constructor(config: HeartbeatConfig) {
    this.config = config;
  }

  /**
   * Report current agent status
   */
  async reportStatus(status: AgentStatus): Promise<void> {
    this.currentStatus = status;

    try {
      const heartbeat = {
        agent_id: status.agentId,
        status: status.status,
        current_run_id: status.currentRunId,
        timestamp: new Date().toISOString(),
        metadata: {
          version: status.version,
          platform: process.platform,
          node_version: process.version,
          ...status.metadata,
        },
      };

      await this.config.publisher.publish(
        this.config.queueName,
        heartbeat,
        { persistent: false } // Heartbeats don't need to be persisted
      );
    } catch (error) {
      console.error("[Heartbeat] Failed to report status:", error);
      // Don't throw - heartbeat failures shouldn't crash the agent
    }
  }

  /**
   * Start periodic health reporting
   */
  startMonitoring(intervalMs: number): void {
    if (this.intervalId) {
      console.warn("[Heartbeat] Monitoring already started");
      return;
    }

    // Send initial heartbeat
    if (this.currentStatus) {
      this.reportStatus(this.currentStatus).catch(console.error);
    }

    // Start periodic heartbeats
    this.intervalId = setInterval(() => {
      if (this.currentStatus) {
        this.reportStatus(this.currentStatus).catch(console.error);
      }
    }, intervalMs);

    console.log(`[Heartbeat] Started monitoring (interval: ${intervalMs}ms)`);
  }

  /**
   * Stop health reporting
   */
  stopMonitoring(): void {
    if (this.intervalId) {
      clearInterval(this.intervalId);
      this.intervalId = null;
      console.log("[Heartbeat] Stopped monitoring");
    }
  }
}
