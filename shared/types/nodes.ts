export type NodeTypes =
  | "goto"
  | "inputdata"
  | "click"
  | "keypress"
  | "findelement"
  | "loaddata"
  | "waitduration"
  | "screenshot";

export type DraggableNodeLabel = 'api' | 'interaction' | 'data' | 'visual' | 'function' | 'logic' | 'trigger';

export interface NodeIO {
  id: string,
  label: string,
  type: string,
  default?: any,
  values?: any[]
}

export type BaseDraggableNode = {
  label: DraggableNodeLabel
  title: string
  color?: string
}