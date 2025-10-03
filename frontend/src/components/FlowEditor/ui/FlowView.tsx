import ReactFlow, { MiniMap, Background } from "reactflow";
import "reactflow/dist/style.css";
import { nodes as toolboxNodes } from "../model/nodes";
import type { Edge, Node } from "@/api-client";
import type { Node as RFNode } from "reactflow";
import useStore from "../store";
import { useState } from "react";
import nodeTypes from "./nodes";
import type { NodeTypes } from "@shared/types/nodes";
import type { DraggableNode } from "../model";
export interface FlowViewProps {
  nodes: Node[];
  edges: Edge[];
}

const getNodeColor = (n: RFNode<DraggableNode>): string => {
  if (n.style?.backgroundColor) return n.style?.backgroundColor;
  if (!!n.type && n.type in toolboxNodes)
    return toolboxNodes[n.type as NodeTypes].color!;
  return "#e2e2e2";
};

function FlowView(props: Readonly<FlowViewProps>) {
  const [isInitialRender, setIsInitialRender] = useState<boolean>(true);
  const {
    nodes: flowNodes,
    edges: flowEdges,
    appendNode,
    appendEdge,
    clear: clearFlow,
  } = useStore();

  if (isInitialRender) {
    setIsInitialRender(false);
    clearFlow();
    props.nodes.forEach((n) => appendNode(n));
    props.edges.forEach((e) => appendEdge(e));
  }

  return (
    <ReactFlow
      nodeTypes={nodeTypes}
      nodes={flowNodes}
      edges={flowEdges}
      fitView
      proOptions={{ hideAttribution: true }}
      maxZoom={2}
      multiSelectionKeyCode={"DISABLED"}
      snapToGrid={true}
      elementsSelectable={false}
      nodesConnectable={false}
      nodesDraggable={false}
    >
      <MiniMap
        className="bg-white dark:bg-neutral-900"
        nodeColor={getNodeColor}
      />
      <Background gap={10} className="bg-gray-100 dark:bg-neutral-900" />
    </ReactFlow>
  );
}

export default FlowView;
