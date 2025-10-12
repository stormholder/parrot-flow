/**
 * Start Node Executor
 *
 * Executes the 'start' node type.
 * This is typically the entry point of a scenario and doesn't perform any action.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class StartNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    // Start node doesn't do anything - it's just a marker
    // We can use it to initialize variables if needed
    return this.success({ message: 'Scenario started' });
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    // Start node has no required parameters
    return { valid: true };
  }
}
