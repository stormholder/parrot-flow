/**
 * Graph traversal utilities for scenario execution
 *
 * This module provides topological sorting and graph traversal
 * for executing scenario nodes in the correct order.
 *
 * Aligned with backend domain: backend/internal/domain/scenario/value_objects.go
 */

import { Context } from '../types/generated/messages.js';

/**
 * Internal graph representation optimized for traversal
 */
export class ScenarioGraph {
  private vertices: Set<string>;
  private inputs: Map<string, string[]>;    // nodeId -> list of predecessor nodeIds
  private outputs: Map<string, string[]>;   // nodeId -> list of successor nodeIds
  private branchMarks: Map<string, string[]>; // nodeId -> list of source_handles for conditional branches
  private readonly conditions: Map<string, string>;   // edgeKey -> condition string

  constructor(context: Context) {
    this.vertices = new Set();
    this.inputs = new Map();
    this.outputs = new Map();
    this.branchMarks = new Map();
    this.conditions = new Map();

    // Extract all vertices from nodes (blocks)
    for (const node of context.blocks) {
      this.vertices.add(node.id);
      // Initialize empty arrays for each vertex
      this.inputs.set(node.id, []);
      this.outputs.set(node.id, []);
    }

    // Build edges
    for (const edge of context.edges) {
      this.addEdge(edge.source, edge.target, edge.source_handle, edge.condition);
    }
  }

  /**
   * Add an edge to the graph
   */
  private addEdge(source: string, target: string, handle?: string, condition?: string): void {
    // Add to outputs map
    const outputs = this.outputs.get(source) || [];
    outputs.push(target);
    this.outputs.set(source, outputs);

    // Add to inputs map
    const inputs = this.inputs.get(target) || [];
    inputs.push(source);
    this.inputs.set(target, inputs);

    // Track branch marks if handle is provided
    if (handle) {
      const marks = this.branchMarks.get(target) || [];
      marks.push(handle);
      this.branchMarks.set(target, marks);
    }

    // Track condition if provided
    if (condition) {
      const edgeKey = `${source}->${target}`;
      this.conditions.set(edgeKey, condition);
    }
  }

  /**
   * Perform topological sort using DFS
   * Returns execution order and detects cycles
   *
   * @returns Array of block IDs in execution order
   * @throws Error if graph contains cycles
   */
  public topologicalSort(): string[] {
    const order: string[] = [];
    const permanentMark = new Set<string>();
    const temporaryMark = new Set<string>();
    let isAcyclic = true;

    const visit = (nodeId: string): void => {
      if (temporaryMark.has(nodeId)) {
        // Found a cycle
        isAcyclic = false;
        return;
      }

      if (!permanentMark.has(nodeId) && !temporaryMark.has(nodeId)) {
        temporaryMark.add(nodeId);

        // Visit all predecessors (inputs)
        const predecessors = this.inputs.get(nodeId) || [];
        for (const predecessor of predecessors) {
          visit(predecessor);

          // Propagate branch marks
          const marks = this.branchMarks.get(predecessor);
          if (marks) {
            const currentMarks = this.branchMarks.get(nodeId) || [];
            this.branchMarks.set(nodeId, [...currentMarks, ...marks]);
          }

          if (!isAcyclic) {
            return;
          }
        }

        temporaryMark.delete(nodeId);
        permanentMark.add(nodeId);
        order.unshift(nodeId); // Add to beginning (reverse postorder)
      }
    };

    // Visit all vertices
    for (const nodeId of this.inputs.keys()) {
      if (!permanentMark.has(nodeId)) {
        visit(nodeId);
        if (!isAcyclic) {
          throw new Error('Graph contains cycles - not a valid DAG');
        }
      }
    }

    // Reverse the order to get correct execution sequence
    return order.reverse();
  }

  /**
   * Get all direct successors of a node (BFS step)
   *
   * @param nodeId - The node to get successors for
   * @param visited - Set of already visited nodes
   * @returns Array of unvisited successor node IDs
   */
  public getNextNodes(nodeId: string, visited: Set<string>): string[] {
    const successors = this.outputs.get(nodeId) || [];
    return successors.filter(successor => !visited.has(successor));
  }

  /**
   * Get all predecessors of a node
   */
  public getPredecessors(nodeId: string): string[] {
    return this.inputs.get(nodeId) || [];
  }

  /**
   * Get all successors of a node
   */
  public getSuccessors(nodeId: string): string[] {
    return this.outputs.get(nodeId) || [];
  }

  /**
   * Get branch marks for a node
   */
  public getBranchMarks(nodeId: string): string[] {
    return this.branchMarks.get(nodeId) || [];
  }

  /**
   * Get condition for an edge between two nodes
   */
  public getCondition(source: string, target: string): string | undefined {
    const edgeKey = `${source}->${target}`;
    return this.conditions.get(edgeKey);
  }

  /**
   * Check if a node has any predecessors
   */
  public isStartNode(nodeId: string): boolean {
    const predecessors = this.inputs.get(nodeId) || [];
    return predecessors.length === 0;
  }

  /**
   * Check if a node has any successors
   */
  public isEndNode(nodeId: string): boolean {
    const successors = this.outputs.get(nodeId) || [];
    return successors.length === 0;
  }

  /**
   * Get all start nodes (nodes with no predecessors)
   */
  public getStartNodes(): string[] {
    const startNodes: string[] = [];
    for (const nodeId of this.vertices) {
      if (this.isStartNode(nodeId)) {
        startNodes.push(nodeId);
      }
    }
    return startNodes;
  }

  /**
   * Get all end nodes (nodes with no successors)
   */
  public getEndNodes(): string[] {
    const endNodes: string[] = [];
    for (const nodeId of this.vertices) {
      if (this.isEndNode(nodeId)) {
        endNodes.push(nodeId);
      }
    }
    return endNodes;
  }

  /**
   * Validate the graph structure
   *
   * @returns Validation result with any errors found
   */
  public validate(): { valid: boolean; errors: string[] } {
    const errors: string[] = [];

    // Check for cycles
    try {
      this.topologicalSort();
    } catch (error) {
      errors.push(`Cycle detected: ${error instanceof Error ? error.message : String(error)}`);
    }

    // Check that all vertices have at least one connection
    for (const nodeId of this.vertices) {
      const hasInput = (this.inputs.get(nodeId) || []).length > 0;
      const hasOutput = (this.outputs.get(nodeId) || []).length > 0;

      if (!hasInput && !hasOutput) {
        errors.push(`Block ${nodeId} is isolated (no connections)`);
      }
    }

    // Check for multiple start nodes (warning, not error)
    const startNodes = this.getStartNodes();
    if (startNodes.length > 1) {
      errors.push(`Warning: Multiple start nodes found: ${startNodes.join(', ')}`);
    }

    // Check for no start nodes
    if (startNodes.length === 0) {
      errors.push('No start node found (all nodes have predecessors)');
    }

    return {
      valid: errors.length === 0,
      errors
    };
  }
}
