/**
 * Node Executor Interface
 *
 * Defines the contract that all node executors must implement.
 * Each node type (goto, click, inputdata, etc.) will have its own
 * executor implementing this interface.
 */

import type { Node, Parameter } from '../../../types/generated/messages.js';
import type { ExecutionContext } from '../../context/ExecutionContext.js';

/**
 * Result of node execution
 */
export interface NodeExecutionResult {
  /**
   * Whether the execution was successful
   */
  success: boolean;

  /**
   * Output data from the node execution
   */
  output?: Record<string, any>;

  /**
   * Error message if execution failed
   */
  error?: string;

  /**
   * Additional metadata about the execution
   */
  metadata?: Record<string, any>;
}

/**
 * Interface for all node executors
 */
export interface INodeExecutor {
  /**
   * Execute the node
   *
   * @param node - The node to execute
   * @param parameters - Input parameters for this node
   * @param context - Shared execution context
   * @returns Result of the execution
   */
  execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult>;

  /**
   * Validate that the node can be executed with the given parameters
   *
   * @param node - The node to validate
   * @param parameters - Input parameters for this node
   * @returns Validation result with error message if invalid
   */
  validate(
    node: Node,
    parameters: Parameter[]
  ): { valid: boolean; error?: string };
}
