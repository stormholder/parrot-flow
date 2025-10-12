/**
 * Goto Node Executor
 *
 * Executes the 'goto' node type.
 * Navigates to a specified URL in the browser.
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class GotoNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get URL parameter
      const urlParam = this.getRequiredParameter(parameters, 'url');
      const url = this.resolveValue(urlParam, context);

      // Get optional timeout (use browser config default or 30s)
      const timeout = this.getParameterWithDefault(
        parameters,
        'timeout',
        context.browserConfig?.timeout || 30000
      );

      // Navigate to URL
      const startTime = Date.now();
      await context.page.goto(url, {
        timeout,
        waitUntil: 'domcontentloaded'
      });
      const loadTime = Date.now() - startTime;

      // Get final URL (might be different due to redirects)
      const finalUrl = context.page.url();

      return this.success(
        {
          url: finalUrl,
          loadTime
        },
        {
          originalUrl: url,
          redirected: url !== finalUrl
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to navigate to URL: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    return this.validateRequiredParameters(parameters, ['url']);
  }
}
