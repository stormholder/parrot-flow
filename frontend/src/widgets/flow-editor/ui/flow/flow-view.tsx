import ReactFlow, { MiniMap, Background } from "reactflow";
import "reactflow/dist/style.css";
import useStore from "../../store";
import { useState } from "react";
import nodeTypes from "../nodes";
import type { FlowViewProps } from "../../model";
import { getNodeColor } from "../../model/configuration";

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
