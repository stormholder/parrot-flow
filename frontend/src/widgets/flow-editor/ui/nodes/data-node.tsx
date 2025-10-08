import type { CustomNodeProps } from "../../types";
import BaseNode from "./base-node";

const DataNode = ({ data }: CustomNodeProps) => {
  const { label, title, icon, color } = data;
  return (
    <BaseNode
      label={label}
      title={title}
      icon={icon}
      color={color}
      leftHandle={false}
    />
  );
};

export default DataNode;
