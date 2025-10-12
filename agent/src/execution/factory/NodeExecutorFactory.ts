/**
 * Node Executor Factory
 *
 * Factory class that creates the appropriate executor instance based on node type.
 * Implements the Factory pattern for creating node executors.
 */

import type { Node } from "../../types/generated/messages.js";
import type { INodeExecutor } from "../executors/base/INodeExecutor.js";

// Import all executor implementations
import { StartNodeExecutor } from "../executors/StartNodeExecutor.js";
import { GotoNodeExecutor } from "../executors/GotoNodeExecutor.js";
import { WaitDurationNodeExecutor } from "../executors/WaitDurationNodeExecutor.js";
import { FindElementNodeExecutor } from "../executors/FindElementNodeExecutor.js";
import { ClickNodeExecutor } from "../executors/ClickNodeExecutor.js";
import { InputDataNodeExecutor } from "../executors/InputDataNodeExecutor.js";
import { KeyPressNodeExecutor } from "../executors/KeyPressNodeExecutor.js";
import { ScreenshotNodeExecutor } from "../executors/ScreenshotNodeExecutor.js";
import { LoadDataNodeExecutor } from "../executors/LoadDataNodeExecutor.js";
import { ExtractDataNodeExecutor } from "../executors/ExtractDataNodeExecutor.js";

/**
 * Factory for creating node executors
 */
export class NodeExecutorFactory {
  // Executor registry (singleton instances)
  private static readonly executors: Map<string, INodeExecutor> = new Map();

  /**
   * Initialize the factory with executor instances
   * This is called once when the factory is first used
   */
  private static initialize(): void {
    if (this.executors.size === 0) {
      this.executors.set("start", new StartNodeExecutor());
      this.executors.set("goto", new GotoNodeExecutor());
      this.executors.set("waitduration", new WaitDurationNodeExecutor());
      this.executors.set("findelement", new FindElementNodeExecutor());
      this.executors.set("click", new ClickNodeExecutor());
      this.executors.set("inputdata", new InputDataNodeExecutor());
      this.executors.set("keypress", new KeyPressNodeExecutor());
      this.executors.set("screenshot", new ScreenshotNodeExecutor());
      this.executors.set("loaddata", new LoadDataNodeExecutor());
      this.executors.set("extractdata", new ExtractDataNodeExecutor());
    }
  }

  /**
   * Create an executor for the given node type
   *
   * @param node - The node to create an executor for
   * @returns The appropriate executor instance
   * @throws Error if node type is not supported
   */
  static create(node: Node): INodeExecutor {
    this.initialize();

    const executor = this.executors.get(node.node_type);

    if (!executor) {
      throw new Error(`Unsupported node type: ${node.node_type}`);
    }

    return executor;
  }

  /**
   * Check if a node type is supported
   *
   * @param nodeType - The node type to check
   * @returns True if the node type is supported
   */
  static isSupported(nodeType: string): boolean {
    this.initialize();
    return this.executors.has(nodeType);
  }

  /**
   * Get all supported node types
   *
   * @returns Array of supported node types
   */
  static getSupportedNodeTypes(): string[] {
    this.initialize();
    return Array.from(this.executors.keys());
  }

  /**
   * Register a custom executor for a node type
   * This allows extending the factory with custom node types
   *
   * @param nodeType - The node type
   * @param executor - The executor instance
   */
  static register(nodeType: string, executor: INodeExecutor): void {
    this.initialize();
    this.executors.set(nodeType, executor);
  }
}
