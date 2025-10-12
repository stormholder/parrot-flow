/**
 * FindElement Node Executor
 *
 * Executes the 'findelement' node type.
 * Finds an element on the page using a CSS selector or XPath.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class FindElementNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get selector parameter
      const selectorParam = this.getRequiredParameter(parameters, 'selector');
      const selector = String(this.resolveValue(selectorParam, context));

      // Get optional timeout
      const timeout = this.getParameterWithDefault(
        parameters,
        'timeout',
        context.browserConfig?.timeout || 30000
      );

      // Get optional output variable name
      const outputVar = this.getParameter(parameters, 'outputVariable');

      // Wait for element to be present
      const startTime = Date.now();
      const element = await context.page.waitForSelector(selector, {
        timeout,
        state: 'attached'
      });
      const findTime = Date.now() - startTime;

      if (!element) {
        return this.failure(`Element not found: ${selector}`);
      }

      // Get element information
      const isVisible = await element.isVisible();
      const isEnabled = await element.isEnabled();
      const boundingBox = await element.boundingBox();

      // Store element reference if output variable specified
      if (outputVar) {
        context.setVariable(outputVar, {
          selector,
          found: true,
          visible: isVisible,
          enabled: isEnabled
        });
      }

      return this.success(
        {
          selector,
          found: true,
          visible: isVisible,
          enabled: isEnabled,
          position: boundingBox
        },
        {
          findTime
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to find element: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    return this.validateRequiredParameters(parameters, ['selector']);
  }
}
