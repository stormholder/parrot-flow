/**
 * Example usage of the ScenarioGraph traversal
 *
 * This demonstrates how to use the graph traversal for a realistic scenario
 */

import { ScenarioGraph } from '../../src/graph/traversal.js';
import { Context } from '../../src/types/generated/messages.js';

/**
 * Example: Login flow scenario
 *
 * Flow:
 * START -> GOTO (login page) -> FIND (username field) -> INPUT (username)
 *       -> FIND (password field) -> INPUT (password) -> CLICK (submit)
 *       -> WAIT (for dashboard) -> SCREENSHOT -> END
 */
function createLoginFlowExample(): Context {
  return {
    blocks: [
      { id: 'start', node_type: 'start', position: { x: 100, y: 100 } },
      { id: 'goto_login', node_type: 'goto', position: { x: 100, y: 200 } },
      { id: 'find_username', node_type: 'findelement', position: { x: 100, y: 300 } },
      { id: 'input_username', node_type: 'inputdata', position: { x: 100, y: 400 } },
      { id: 'find_password', node_type: 'findelement', position: { x: 100, y: 500 } },
      { id: 'input_password', node_type: 'inputdata', position: { x: 100, y: 600 } },
      { id: 'click_submit', node_type: 'click', position: { x: 100, y: 700 } },
      { id: 'wait_dashboard', node_type: 'waitduration', position: { x: 100, y: 800 } },
      { id: 'take_screenshot', node_type: 'screenshot', position: { x: 100, y: 900 } }
    ],
    edges: [
      { id: 'e1', source: 'start', target: 'goto_login' },
      { id: 'e2', source: 'goto_login', target: 'find_username' },
      { id: 'e3', source: 'find_username', target: 'input_username' },
      { id: 'e4', source: 'input_username', target: 'find_password' },
      { id: 'e5', source: 'find_password', target: 'input_password' },
      { id: 'e6', source: 'input_password', target: 'click_submit' },
      { id: 'e7', source: 'click_submit', target: 'wait_dashboard' },
      { id: 'e8', source: 'wait_dashboard', target: 'take_screenshot' }
    ]
  };
}

/**
 * Example: Conditional flow with branches
 *
 * Flow:
 * START -> GOTO -> FIND (element)
 *       -> [success] CLICK
 *       -> [error] SCREENSHOT (error) -> END
 * CLICK -> WAIT -> SCREENSHOT (success) -> END
 */
function createConditionalFlowExample(): Context {
  return {
    blocks: [
      { id: 'start', node_type: 'start', position: { x: 100, y: 100 } },
      { id: 'goto_page', node_type: 'goto', position: { x: 100, y: 200 } },
      { id: 'find_element', node_type: 'findelement', position: { x: 100, y: 300 } },
      { id: 'click_element', node_type: 'click', position: { x: 50, y: 400 } },
      { id: 'error_screenshot', node_type: 'screenshot', position: { x: 200, y: 400 } },
      { id: 'wait_success', node_type: 'waitduration', position: { x: 50, y: 500 } },
      { id: 'success_screenshot', node_type: 'screenshot', position: { x: 50, y: 600 } }
    ],
    edges: [
      { id: 'e1', source: 'start', target: 'goto_page' },
      { id: 'e2', source: 'goto_page', target: 'find_element' },
      { id: 'e3', source: 'find_element', target: 'click_element', source_handle: 'success' },
      { id: 'e4', source: 'find_element', target: 'error_screenshot', source_handle: 'error' },
      { id: 'e5', source: 'click_element', target: 'wait_success' },
      { id: 'e6', source: 'wait_success', target: 'success_screenshot' }
    ]
  };
}

/**
 * Example: Parallel data extraction
 *
 * Flow:
 * START -> GOTO -> EXTRACT_TITLE
 *               -> EXTRACT_PRICE  -> COMBINE -> END
 *               -> EXTRACT_DESCRIPTION
 */
