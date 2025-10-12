/**
 * ExtractData Node Executor
 *
 * Executes the 'extractdata' node type.
 * Extracts data from the page (text content, attributes, etc.).
 */

import type { Node, Parameter } from '../../types/generated/messages.js';
import type { ExecutionContext } from '../context/ExecutionContext.js';
import type { NodeExecutionResult } from './base/INodeExecutor.js';
import { BaseNodeExecutor } from './base/BaseNodeExecutor.js';

export class ExtractDataNodeExecutor extends BaseNodeExecutor {
  async execute(
    node: Node,
    parameters: Parameter[],
    context: ExecutionContext
  ): Promise<NodeExecutionResult> {
    try {
      // Get required parameters
      const selectorParam = this.getRequiredParameter(parameters, 'selector');
      const selector = String(this.resolveValue(selectorParam, context));

      // Get optional parameters
      const timeout = this.getParameterWithDefault(
        parameters,
        'timeout',
        context.browserConfig?.timeout || 30000
      );
      const extractType = this.getParameterWithDefault(parameters, 'extractType', 'textContent') as
        | 'textContent'
        | 'innerText'
        | 'innerHTML'
        | 'attribute'
        | 'value';
      const attribute = this.getParameter(parameters, 'attribute');
      const outputVar = this.getParameter(parameters, 'outputVariable');
      const multiple = this.getParameterWithDefault(parameters, 'multiple', false);

      const startTime = Date.now();
      let extractedData: any;

      if (multiple) {
        // Extract from multiple elements
        const elements = await context.page.$$(selector);

        if (elements.length === 0) {
          return this.failure(`No elements found: ${selector}`);
        }

        extractedData = [];
        for (const element of elements) {
          const data = await this.extractFromElement(element, extractType, attribute);
          extractedData.push(data);
        }
      } else {
        // Extract from single element
        const element = await context.page.waitForSelector(selector, {
          timeout,
          state: 'attached'
        });

        if (!element) {
          return this.failure(`Element not found: ${selector}`);
        }

        extractedData = await this.extractFromElement(element, extractType, attribute);
      }

      const extractTime = Date.now() - startTime;

      // Store extracted data in context if output variable specified
      if (outputVar) {
        context.setVariable(String(outputVar), extractedData);
      }

      return this.success(
        {
          selector,
          extractType,
          data: extractedData,
          count: multiple ? (extractedData as any[]).length : 1
        },
        {
          extractTime,
          multiple
        }
      );
    } catch (error) {
      return this.failure(
        `Failed to extract data: ${error instanceof Error ? error.message : String(error)}`
      );
    }
  }

  /**
   * Helper method to extract data from an element
   */
  private async extractFromElement(
    element: any,
    extractType: string,
    attribute?: any
  ): Promise<any> {
    switch (extractType) {
      case 'textContent':
        return await element.textContent();

      case 'innerText':
        return await element.innerText();

      case 'innerHTML':
        return await element.innerHTML();

      case 'attribute':
        if (!attribute) {
          throw new Error('Attribute name is required when extractType is "attribute"');
        }
        return await element.getAttribute(String(attribute));

      case 'value':
        return await element.inputValue();

      default:
        throw new Error(`Unknown extract type: ${extractType}`);
    }
  }

  validate(node: Node, parameters: Parameter[]): { valid: boolean; error?: string } {
    const validation = this.validateRequiredParameters(parameters, ['selector']);
    if (!validation.valid) {
      return validation;
    }

    // If extractType is 'attribute', attribute parameter must be provided
    const extractType = this.getParameter(parameters, 'extractType');
    if (extractType === 'attribute') {
      const attribute = this.getParameter(parameters, 'attribute');
      if (!attribute) {
        return {
          valid: false,
          error: 'Attribute parameter is required when extractType is "attribute"'
        };
      }
    }

    return { valid: true };
  }
}
