/**
 * LoadData Node Executor
 *
 * Executes the 'loaddata' node type.
 * Loads data into the execution context from external sources or static values.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class LoadDataNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get required parameters
      const dataParam = this.getRequiredParameter(parameters, 'data');
      const data = this.resolveValue(dataParam, context);

      // Get optional variable name to store the data
      const variableName = this.getParameter(parameters, 'variableName');

      // Store data in context if variable name is provided
      if (variableName) {
        context.setVariable(String(variableName), data);
      }

      // Get optional data type hint
      const dataType = this.getParameter(parameters, 'dataType') || typeof data;

      return this.success(
        {
          data,
          dataType,
          stored: !!variableName,
          variableName
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to load data: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    return this.validateRequiredParameters(parameters, ['data']);
  }
}
