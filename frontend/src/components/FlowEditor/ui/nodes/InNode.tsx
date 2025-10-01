import type { CustomNodeProps } from "../../model";
import BaseNode from "../BaseNode";

const InNode = ({ data }: CustomNodeProps) => {
  const { label, title, icon, color, inputs } = data;

  return (
    <BaseNode label={label} title={title} icon={icon} color={color}>
      <BaseNode.Body>
        <BaseNode.Section>
          {inputs?.map((i) => (
            <BaseNode.Param
              key={i.id}
              id={i.id}
              label={i.label}
              type={i.type}
              direction="left"
            />
          ))}
        </BaseNode.Section>
        <BaseNode.Section></BaseNode.Section>
      </BaseNode.Body>
    </BaseNode>
  );
};

export default InNode;
