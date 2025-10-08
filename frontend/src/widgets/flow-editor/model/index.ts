import type { Edge, Node } from "@/shared/api-client";
import type { BaseDraggableNode, NodeIO } from "@shared/types/nodes";
import type { IconType } from "react-icons";
import type { NodeProps } from "reactflow";

export type DraggableNode = BaseDraggableNode & { icon: IconType };

export interface IONodeProps {
  inputs?: NodeIO[];
  outputs?: NodeIO[];
}

export type CustomNodeProps = NodeProps<IONodeProps & DraggableNode>;

export interface BaseNodeProps {
  label: string;
  title: string;
  icon: IconType;
  color?: string;
  children?: React.ReactNode;
  leftHandle?: boolean;
  rightHandle?: boolean;
}

export interface BaseNodeBodyProps {
  children?: React.ReactNode;
}

export interface BaseNodeBodySectionProps {
  children?: React.ReactNode;
}

export interface BaseNodeParamProps {
  direction?: "left" | "right";
  params?: string[];
  id: string;
  label: string;
  type: string;
}

export interface FlowViewProps {
  nodes: Node[];
  edges: Edge[];
}

export interface FlowProps extends FlowViewProps {
  showAside: boolean;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  parameters?: any;
  onScenarioChange?: ({ nodes, edges }: FlowViewProps) => void;
  onNodeSelectionChange?: (nodeId: string | undefined) => void;
}
