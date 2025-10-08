import type { Edge as ScenarioBlockEdge } from "@/shared/api-client/models/Edge";
import type { BaseDraggableNode, NodeIO } from "@shared/types/nodes";
import type { IconType } from "react-icons";
import type {
  Connection,
  Edge as RFEdge,
  Node as RFNode,
  HandleType,
  OnConnect,
  OnEdgesChange,
  OnNodesChange,
} from "reactflow";

// Re-export React Flow types
export type {
  Node,
  Edge,
  Connection,
  NodeChange,
  EdgeChange,
  HandleType,
} from "reactflow";

// Store types
export type NewNodeProps = {
  id: string;
  type: string;
  position: {
    x: number;
    y: number;
  };
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  data?: any;
};

export type NewEdgeProps = ScenarioBlockEdge;

export type RFState = {
  edgeUpdateSuccessful: boolean;
  nodes: RFNode[];
  edges: RFEdge[];
  onSelectionChange: (params: { nodes: RFNode[]; edges: RFEdge[] }) => void;
  onNodesChange: OnNodesChange;
  onEdgesChange: OnEdgesChange;
  onEdgeUpdateStart: (
    event: React.MouseEvent,
    edge: RFEdge,
    handleType: HandleType
  ) => void;
  onEdgeUpdate: (oldEdge: RFEdge, newConnection: Connection) => void;
  onEdgeUpdateEnd: (
    event: React.MouseEvent,
    edge: RFEdge,
    handleType: HandleType
  ) => void;
  onConnect: OnConnect;
  appendNode: (nodeProps: NewNodeProps) => void;
  appendEdge: (edgeProps: NewEdgeProps) => void;
  clear: () => void;
};

// Node types
export type DraggableNode = BaseDraggableNode & { icon: IconType };

export interface IONodeProps {
  inputs?: NodeIO[];
  outputs?: NodeIO[];
}
