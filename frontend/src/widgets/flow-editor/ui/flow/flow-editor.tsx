import { useEffect, useState } from "react";
import type { FlowProps } from "../../model";
import { nodes as toolboxNodes } from "../../model/nodes";
import useStore from "../../store";
import ReactFlow, {
  Background,
  MiniMap,
  ReactFlowProvider,
  type Connection,
  type Edge,
  type Node,
  type HandleType,
  type NodeChange,
} from "reactflow";
import CustomDndContext from "../dnd/custom-dnd-context";
import { getNodeColor } from "../../model/configuration";
import nodeTypes from "../nodes";
import { isAppleOS } from "@/shared/lib/utils";

function FlowEditor(props: Readonly<FlowProps>) {
  const {
    nodes: flowNodes,
    edges: flowEdges,
    appendNode,
    appendEdge,
    clear: clearFlow,
    onConnect,
    onNodesChange,
    onSelectionChange,
    onEdgesChange,
    onEdgeUpdateStart,
    onEdgeUpdate,
    onEdgeUpdateEnd,
  } = useStore();

  const {
    nodes: propNodes,
    edges: propEdges,
    parameters,
    onScenarioChange,
    onNodeSelectionChange,
  } = props;

  const [hasChanges, setHasChanges] = useState<boolean>(false);

  useEffect(() => {
    clearFlow();
    propNodes.forEach((n) => appendNode(n));
    propEdges.forEach((e) => appendEdge(e));
    toolboxNodes.loaddata.parameters = parameters;
  }, []);

  useEffect(() => {
    if (hasChanges && !!onScenarioChange) {
      setHasChanges(false);
      onScenarioChange({
        nodes: flowNodes.map((fn) => {
          return {
            id: fn.id,
            position: fn.position,
            type: fn.type!,
          };
        }),
        edges: flowEdges.map((fe) => {
          return {
            id: fe.id,
            source: fe.source,
            target: fe.target,
            sourceHandle: fe.sourceHandle,
            targetHandle: fe.targetHandle,
          };
        }),
      });
    }
  }, [hasChanges]);

  const handleNodesChange = (changes: NodeChange[]) => {
    const changesCount = changes.filter(
      (c) => c.type === "add" || c.type === "remove" || c.type === "position"
    ).length;
    setHasChanges(changesCount > 0);
    onNodesChange(changes);
  };

  const handleConnect = (connection: Connection) => {
    if (!connection.source || !connection.target) {
      return;
    }
    setHasChanges(true);
    onConnect(connection);
  };

  const handleEdgeUpdateEnd = (
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    event: any,
    edge: Edge,
    handleType: HandleType
  ) => {
    setHasChanges(true);
    onEdgeUpdateEnd(event, edge, handleType);
  };

  const handleNodeSelectionChange = (change: {
    nodes: Node[];
    edges: Edge[];
  }) => {
    onSelectionChange(change);
    const id: string | undefined =
      change.nodes.length > 0 ? change.nodes[0].id : undefined;
    if (onNodeSelectionChange) onNodeSelectionChange(id);
  };

  return (
    <ReactFlowProvider>
      <CustomDndContext>
        <div className="h-full overflow-y-hidden">
          <div className="h-full w-full">
            {/* <AsideToolBox 
              nodes={toolboxNodes} 
              className={
                clsx(
                  'bg-white dark:bg-neutral-800',
                  'absolute top-0 left-0 transition-transform duration-300 ease-in-out', 
                  !props.showAside ? 'translate-x-[-15rem]' : ''
                )
              }
            /> */}
            <ReactFlow
              nodeTypes={nodeTypes}
              nodes={flowNodes}
              edges={flowEdges}
              onSelectionChange={handleNodeSelectionChange}
              onNodesChange={handleNodesChange}
              onEdgesChange={onEdgesChange}
              onEdgeUpdateStart={onEdgeUpdateStart}
              onEdgeUpdate={onEdgeUpdate}
              onEdgeUpdateEnd={handleEdgeUpdateEnd}
              onConnect={handleConnect}
              fitView
              proOptions={{ hideAttribution: true }}
              maxZoom={2}
              multiSelectionKeyCode={"DISABLED"}
              deleteKeyCode={isAppleOS() ? "Backspace" : "Delete"}
              snapToGrid={true}
            >
              <MiniMap
                className="bg-white dark:bg-neutral-900"
                nodeColor={getNodeColor}
              />
              <Background
                gap={10}
                className="bg-gray-100 dark:bg-neutral-900"
              />
            </ReactFlow>
          </div>
        </div>
      </CustomDndContext>
    </ReactFlowProvider>
  );
}

export default FlowEditor;
