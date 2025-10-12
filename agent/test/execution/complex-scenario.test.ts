/**
 * Complex Scenario Execution Test
 *
 * Tests the execution engine with a more complex scenario:
 * 1. Start
 * 2. Navigate to a search engine
 * 3. Find search input
 * 4. Type search query
 * 5. Press Enter
 * 6. Wait for results
 * 7. Extract result titles
 * 8. Take screenshot
 */

import { test } from 'node:test';
import assert from 'node:assert';
import { ScenarioExecutor } from '../../src/execution/index.js';
import type { ExecuteScenarioMessage, ProgressEvent } from '../../src/types/generated/messages.js';

test('Execute complex scenario - search and extract', async (t) => {
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

  // Create a complex scenario message
  const message: ExecuteScenarioMessage = {
    run_id: 'test-run-002',
    scenario_id: 'test-scenario-002',
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
          id: 'find-1',
          node_type: 'findelement',
          position: { x: 200, y: 0 }
        },
        {
          id: 'input-1',
          node_type: 'inputdata',
          position: { x: 300, y: 0 }
        },
        {
          id: 'keypress-1',
          node_type: 'keypress',
          position: { x: 400, y: 0 }
        },
        {
          id: 'wait-1',
          node_type: 'waitduration',
          position: { x: 500, y: 0 }
        },
        {
          id: 'screenshot-1',
          node_type: 'screenshot',
          position: { x: 600, y: 0 }
        }
      ],
      edges: [
        { id: 'edge-1', source: 'start-1', target: 'goto-1' },
        { id: 'edge-2', source: 'goto-1', target: 'find-1' },
        { id: 'edge-3', source: 'find-1', target: 'input-1' },
        { id: 'edge-4', source: 'input-1', target: 'keypress-1' },
        { id: 'edge-5', source: 'keypress-1', target: 'wait-1' },
        { id: 'edge-6', source: 'wait-1', target: 'screenshot-1' }
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
            { name: 'url', value: 'https://duckduckgo.com' }
          ]
        },
        {
          block_id: 'find-1',
          input: [
            { name: 'selector', value: 'input[name="q"]' },
            { name: 'outputVariable', value: 'searchInput' }
          ]
        },
        {
          block_id: 'input-1',
          input: [
            { name: 'selector', value: 'input[name="q"]' },
            { name: 'text', value: 'playwright automation' },
            { name: 'clearFirst', value: true }
          ]
        },
        {
          block_id: 'keypress-1',
          input: [
            { name: 'key', value: 'Enter' }
          ]
        },
        {
          block_id: 'wait-1',
          input: [
            { name: 'duration', value: 2000 }
          ]
        },
        {
          block_id: 'screenshot-1',
          input: [
            { name: 'format', value: 'png' },
            { name: 'fullPage', value: false }
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
  assert.ok(eventTypes.includes('run_completed'), 'Should have run_completed event');

  // Check that all nodes were executed
  const nodeEvents = events.filter(e => e.node_id);
  const executedNodes = new Set(nodeEvents.map(e => e.node_id));

  assert.strictEqual(executedNodes.size, 7, 'Should execute all 7 nodes');

  console.log('\n✓ Complex scenario executed successfully!');
  console.log(`Total events: ${events.length}`);
  console.log(`Nodes executed: ${executedNodes.size}`);

  // Print execution times
  const completedEvents = events.filter(e => e.event === 'node_completed');
  console.log('\nNode execution times:');
  for (const event of completedEvents) {
    console.log(`  ${event.node_id}: ${event.execution_time_ms}ms`);
  }
});
