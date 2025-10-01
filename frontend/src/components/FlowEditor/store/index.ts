import { create } from "zustand";
import type {
  Connection,
  Node,
  Edge,
  NodeChange,
  EdgeChange,
  HandleType,
} from "reactflow";
import {
  MarkerType,
  addEdge,
  updateEdge,
  applyNodeChanges,
  applyEdgeChanges,
} from "reactflow";
import type { NewEdgeProps, NewNodeProps, RFState } from "./types";
import type { NodeTypes } from "@shared/types/nodes";
import { nanoid } from "@/utils";
import { getNodeConfig } from "../model/nodes";

const initialNodes: Node[] = [];

const initialEdges: Edge[] = [];

const useStore = create<RFState>((set, get) => ({
  edgeUpdateSuccessful: true,
  nodes: initialNodes,
  edges: initialEdges,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  onSelectionChange: (_params: { nodes: Node[]; edges: Edge[] }) => {
    // console.log(params);
  },
  onNodesChange: (changes: NodeChange[]) => {
    set({
      nodes: applyNodeChanges(changes, get().nodes),
    });
  },
  onEdgesChange: (changes: EdgeChange[]) => {
    set({
      edges: applyEdgeChanges(changes, get().edges),
    });
  },
  onConnect: (connection: Connection) => {
    if (!connection.source || !connection.target) {
      return;
    }
    const newEdge: Edge = {
      id: nanoid(),
      source: connection.source,
      sourceHandle: connection.sourceHandle,
      target: connection.target,
      targetHandle: connection.targetHandle,
      className: "base-node-edge",
      style: { strokeWidth: 2 },
      markerEnd: {
        type: MarkerType.ArrowClosed,
      },
    };
    set({
      edges: addEdge(newEdge, get().edges),
    });
  },
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  onEdgeUpdateStart: (
    _event: React.MouseEvent,
    _edge: Edge,
    _handle: HandleType
  ) => {
    get().edgeUpdateSuccessful = false;
  },
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  onEdgeUpdateEnd: (
    _event: React.MouseEvent,
    edge: Edge,
    _handle: HandleType
  ) => {
    if (!get().edgeUpdateSuccessful) {
      set({
        edges: get().edges.filter((e: Edge) => e.id !== edge.id),
      });
    }
    set({
      edgeUpdateSuccessful: true,
    });
  },
  onEdgeUpdate: (oldEdge: Edge, newConnection: Connection) => {
    set({
      edgeUpdateSuccessful: true,
    });
    set({
      edges: updateEdge(oldEdge, newConnection, get().edges),
    });
  },
  appendNode: (nodeProps: NewNodeProps) => {
    const newNode: Node = {
      id: nodeProps.id,
      type: nodeProps.type,
      position: nodeProps.position,
      data: getNodeConfig(
        nodeProps.type as NodeTypes,
        nodeProps.id,
        nodeProps.data
      ),
    };
    set({
      nodes: get().nodes.concat(newNode),
    });
  },
  appendEdge: (edgeProps: NewEdgeProps) => {
    const existingEdge = get().edges.find((e) => e.id == edgeProps.id);
    if (existingEdge) {
      return;
    }
    const newEdge: Edge = {
      id: edgeProps.id,
      source: edgeProps.source,
      target: edgeProps.target,
      sourceHandle: edgeProps.sourceHandle,
      targetHandle: edgeProps.targetHandle,
      className: "base-node-edge",
      style: { strokeWidth: 2 },
      markerEnd: {
        type: MarkerType.ArrowClosed,
      },
    };
    set({
      edges: get().edges.concat(newEdge),
    });
  },
  clear: () => {
    set({
      nodes: [],
      edges: [],
    });
  },
}));

export default useStore;
