/**
 * Execution Context
 *
 * Maintains the state and shared resources during scenario execution.
 * This context is passed to each node executor and allows them to:
 * - Access the browser/page instances
 * - Store and retrieve variables (for data flow between nodes)
 * - Report progress events
 * - Access execution configuration
 */

import type { Browser, Page } from 'playwright';
import type { ProgressEvent, BrowserConfig } from '../../types/generated/messages.js';

export interface ExecutionContextConfig {
  runId: string;
  scenarioId: string;
  browserConfig?: BrowserConfig;
  onProgress?: (event: ProgressEvent) => void | Promise<void>;
}

export class ExecutionContext {
  // Browser resources
  public readonly browser: Browser;
  public readonly page: Page;

  // Execution metadata
  public readonly runId: string;
  public readonly scenarioId: string;
  public readonly browserConfig?: BrowserConfig;

  // Variable storage (for data flow between nodes)
  private variables: Map<string, any>;

  // Progress callback
  private onProgress?: (event: ProgressEvent) => void | Promise<void>;

  // Execution state
  private cancelled: boolean = false;
  private paused: boolean = false;

  constructor(
    browser: Browser,
    page: Page,
    config: ExecutionContextConfig
  ) {
    this.browser = browser;
    this.page = page;
    this.runId = config.runId;
    this.scenarioId = config.scenarioId;
    this.browserConfig = config.browserConfig;
    this.onProgress = config.onProgress;
    this.variables = new Map();
  }

  /**
   * Set a variable in the execution context
   */
  setVariable(name: string, value: any): void {
    this.variables.set(name, value);
  }

  /**
   * Get a variable from the execution context
   */
  getVariable(name: string): any {
    return this.variables.get(name);
  }

  /**
   * Check if a variable exists
   */
  hasVariable(name: string): boolean {
    return this.variables.has(name);
  }

  /**
   * Get all variables
   */
  getAllVariables(): Record<string, any> {
    return Object.fromEntries(this.variables);
  }

  /**
   * Report a progress event
   */
  async reportProgress(event: Omit<ProgressEvent, 'run_id' | 'timestamp'>): Promise<void> {
    if (this.onProgress) {
      const fullEvent: ProgressEvent = {
        ...event,
        run_id: this.runId,
        timestamp: new Date().toISOString()
      } as ProgressEvent;
      await this.onProgress(fullEvent);
    }
  }

  /**
   * Mark execution as cancelled
   */
  cancel(): void {
    this.cancelled = true;
  }

  /**
   * Check if execution is cancelled
   */
  isCancelled(): boolean {
    return this.cancelled;
  }

  /**
   * Pause execution
   */
  pause(): void {
    this.paused = true;
  }

  /**
   * Resume execution
   */
  resume(): void {
    this.paused = false;
  }

  /**
   * Check if execution is paused
   */
  isPaused(): boolean {
    return this.paused;
  }

  /**
   * Wait while paused (for implementing pause functionality)
   */
  async waitIfPaused(): Promise<void> {
    while (this.paused && !this.cancelled) {
      await new Promise(resolve => setTimeout(resolve, 100));
    }
  }

  /**
   * Cleanup resources
   */
  async cleanup(): Promise<void> {
    try {
      await this.page?.close();
      await this.browser?.close();
    } catch (error) {
      console.error('Error during cleanup:', error);
    }
  }
}
