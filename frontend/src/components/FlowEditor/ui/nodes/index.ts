import type { NodeTypes } from "@shared/types/nodes";
import StartNode from "./StartNode";
import type { JSX } from "react";
import InNode from "./InNode";
import DataNode from "./DataNode";
import type { CustomNodeProps } from "../../model";

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
