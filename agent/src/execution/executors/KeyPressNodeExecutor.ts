/**
 * KeyPress Node Executor
 *
 * Executes the 'keypress' node type.
 * Simulates keyboard key presses (e.g., Enter, Tab, Escape).
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class KeyPressNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get key parameter
      const keyParam = this.getRequiredParameter(parameters, 'key');
      const key = String(this.resolveValue(keyParam, context));

      // Get optional parameters
      const selector = this.getParameter(parameters, 'selector');
      const delay = this.getParameterWithDefault(parameters, 'delay', 0);

      const startTime = Date.now();

      // If selector is provided, focus the element first
      if (selector) {
        const resolvedSelector = String(this.resolveValue(selector, context));
        const timeout = this.getParameterWithDefault(
          parameters,
          'timeout',
          context.browserConfig?.timeout || 30000
        );

        const element = await context.page.waitForSelector(resolvedSelector, {
          timeout,
          state: 'visible'
        });

        if (!element) {
          return this.failure(`Element not found: ${resolvedSelector}`);
        }

        await element.focus();
      }

      // Press the key
      await context.page.keyboard.press(key, { delay });

      const pressTime = Date.now() - startTime;

      return this.success(
        {
          key,
          pressed: true
        },
        {
          pressTime,
          delay,
          targeted: !!selector
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to press key: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    return this.validateRequiredParameters(parameters, ['key']);
  }
}
