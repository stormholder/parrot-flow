import {
  TbDatabaseImport,
  TbEyeSearch,
  TbHourglassLow,
  TbKeyboard,
  TbPointer,
  TbScreenshot,
  TbWorldSearch,
  TbSquareArrowDown,
} from "react-icons/tb";

import type { DraggableNodeLabel, NodeTypes } from "@shared/types/nodes";
import type { DraggableNode } from ".";
import { nodesConfig } from "./configuration";

export const nodeColors: Record<DraggableNodeLabel, string> = {
  api: "rgb(239 197 118)",
  data: "rgb(224 255 153)",
  function: "rgb(229 231 235)",
  interaction: "rgb(161 255 185)",
  logic: "rgb(162 192 247)",
  trigger: "rgb(245 233 67)",
  visual: "rgb(202 155 243)",
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const nodes: Record<NodeTypes, DraggableNode & { parameters?: any }> = {
  goto: {
    label: "api",
    title: "Go To URL",
    icon: TbWorldSearch,
    color: nodeColors.api,
  },
  inputdata: {
    label: "interaction",
    title: "Input Data",
    icon: TbKeyboard,
    color: nodeColors.interaction,
  },
  click: {
    label: "interaction",
    title: "Click Element",
    icon: TbPointer,
    color: nodeColors.interaction,
  },
  keypress: {
    label: "interaction",
    title: "Press Keyboard key",
    icon: TbSquareArrowDown,
    color: nodeColors.interaction,
  },
  findelement: {
    label: "visual",
    title: "Find Element",
    icon: TbEyeSearch,
    color: nodeColors.visual,
  },
  loaddata: {
    label: "data",
    title: "Insert Data",
    icon: TbDatabaseImport,
    color: nodeColors.data,
  },
  waitduration: {
    label: "api",
    title: "Idle Wait",
    icon: TbHourglassLow,
    color: nodeColors.api,
  },
  screenshot: {
    label: "visual",
    title: "Take Screenshot",
    icon: TbScreenshot,
    color: nodeColors.visual,
  },
};

export const getNodeConfig = (
  type: NodeTypes,
  id: string,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  parameters?: any
) => {
  const configGenerator = nodesConfig[type];

  if (!configGenerator) return;

  return {
    ...nodes[type],
    ...configGenerator(id, parameters),
  };
};
