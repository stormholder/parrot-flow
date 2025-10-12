/**
 * Agent Service
 *
 * Main application service that orchestrates scenario execution.
 * This is the core business logic coordinator that:
 * - Maintains agent state
 * - Handles execution requests
 * - Reports progress and errors
 * - Manages health monitoring
 *
 * Uses dependency injection for all infrastructure concerns.
 */

import { nanoid } from 'nanoid';
import type { IMessageConsumer } from '../ports/messaging/IMessageConsumer.js';
import type { IMessagePublisher } from '../ports/messaging/IMessagePublisher.js';
import type { IHealthMonitor } from '../ports/monitoring/IHealthMonitor.js';
import { ScenarioExecutor } from '../execution/index.js';
import type { ExecuteScenarioMessage, ProgressEvent } from '../types/generated/messages.js';

export interface AgentServiceConfig {
  agentId?: string;
  version: string;
  browserType?: 'chromium' | 'firefox' | 'webkit';
  browserPath?: string;
  messageConsumer: IMessageConsumer<ExecuteScenarioMessage>;
  messagePublisher: IMessagePublisher;
  healthMonitor: IHealthMonitor;
  requestQueueName: string;
}

export class AgentService {
  private readonly agentId: string;
  private readonly config: AgentServiceConfig;
  private currentRunId: string | null = null;
  private status: 'idle' | 'running' | 'error' = 'idle';

  constructor(config: AgentServiceConfig) {
    this.config = config;
    this.agentId = config.agentId || nanoid();
  }

  /**
   * Get the agent ID
   */
  getAgentId(): string {
    return this.agentId;
  }

  /**
   * Get the current agent status
   */
  getStatus(): 'idle' | 'running' | 'error' {
    return this.status;
  }

  /**
   * Get the current run ID (if executing)
   */
  getCurrentRunId(): string | null {
    return this.currentRunId;
  }

  /**
   * Start processing execution requests
   */
  async start(): Promise<void> {
    console.log(`[Agent ${this.agentId}] Starting...`);

    // Update status
    this.updateStatus('idle', null);

    // Start consuming messages
    await this.config.messageConsumer.consume(
      this.config.requestQueueName,
      (message) => this.handleExecutionRequest(message)
    );

    console.log(`[Agent ${this.agentId}] Ready to receive execution requests`);
  }

  /**
   * Stop processing requests
   */
  async stop(): Promise<void> {
    console.log(`[Agent ${this.agentId}] Stopping...`);
    await this.config.messageConsumer.stop();
    console.log(`[Agent ${this.agentId}] Stopped`);
  }

  /**
   * Handle an incoming execution request
   */
  private async handleExecutionRequest(message: ExecuteScenarioMessage): Promise<void> {
    console.log(
      `\n[Agent ${this.agentId}] Starting run ${message.run_id} for scenario ${message.scenario_id}`
    );

    // Update status
    this.updateStatus('running', message.run_id);

    try {
      // Ensure progress queue exists
      const progressQueue = message.reply_queue || `agent.progress.${message.run_id}`;
      await this.config.messagePublisher.assertQueue(progressQueue);

      // Create executor with progress callback
      const executor = new ScenarioExecutor({
        browserType: this.config.browserType || 'chromium',
        browserPath: this.config.browserPath,
        onProgress: async (event: ProgressEvent) => {
          await this.publishProgressEvent(progressQueue, event);
        }
      });

      // Execute scenario
      await executor.execute(message);

      console.log(`[Agent ${this.agentId}] Completed run ${message.run_id}`);

      // Update status
      this.updateStatus('idle', null);
    } catch (error) {
      console.error(`[Agent ${this.agentId}] Execution error:`, error);

      // Send failure event
      await this.publishFailureEvent(message, error);

      // Update status
      this.updateStatus('error', null);

      // Reset to idle after a brief delay
      setTimeout(() => {
        if (this.status === 'error' && this.currentRunId === null) {
          this.updateStatus('idle', null);
        }
      }, 5000);
    }
  }

  /**
   * Publish a progress event
   */
  private async publishProgressEvent(
    queueName: string,
    event: ProgressEvent
  ): Promise<void> {
    try {
      await this.config.messagePublisher.publish(queueName, event, {
        persistent: true
      });
    } catch (error) {
      console.error(`[Agent ${this.agentId}] Failed to publish progress event:`, error);
    }
  }

  /**
   * Publish a failure event
   */
  private async publishFailureEvent(
    message: ExecuteScenarioMessage,
    error: unknown
  ): Promise<void> {
    try {
      const progressQueue = message.reply_queue || `agent.progress.${message.run_id}`;

      const failureEvent: ProgressEvent = {
        run_id: message.run_id,
        event: 'run_failed',
        error: error instanceof Error ? error.message : String(error),
        timestamp: new Date().toISOString()
      };

      await this.config.messagePublisher.publish(progressQueue, failureEvent, {
        persistent: true
      });
    } catch (err) {
      console.error(`[Agent ${this.agentId}] Failed to publish failure event:`, err);
    }
  }

  /**
   * Update agent status and report to health monitor
   */
  private updateStatus(status: 'idle' | 'running' | 'error', runId: string | null): void {
    this.status = status;
    this.currentRunId = runId;

    // Report to health monitor
    this.config.healthMonitor.reportStatus({
      agentId: this.agentId,
      status: this.status,
      currentRunId: this.currentRunId || undefined,
      version: this.config.version
    }).catch((error) => {
      console.error(`[Agent ${this.agentId}] Failed to report status:`, error);
    });
  }
}
