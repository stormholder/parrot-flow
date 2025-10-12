/**
 * WaitDuration Node Executor
 *
 * Executes the 'waitduration' node type.
 * Waits for a specified duration in milliseconds.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class WaitDurationNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get duration parameter
      const durationParam = this.getRequiredParameter(parameters, 'duration');
      const duration = Number(this.resolveValue(durationParam, context));

      if (isNaN(duration) || duration < 0) {
        return this.failure('Duration must be a positive number');
      }

      // Wait for the specified duration
      const startTime = Date.now();
      await this.wait(duration);
      const actualWait = Date.now() - startTime;

      return this.success(
        { waited: actualWait },
        { requestedDuration: duration }
      );
    } catch (error) {
      return this.failure(
        `Failed to wait: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    const validation = this.validateRequiredParameters(parameters, ['duration']);
    if (!validation.valid) {
      return validation;
    }

    // Validate that duration is a number
    const durationParam = this.getParameter(parameters, 'duration');
    const duration = Number(durationParam);
    if (isNaN(duration) || duration < 0) {
      return {
        valid: false,
        error: 'Duration must be a positive number'
      };
    }

    return { valid: true };
  }
}
