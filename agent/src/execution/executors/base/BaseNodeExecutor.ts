/**
 * Base Node Executor
 *
 * Abstract base class that provides common functionality for all node executors.
 * Implements shared logic for parameter handling, validation, and error handling.
 */

import type { Node, Parameter } from '../../../types/generated/messages.js';
import type { ExecutionContext } from '../../context/ExecutionContext.js';
import type { INodeExecutor, NodeExecutionResult } from './INodeExecutor.js';

export abstract class BaseNodeExecutor implements INodeExecutor {
  /**
   * Execute the node (implemented by subclasses)
   */
  abstract execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult>;

  /**
   * Validate the node (can be overridden by subclasses)
   */
  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    // Default validation - can be overridden
    return { valid: true };
  }

  /**
   * Helper: Get a parameter value by name
   */
  protected getParameter(parameters: Parameter[], name: string): any {
    const param = parameters.find(p => p.name === name);
    return param?.value;
  }

  /**
   * Helper: Get a required parameter value by name
   * Throws error if parameter is missing
   */
  protected getRequiredParameter(parameters: Parameter[], name: string): any {
    const value = this.getParameter(parameters, name);
    if (value === undefined || value === null) {
      throw new Error(`Required parameter '${name}' is missing`);
    }
    return value;
  }

  /**
   * Helper: Get parameter with default value
   */
  protected getParameterWithDefault<T>(
    parameters: Parameter[],
    name: string,
    defaultValue: T
  ): T {
    const value = this.getParameter(parameters, name);
    return value !== undefined && value !== null ? value : defaultValue;
  }

  /**
   * Helper: Resolve a parameter value (handles variable references)
   * If value starts with '$', treat it as a variable reference
   */
  protected resolveValue(value: any, context: ExecutionContext): any {
    if (typeof value === 'string' && value.startsWith('$')) {
      const varName = value.substring(1);
      if (context.hasVariable(varName)) {
        return context.getVariable(varName);
      }
      throw new Error(`Variable '${varName}' not found in context`);
    }
    return value;
  }

  /**
   * Helper: Resolve all parameter values
   */
  protected resolveParameters(
    parameters: Parameter[],
    context: ExecutionContext
  ): Record<string, any> {
    const resolved: Record<string, any> = {};
    for (const param of parameters) {
      try {
        resolved[param.name] = this.resolveValue(param.value, context);
      } catch (error) {
        // If variable not found, use the literal value
        resolved[param.name] = param.value;
      }
    }
    return resolved;
  }

  /**
   * Helper: Create a success result
   */
  protected success(output?: Record<string, any>, metadata?: Record<string, any>): NodeExecutionResult {
    return {
      success: true,
      output,
      metadata
    };
  }

  /**
   * Helper: Create a failure result
   */
  protected failure(error: string, metadata?: Record<string, any>): NodeExecutionResult {
    return {
      success: false,
      error,
      metadata
    };
  }

  /**
   * Helper: Validate required parameters exist
   */
  protected validateRequiredParameters(
    parameters: Parameter[],
    requiredParams: string[]
  ): { valid: boolean; error?: string } {
    for (const requiredParam of requiredParams) {
      const param = parameters.find(p => p.name === requiredParam);
      if (!param || param.value === undefined || param.value === null) {
        return {
          valid: false,
          error: `Required parameter '${requiredParam}' is missing`
        };
      }
    }
    return { valid: true };
  }

  /**
   * Helper: Wait for a duration (useful for animations, page loads, etc.)
   */
  protected async wait(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}
