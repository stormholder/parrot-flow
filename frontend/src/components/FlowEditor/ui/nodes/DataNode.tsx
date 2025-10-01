import type { CustomNodeProps } from "../../model";
import BaseNode from "../BaseNode";

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
