import type { NodeTypes } from "@shared/types/nodes";
import type { DraggableNode, IONodeProps } from ".";
import { capitalize } from "@/shared/lib/utils";
import { nodes as toolboxNodes } from "./nodes";
import type { Node as RFNode } from "reactflow";

export const nodesConfig: Record<
  NodeTypes,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  (id: string, parameters?: any) => IONodeProps
> = {
  goto: (id) => ({
    inputs: [
      {
        id: id + "-url",
        label: "URL",
        type: "S",
      },
      {
        id: id + "-timeout",
        label: "Timeout",
        type: "N",
        default: 3000,
      },
    ],
  }),
  inputdata: (id) => ({
    inputs: [
      {
        id: id + "-selector",
        label: "Selector",
        type: "S",
      },
      {
        id: id + "-type",
        label: "Type",
        type: "S",
        default: "input",
        values: ["input", "select", "checkbox", "radio"],
      },
      {
        id: id + "-value",
        label: "Value",
        type: "S",
      },
    ],
  }),
  click: (id) => ({
    inputs: [
      {
        id: id + "-selector",
        label: "Selector",
        type: "S",
      },
    ],
  }),
  keypress: (id) => ({
    inputs: [
      {
        id: id + "-key",
        label: "Key",
        type: "S",
        values: [
          "Enter",
          "Space",
          "Tab",
          "Backspace",
          "PageUp",
          "PageDown",
          "Home",
          "End",
          "F1",
          "F2",
          "F3",
          "F4",
          "F5",
          "F6",
          "F7",
          "F8",
          "F9",
          "F10",
          "F11",
          "F12",
        ],
      },
    ],
  }),
  findelement: (id) => ({
    inputs: [
      {
        id: id + "-selector",
        label: "selector",
        type: "S",
      },
      {
        id: id + "-timeout",
        label: "timeout",
        type: "N",
      },
    ],
  }),
  loaddata: (id, parameters) => ({
    inputs:
      !!parameters && !!parameters.input
        ? // eslint-disable-next-line @typescript-eslint/no-explicit-any
          parameters.input.map((p: any) => {
            return {
              id: id + "-" + p.name,
              label: capitalize(p.name),
              type: p.type,
            };
          })
        : [],
    outputs: [],
  }),
  waitduration: (id) => ({
    inputs: [
      {
        id: id + "-duration",
        label: "Duration",
        type: "N",
        default: 3000,
      },
    ],
  }),
  screenshot: (id) => ({
    inputs: [
      {
        id: id + "-filename",
        label: "filename",
        type: "S",
      },
    ],
  }),
};

export const getNodeColor = (n: RFNode<DraggableNode>): string => {
  if (n.style?.backgroundColor) return n.style?.backgroundColor;
  if (!!n.type && n.type in toolboxNodes)
    return toolboxNodes[n.type as NodeTypes].color!;
  return "#e2e2e2";
};
