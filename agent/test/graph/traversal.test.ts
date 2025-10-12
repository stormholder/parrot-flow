/**
 * Tests for graph traversal utilities
 */

import { describe, it } from 'node:test';
import assert from 'node:assert';
import { ScenarioGraph } from '../../src/graph/traversal.js';
import { Context } from '../../src/types/generated/messages.js';

describe('ScenarioGraph', () => {
  describe('topologicalSort', () => {
    it('should sort a simple linear graph', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'B', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const order = sg.topologicalSort();

      assert.deepStrictEqual(order, ['A', 'B', 'C']);
    });

    it('should sort a graph with parallel branches', () => {
      // A -> B -> D
      // A -> C -> D
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'D', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'A', target: 'C' },
          { id: 'e3', source: 'B', target: 'D' },
          { id: 'e4', source: 'C', target: 'D' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const order = sg.topologicalSort();

      // A should be first, D should be last
      assert.strictEqual(order[0], 'A');
      assert.strictEqual(order[3], 'D');
      // B and C should be in the middle (order doesn't matter)
      assert.ok(order.includes('B'));
      assert.ok(order.includes('C'));
    });

    it('should detect cycles in the graph', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'B', target: 'C' },
          { id: 'e3', source: 'C', target: 'A' } // Creates cycle
        ]
      };

      const sg = new ScenarioGraph(context);

      assert.throws(
        () => sg.topologicalSort(),
        /cycle/i
      );
    });

    it('should handle a single node graph', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } }
        ],
        edges: []
      };

      const sg = new ScenarioGraph(context);
      const order = sg.topologicalSort();

      assert.deepStrictEqual(order, ['A']);
    });
  });

  describe('getNextNodes', () => {
    it('should return unvisited successors', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'goto', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'A', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const visited = new Set<string>(['A']);
      const next = sg.getNextNodes('A', visited);

      assert.strictEqual(next.length, 2);
      assert.ok(next.includes('B'));
      assert.ok(next.includes('C'));
    });

    it('should filter out visited nodes', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'goto', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'A', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const visited = new Set<string>(['A', 'B']);
      const next = sg.getNextNodes('A', visited);

      assert.strictEqual(next.length, 1);
      assert.strictEqual(next[0], 'C');
    });
  });

  describe('getStartNodes', () => {
    it('should identify nodes with no predecessors', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'B', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const startNodes = sg.getStartNodes();

      assert.deepStrictEqual(startNodes, ['A']);
    });

    it('should identify multiple start nodes', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'C' },
          { id: 'e2', source: 'B', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const startNodes = sg.getStartNodes();

      assert.strictEqual(startNodes.length, 2);
      assert.ok(startNodes.includes('A'));
      assert.ok(startNodes.includes('B'));
    });
  });

  describe('getEndNodes', () => {
    it('should identify nodes with no successors', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'B', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const endNodes = sg.getEndNodes();

      assert.deepStrictEqual(endNodes, ['C']);
    });
  });

  describe('validate', () => {
    it('should validate a correct graph', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'B', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const result = sg.validate();

      assert.strictEqual(result.valid, true);
      assert.strictEqual(result.errors.length, 0);
    });

    it('should detect isolated nodes', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' }
          // C is isolated
        ]
      };

      const sg = new ScenarioGraph(context);
      const result = sg.validate();

      assert.strictEqual(result.valid, false);
      assert.ok(result.errors.some(e => e.includes('isolated')));
    });

    it('should detect cycles', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B' },
          { id: 'e2', source: 'B', target: 'A' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const result = sg.validate();

      assert.strictEqual(result.valid, false);
      assert.ok(result.errors.some(e => e.toLowerCase().includes('cycle')));
    });

    it('should warn about multiple start nodes', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'C', node_type: 'click', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'C' },
          { id: 'e2', source: 'B', target: 'C' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const result = sg.validate();

      assert.strictEqual(result.valid, false);
      assert.ok(result.errors.some(e => e.includes('Multiple start nodes')));
    });
  });

  describe('branch marks', () => {
    it('should track branch marks from edges', () => {
      const context: Context = {
        blocks: [
          { id: 'A', node_type: 'start', position: { x: 0, y: 0 } },
          { id: 'B', node_type: 'goto', position: { x: 0, y: 0 } }
        ],
        edges: [
          { id: 'e1', source: 'A', target: 'B', source_handle: 'success' }
        ]
      };

      const sg = new ScenarioGraph(context);
      const marks = sg.getBranchMarks('B');

      assert.deepStrictEqual(marks, ['success']);
    });
  });
});
