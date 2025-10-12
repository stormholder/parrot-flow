/**
 * Click Node Executor
 *
 * Executes the 'click' node type.
 * Clicks on an element identified by a CSS selector or XPath.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class ClickNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get selector parameter
      const selectorParam = this.getRequiredParameter(parameters, 'selector');
      const selector = String(this.resolveValue(selectorParam, context));

      // Get optional parameters
      const timeout = this.getParameterWithDefault(
        parameters,
        'timeout',
        context.browserConfig?.timeout || 30000
      );
      const clickCount = this.getParameterWithDefault(parameters, 'clickCount', 1);
      const button = this.getParameterWithDefault(parameters, 'button', 'left') as 'left' | 'right' | 'middle';
      const waitForNavigation = this.getParameterWithDefault(parameters, 'waitForNavigation', false);

      // Wait for element to be visible and enabled
      const startTime = Date.now();
      const element = await context.page.waitForSelector(selector, {
        timeout,
        state: 'visible'
      });

      if (!element) {
        return this.failure(`Element not found: ${selector}`);
      }

      // Check if element is enabled
      const isEnabled = await element.isEnabled();
      if (!isEnabled) {
        return this.failure(`Element is disabled: ${selector}`);
      }

      // Perform click (with optional navigation wait)
      if (waitForNavigation) {
        await Promise.all([
          context.page.waitForLoadState('networkidle', { timeout }),
          element.click({ clickCount, button })
        ]);
      } else {
        await element.click({ clickCount, button });
      }

      const clickTime = Date.now() - startTime;

      return this.success(
        {
          selector,
          clicked: true,
          clickCount
        },
        {
          clickTime,
          button
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to click element: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    return this.validateRequiredParameters(parameters, ['selector']);
  }
}
