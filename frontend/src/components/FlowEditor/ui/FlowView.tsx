import ReactFlow, { MiniMap, Background } from "reactflow";
import "reactflow/dist/style.css";
import type { RFState } from "../store/types";
import { nodes as toolboxNodes } from "../model/nodes";
import type { Edge, Node } from "@/api-client";
import useStore from "../store";
import { useState } from "react";
import nodeTypes from "./nodes";

const selector = (state: RFState) => ({
  flowNodes: state.nodes,
  flowEdges: state.edges,
  appendNode: state.appendNode,
  appendEdge: state.appendEdge,
  clearFlow: state.clear,
});

export interface FlowViewProps {
  nodes: Node[];
  edges: Edge[];
}

export const FlowView: React.FC<FlowViewProps> = ({ nodes, edges }) => {
  const [isInitialRender, setIsInitialRender] = useState(true);
  const { flowNodes, flowEdges, appendNode, appendEdge, clearFlow } =
    useStore(selector);

  if (isInitialRender) {
    setIsInitialRender(false);
    clearFlow();
    nodes.forEach((n) => appendNode(n));
    edges.forEach((e) => appendEdge(e));
  }

  return (
    <ReactFlow
      nodes={flowNodes}
      edges={flowEdges}
      nodeTypes={nodeTypes}
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
        nodeColor={(n) => {
          if (n.style?.backgroundColor) return n.style?.backgroundColor;
          if (!!n.type && n.type in toolboxNodes)
            return (toolboxNodes as any)[n.type].color;
          return "#e2e2e2";
        }}
      />
      <Background gap={10} className="bg-gray-100 dark:bg-neutral-900" />
    </ReactFlow>
  );
};
