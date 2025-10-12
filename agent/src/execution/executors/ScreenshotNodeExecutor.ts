/**
 * Screenshot Node Executor
 *
 * Executes the 'screenshot' node type.
 * Takes a screenshot of the page or a specific element.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class ScreenshotNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get optional parameters
      const selector = this.getParameter(parameters, 'selector');
      const fullPage = this.getParameterWithDefault(parameters, 'fullPage', false);
      const format = this.getParameterWithDefault(parameters, 'format', 'png') as 'png' | 'jpeg';
      const quality = this.getParameterWithDefault(parameters, 'quality', 80);
      const outputVar = this.getParameter(parameters, 'outputVariable');

      const startTime = Date.now();
      let screenshot: Buffer;

      // Take screenshot of element or full page
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

        screenshot = await element.screenshot({
          type: format,
          ...(format === 'jpeg' && { quality })
        });
      } else {
        screenshot = await context.page.screenshot({
          fullPage,
          type: format,
          ...(format === 'jpeg' && { quality })
        });
      }

      const screenshotTime = Date.now() - startTime;
      const base64 = screenshot.toString('base64');
      const size = screenshot.length;

      // Store screenshot in context if output variable specified
      if (outputVar) {
        context.setVariable(outputVar, {
          base64,
          format,
          size
        });
      }

      return this.success(
        {
          screenshot: base64,
          format,
          size,
          fullPage: !selector && fullPage
        },
        {
          screenshotTime,
          targeted: !!selector
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to take screenshot: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    // Screenshot has no required parameters
    return { valid: true };
  }
}
