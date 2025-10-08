import { useCallback, useId } from "react";
import { useReactFlow } from "reactflow";
import { DndContext, type DragEndEvent } from "@dnd-kit/core";
import { nanoid } from "@/utils";
import { getNodeConfig } from "../../model/nodes";
import CustomDragOverlay from "./custom-drag-overlay";

interface CustomDndContextProps {
  children: React.ReactNode;
}

const CustomDndContext = ({ children }: CustomDndContextProps) => {
  const ctxId = useId();
  const { setNodes, screenToFlowPosition } = useReactFlow();

  const onDragEnd = useCallback(
    ({ active }: DragEndEvent) => {
      const type = active.data.current?.type;
      const data = active.data.current?.data;
      if (!type) return;

      const el = document.getElementById("overlay-id")!;
      const clientRect = el.getBoundingClientRect();

      const position = screenToFlowPosition({
        x: clientRect.left,
        y: clientRect.top,
      });
      const ID = nanoid();
      const newNode = {
        id: ID,
        type,
        position,
        data: getNodeConfig(type, ID, data),
      };

      setNodes((nds) => nds.concat(newNode));
    },
    [screenToFlowPosition, setNodes]
  );

  return (
    <DndContext id={ctxId} onDragEnd={onDragEnd}>
      <CustomDragOverlay />
      {children}
    </DndContext>
  );
};

export default CustomDndContext;
