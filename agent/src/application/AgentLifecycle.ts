/**
 * Agent Lifecycle Manager
 *
 * Manages the lifecycle of the agent service:
 * - Initialization and startup
 * - Graceful shutdown on signals
 * - Resource cleanup
 * - Health monitoring coordination
 */

import type { AgentService } from "./AgentService.js";
import type { IHealthMonitor } from "../ports/monitoring/IHealthMonitor.js";
import { NodeExecutorFactory } from "../execution/index.js";

export interface LifecycleConfig {
  agentService: AgentService;
  healthMonitor: IHealthMonitor;
  onShutdown?: () => Promise<void>;
  heartbeatIntervalMs?: number;
}

export class AgentLifecycle {
  private readonly config: LifecycleConfig;
  private isShuttingDown = false;

  constructor(config: LifecycleConfig) {
    this.config = config;
  }

  /**
   * Start the agent and all its services
   */
  async start(): Promise<void> {
    this.printBanner();

    // Start agent service
    await this.config.agentService.start();

    // Start health monitoring
    const interval = this.config.heartbeatIntervalMs || 10000; // Default 10s
    this.config.healthMonitor.startMonitoring(interval);

    // Print ready message
    this.printReadyMessage();

    // Setup graceful shutdown handlers
    this.setupShutdownHandlers();
  }

  /**
   * Shutdown the agent gracefully
   */
  async shutdown(): Promise<void> {
    if (this.isShuttingDown) {
      console.log("[Lifecycle] Shutdown already in progress");
      return;
    }

    this.isShuttingDown = true;
    console.log("\n[Lifecycle] Shutting down gracefully...");

    try {
      // Stop health monitoring
      this.config.healthMonitor.stopMonitoring();

      // Stop agent service
      await this.config.agentService.stop();

      // Custom shutdown hook
      if (this.config.onShutdown) {
        await this.config.onShutdown();
      }

      console.log("[Lifecycle] Shutdown complete");
      process.exit(0);
    } catch (error) {
      console.error("[Lifecycle] Error during shutdown:", error);
      process.exit(1);
    }
  }

  /**
   * Setup handlers for graceful shutdown signals
   */
  private setupShutdownHandlers(): void {
    const shutdownHandler = () => {
      this.shutdown().catch((error) => {
        console.error("[Lifecycle] Shutdown error:", error);
        process.exit(1);
      });
    };

    process.on("SIGINT", shutdownHandler);
    process.on("SIGTERM", shutdownHandler);

    // Handle uncaught errors
    process.on("uncaughtException", (error) => {
      console.error("[Lifecycle] Uncaught exception:", error);
      this.shutdown().catch(() => process.exit(1));
    });

    process.on("unhandledRejection", (reason, promise) => {
      console.error(
        "[Lifecycle] Unhandled rejection at:",
        promise,
        "reason:",
        reason
      );
      this.shutdown().catch(() => process.exit(1));
    });
  }

  /**
   * Print startup banner
   */
  private printBanner(): void {
    const agentId = this.config.agentService.getAgentId();

    console.log("=".repeat(60));
    console.log("🦜 Parrot Flow Agent Service");
    console.log("=".repeat(60));
    console.log(`Agent ID: ${agentId}`);
    console.log(`Node: ${process.version}`);
    console.log(`Platform: ${process.platform}`);
    console.log("=".repeat(60));
  }

  /**
   * Print ready message with supported node types
   */
  private printReadyMessage(): void {
    console.log("=".repeat(60));
    console.log("\n✓ Agent is ready to receive execution requests\n");
    console.log("Supported node types:");

    const supportedTypes = NodeExecutorFactory.getSupportedNodeTypes();
    supportedTypes.forEach((type) => console.log(`  - ${type}`));

    console.log("\nWaiting for messages...\n");
  }
}