function createParallelExtractionExample(): Context {
  return {
    blocks: [
      { id: 'start', node_type: 'start', position: { x: 100, y: 100 } },
      { id: 'goto_product', node_type: 'goto', position: { x: 100, y: 200 } },
      { id: 'extract_title', node_type: 'extractdata', position: { x: 50, y: 300 } },
      { id: 'extract_price', node_type: 'extractdata', position: { x: 100, y: 300 } },
      { id: 'extract_desc', node_type: 'extractdata', position: { x: 150, y: 300 } },
      { id: 'combine_data', node_type: 'loaddata', position: { x: 100, y: 400 } }
    ],
    edges: [
      { id: 'e1', source: 'start', target: 'goto_product' },
      { id: 'e2', source: 'goto_product', target: 'extract_title' },
      { id: 'e3', source: 'goto_product', target: 'extract_price' },
      { id: 'e4', source: 'goto_product', target: 'extract_desc' },
      { id: 'e5', source: 'extract_title', target: 'combine_data' },
      { id: 'e6', source: 'extract_price', target: 'combine_data' },
      { id: 'e7', source: 'extract_desc', target: 'combine_data' }
    ]
  };
}

// Demo usage
function demonstrateTraversal() {
  console.log('=== Context Traversal Examples ===\n');

  // Example 1: Linear Login Flow
  console.log('1. Linear Login Flow:');
  const loginContext = new ScenarioGraph(createLoginFlowExample());
  const loginValidation = loginContext.validate();
  console.log('   Valid:', loginValidation.valid);

  if (loginValidation.valid) {
    const loginOrder = loginContext.topologicalSort();
    console.log('   Execution order:', loginOrder.join(' → '));
    console.log('   Start nodes:', loginContext.getStartNodes());
    console.log('   End nodes:', loginContext.getEndNodes());
  }
  console.log();

  // Example 2: Conditional Flow
  console.log('2. Conditional Flow with Branches:');
  const conditionalContext = new ScenarioGraph(createConditionalFlowExample());
  const conditionalValidation = conditionalContext.validate();
  console.log('   Valid:', conditionalValidation.valid);

  if (conditionalValidation.valid) {
    const conditionalOrder = conditionalContext.topologicalSort();
    console.log('   Execution order:', conditionalOrder.join(' → '));
    console.log('   Branch marks for click_element:', conditionalContext.getBranchMarks('click_element'));
    console.log('   Branch marks for error_screenshot:', conditionalContext.getBranchMarks('error_screenshot'));
  }
  console.log();

  // Example 3: Parallel Extraction
  console.log('3. Parallel Data Extraction:');
  const parallelContext = new ScenarioGraph(createParallelExtractionExample());
  const parallelValidation = parallelContext.validate();
  console.log('   Valid:', parallelValidation.valid);

  if (parallelValidation.valid) {
    const parallelOrder = parallelContext.topologicalSort();
    console.log('   Execution order:', parallelOrder.join(' → '));
    console.log('   Predecessors of combine_data:', parallelContext.getPredecessors('combine_data'));
    console.log('   Note: extract_* blocks can run in parallel!');
  }
  console.log();

  // Example 4: Invalid Context (cycle)
  console.log('4. Invalid Context (contains cycle):');
  const cyclicContext: Context = {
    blocks: [
      { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
      { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
      { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
    ],
    edges: [
      { id: 'e1', source: 'A', target: 'B' },
      { id: 'e2', source: 'B', target: 'C' },
      { id: 'e3', source: 'C', target: 'B' } // Creates cycle!
    ]
  };
  const invalidContext = new ScenarioGraph(cyclicContext);
  const invalidValidation = invalidContext.validate();
  console.log('   Valid:', invalidValidation.valid);
  console.log('   Errors:', invalidValidation.errors);
}

// Run the demo if this file is executed directly
if (import.meta.url === `file://${process.argv[1]}`) {
  demonstrateTraversal();
}

export {
  createLoginFlowExample,
  createConditionalFlowExample,
  createParallelExtractionExample,
  demonstrateTraversal
};
