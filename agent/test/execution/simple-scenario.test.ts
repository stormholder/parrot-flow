/**
 * Simple Scenario Execution Test
 *
 * Tests the execution engine with a basic scenario:
 * 1. Start
 * 2. Navigate to a website
 * 3. Take a screenshot
 */

import { test } from 'node:test';
import assert from 'node:assert';
import { ScenarioExecutor } from '../../src/execution/index.js';
import type { ExecuteScenarioMessage, ProgressEvent } from '../../src/types/generated/messages.js';

test('Execute simple scenario - navigate and screenshot', async (t) => {
  // Track progress events
  const events: ProgressEvent[] = [];

  // Create executor
  const executor = new ScenarioExecutor({
    browserType: 'chromium',
    onProgress: async (event) => {
      events.push(event);
      console.log(`[${event.event}] ${event.node_id || 'run'} - ${event.error || 'OK'}`);
    }
  });

  // Create a simple scenario message
  const message: ExecuteScenarioMessage = {
    run_id: 'test-run-001',
    scenario_id: 'test-scenario-001',
    context: {
      blocks: [
        {
          id: 'start-1',
          node_type: 'start',
          position: { x: 0, y: 0 }
        },
        {
          id: 'goto-1',
          node_type: 'goto',
          position: { x: 100, y: 0 }
        },
        {
          id: 'screenshot-1',
          node_type: 'screenshot',
          position: { x: 200, y: 0 }
        }
      ],
      edges: [
        {
          id: 'edge-1',
          source: 'start-1',
          target: 'goto-1'
        },
        {
          id: 'edge-2',
          source: 'goto-1',
          target: 'screenshot-1'
        }
      ]
    },
    input_data: {
      parameters: [
        {
          block_id: 'start-1',
          input: []
        },
        {
          block_id: 'goto-1',
          input: [
            {
              name: 'url',
              value: 'https://example.com'
            }
          ]
        },
        {
          block_id: 'screenshot-1',
          input: [
            {
              name: 'format',
              value: 'png'
            },
            {
              name: 'fullPage',
              value: false
            }
          ]
        }
      ]
    },
    parameters: {
      input: [],
      output: []
    },
    browser_config: {
      headless: true,
      timeout: 30000
    }
  };

  // Execute the scenario
  await executor.execute(message);

  // Verify progress events
  assert.ok(events.length > 0, 'Should have progress events');

  // Check that we have the expected events
  const eventTypes = events.map(e => e.event);
  assert.ok(eventTypes.includes('run_started'), 'Should have run_started event');
  assert.ok(eventTypes.includes('node_started'), 'Should have node_started events');
  assert.ok(eventTypes.includes('node_completed'), 'Should have node_completed events');
  assert.ok(eventTypes.includes('run_completed'), 'Should have run_completed event');

  // Check that all nodes were executed
  const nodeEvents = events.filter(e => e.node_id);
  const executedNodes = new Set(nodeEvents.map(e => e.node_id));
  assert.ok(executedNodes.has('start-1'), 'Should execute start node');
  assert.ok(executedNodes.has('goto-1'), 'Should execute goto node');
  assert.ok(executedNodes.has('screenshot-1'), 'Should execute screenshot node');

  console.log('\n✓ Scenario executed successfully!');
  console.log(`Total events: ${events.length}`);
  console.log(`Nodes executed: ${executedNodes.size}`);
});

test('Verify factory supports all node types', async (t) => {
  const supportedTypes = ScenarioExecutor.getSupportedNodeTypes();

  console.log('Supported node types:', supportedTypes);

  assert.ok(supportedTypes.includes('start'), 'Should support start');
  assert.ok(supportedTypes.includes('goto'), 'Should support goto');
  assert.ok(supportedTypes.includes('waitduration'), 'Should support waitduration');
  assert.ok(supportedTypes.includes('findelement'), 'Should support findelement');
  assert.ok(supportedTypes.includes('click'), 'Should support click');
  assert.ok(supportedTypes.includes('inputdata'), 'Should support inputdata');
  assert.ok(supportedTypes.includes('keypress'), 'Should support keypress');
  assert.ok(supportedTypes.includes('screenshot'), 'Should support screenshot');
  assert.ok(supportedTypes.includes('loaddata'), 'Should support loaddata');
  assert.ok(supportedTypes.includes('extractdata'), 'Should support extractdata');

  assert.strictEqual(supportedTypes.length, 10, 'Should support exactly 10 node types');
});
