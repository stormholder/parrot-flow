import type {
  Connection,
  Node,
  Edge,
  OnNodesChange,
  OnEdgesChange,
  OnConnect,
  HandleType,
} from "reactflow";
import type { Edge as ScenarioBlockEdge } from "@/api-client/models/Edge";

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
  nodes: Node[];
  edges: Edge[];
  onSelectionChange: (params: { nodes: Node[]; edges: Edge[] }) => void;
  onNodesChange: OnNodesChange;
  onEdgesChange: OnEdgesChange;
  onEdgeUpdateStart: (
    event: React.MouseEvent,
    edge: Edge,
    handleType: HandleType
  ) => void;
  onEdgeUpdate: (oldEdge: Edge, newConnection: Connection) => void;
  onEdgeUpdateEnd: (
    event: React.MouseEvent,
    edge: Edge,
    handleType: HandleType
  ) => void;
  onConnect: OnConnect;
  appendNode: (nodeProps: NewNodeProps) => void;
  appendEdge: (edgeProps: NewEdgeProps) => void;
  clear: () => void;
};
