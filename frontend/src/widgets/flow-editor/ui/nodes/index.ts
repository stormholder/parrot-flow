import type { NodeTypes } from "@shared/types/nodes";
import StartNode from "./start-node";
import type { JSX } from "react";
import InNode from "./in-node";
import DataNode from "./data-node";
import type { CustomNodeProps } from "../../types";

const nodeTypes: Record<
  NodeTypes | "start",
  (props: CustomNodeProps) => JSX.Element
> = {
  start: StartNode,
  goto: InNode,
  inputdata: InNode,
  click: InNode,
  keypress: InNode,
  findelement: InNode,
  loaddata: DataNode,
  waitduration: InNode,
  screenshot: InNode,
};

export default nodeTypes;
