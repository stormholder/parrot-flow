import { TbPlayerPlayFilled } from "react-icons/tb";
import BaseNode from "../BaseNode";
import { nodeColors } from "../../model/nodes";

const StartNode = () => {
  return (
    <BaseNode
      label="trigger"
      title="start"
      icon={TbPlayerPlayFilled}
      color={nodeColors.trigger}
      leftHandle={false}
    />
  );
};

export default StartNode;
