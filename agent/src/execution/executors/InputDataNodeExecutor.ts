/**
 * InputData Node Executor
 *
 * Executes the 'inputdata' node type.
 * Types text into an input field identified by a CSS selector.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class InputDataNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get required parameters
      const selectorParam = this.getRequiredParameter(parameters, 'selector');
      const selector = String(this.resolveValue(selectorParam, context));

      const textParam = this.getRequiredParameter(parameters, 'text');
      const text = String(this.resolveValue(textParam, context));

      // Get optional parameters
      const timeout = this.getParameterWithDefault(
        parameters,
        'timeout',
        context.browserConfig?.timeout || 30000
      );
      const clearFirst = this.getParameterWithDefault(parameters, 'clearFirst', true);
      const delay = this.getParameterWithDefault(parameters, 'delay', 0); // Delay between keystrokes (ms)

      // Wait for element to be visible
      const startTime = Date.now();
      const element = await context.page.waitForSelector(selector, {
        timeout,
        state: 'visible'
      });

      if (!element) {
        return this.failure(`Element not found: ${selector}`);
      }

      // Clear the input field first if requested
      if (clearFirst) {
        await element.fill('');
      }

      // Type the text
      if (delay > 0) {
        await element.type(text, { delay });
      } else {
        await element.fill(text);
      }

      const inputTime = Date.now() - startTime;

      // Verify the value was set
      const actualValue = await element.inputValue();

      return this.success(
        {
          selector,
          text: text,
          length: text.length,
          verified: actualValue === text
        },
        {
          inputTime,
          clearFirst,
          delay
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to input data: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    return this.validateRequiredParameters(parameters, ['selector', 'text']);
  }
}
