/**
 * Scenario Executor
 *
 * Main orchestrator for executing browser automation scenarios.
 * Manages the execution lifecycle:
 * 1. Initializes browser and execution context
 * 2. Validates and traverses the scenario graph
 * 3. Executes nodes in the correct order using node executors
 * 4. Handles errors and reports progress
 * 5. Cleans up resources
 */

import { chromium, firefox, webkit, type Browser } from 'playwright';
import type {
  ExecuteScenarioMessage,
  Node,
  NodeParameters,
  ProgressEvent
} from '../types/generated/messages.js';
import { ScenarioGraph } from '../graph/index.js';
import { ExecutionContext } from './context/ExecutionContext.js';
import { NodeExecutorFactory } from './factory/NodeExecutorFactory.js';

export interface ScenarioExecutorOptions {
  /**
   * Browser type to use (default: chromium)
   */
  browserType?: 'chromium' | 'firefox' | 'webkit';

  /**
   * Callback for progress events
   */
  onProgress?: (event: ProgressEvent) => void | Promise<void>;

  /**
   * Browser executable path (optional)
   */
  browserPath?: string;
}

export class ScenarioExecutor {
  private readonly options: ScenarioExecutorOptions;

  constructor(options: ScenarioExecutorOptions = {}) {
    this.options = {
      browserType: options.browserType || 'chromium',
      onProgress: options.onProgress,
      browserPath: options.browserPath
    };
  }

  /**
   * Execute a scenario
   *
   * @param message - The execution message containing scenario and parameters
   */
  async execute(message: ExecuteScenarioMessage): Promise<void> {
    let context: ExecutionContext | null = null;

    try {
      // Report run started
      await this.reportProgress(message, {
        event: 'run_started',
        data: {
          scenario_id: message.scenario_id,
          run_id: message.run_id
        }
      });

      // Step 1: Validate the scenario graph
      const graph = new ScenarioGraph(message.context);
      const validation = graph.validate();

      if (!validation.valid) {
        throw new Error(`Invalid scenario graph: ${validation.errors.join(', ')}`);
      }

      // Step 2: Initialize browser and execution context
      context = await this.initializeContext(message);

      // Step 3: Get execution order
      const executionOrder = graph.topologicalSort();

      // Step 4: Execute nodes in order
      for (const nodeId of executionOrder) {
        // Check if execution is cancelled
        if (context.isCancelled()) {
          await this.reportProgress(message, {
            event: 'run_cancelled',
            data: { reason: 'Cancelled by user' }
          });
          break;
        }

        // Wait if paused
        await context.waitIfPaused();

        // Find the node
        const node = message.context.blocks.find(n => n.id === nodeId);
        if (!node) {
          throw new Error(`Node not found: ${nodeId}`);
        }

        // Execute the node
        await this.executeNode(node, message, context);
      }

      // Step 5: Report completion
      if (!context.isCancelled()) {
        await this.reportProgress(message, {
          event: 'run_completed',
          data: {
            variables: context.getAllVariables()
          }
        });
      }
    } catch (error) {
      // Report failure
      await this.reportProgress(message, {
        event: 'run_failed',
        error: error instanceof Error ? error.message : String(error)
      });

      throw error;
    } finally {
      // Step 6: Cleanup resources
      if (context) {
        await context.cleanup();
      }
    }
  }

  /**
   * Execute a single node
   */
  private async executeNode(
    node: Node,
    message: ExecuteScenarioMessage,
    context: ExecutionContext
  ): Promise<void> {
    const startTime = Date.now();

    try {
      // Report node started
      await this.reportProgress(message, {
        event: 'node_started',
        node_id: node.id,
        data: {
          node_type: node.node_type
        }
      });

      // Get node parameters
      const nodeParams = this.getNodeParameters(node.id, message.input_data.parameters);

      // Get executor for this node type
      const executor = NodeExecutorFactory.create(node);

      // Validate node
      const validation = executor.validate(node, nodeParams);
      if (!validation.valid) {
        throw new Error(`Node validation failed: ${validation.error}`);
      }

      // Execute node
      const result = await executor.execute(node, nodeParams, context);

      // Handle result
      if (result.success) {
        // Store output in context if available
        if (result.output) {
          for (const [key, value] of Object.entries(result.output)) {
            context.setVariable(`${node.id}.${key}`, value);
          }
        }

        // Report node completed
        await this.reportProgress(message, {
          event: 'node_completed',
          node_id: node.id,
          execution_time_ms: Date.now() - startTime,
          data: {
            output: result.output,
            metadata: result.metadata
          }
        });
      } else {
        throw new Error(result.error || 'Node execution failed');
      }
    } catch (error) {
      // Report node failed
      await this.reportProgress(message, {
        event: 'node_failed',
        node_id: node.id,
        execution_time_ms: Date.now() - startTime,
        error: error instanceof Error ? error.message : String(error)
      });

      throw error;
    }
  }

  /**
   * Initialize browser and execution context
   */
  private async initializeContext(
    message: ExecuteScenarioMessage
  ): Promise<ExecutionContext> {
    // Launch browser
    const browser = await this.launchBrowser(message);

    // Create browser context
    const browserContext = await browser.newContext({
      viewport: message.browser_config?.viewport ? {
        width: message.browser_config.viewport.width || 1920,
        height: message.browser_config.viewport.height || 1080
      } : undefined,
      userAgent: message.browser_config?.userAgent,
      locale: message.browser_config?.locale,
      timezoneId: message.browser_config?.timezone
    });

    // Create page
    const page = await browserContext.newPage();

    // Set default timeout
    if (message.browser_config?.timeout) {
      page.setDefaultTimeout(message.browser_config.timeout);
    }

    // Create execution context
    return new ExecutionContext(browser, page, {
      runId: message.run_id,
      scenarioId: message.scenario_id,
      browserConfig: message.browser_config,
      onProgress: this.options.onProgress
    });
  }

  /**
   * Launch browser based on configuration
   */
  private async launchBrowser(message: ExecuteScenarioMessage): Promise<Browser> {
    const config = {
      headless: message.browser_config?.headless ?? true,
      ...(this.options.browserPath && { executablePath: this.options.browserPath })
    };

    switch (this.options.browserType) {
      case 'firefox':
        return await firefox.launch(config);
      case 'webkit':
        return await webkit.launch(config);
      case 'chromium':
      default:
        return await chromium.launch(config);
    }
  }

  /**
   * Get parameters for a specific node
   */
  private getNodeParameters(nodeId: string, allParameters: NodeParameters[]) {
    const nodeParams = allParameters.find(p => p.block_id === nodeId);
    return nodeParams?.input || [];
  }

  /**
   * Report a progress event
   */
  private async reportProgress(
    message: ExecuteScenarioMessage,
    event: Omit<ProgressEvent, 'run_id' | 'timestamp'>
  ): Promise<void> {
    if (this.options.onProgress) {
      const fullEvent: ProgressEvent = {
        ...event,
        run_id: message.run_id,
        timestamp: new Date().toISOString()
      } as ProgressEvent;
      await this.options.onProgress(fullEvent);
    }
  }

  /**
   * Get supported node types
   */
  static getSupportedNodeTypes(): string[] {
    return NodeExecutorFactory.getSupportedNodeTypes();
  }
}
