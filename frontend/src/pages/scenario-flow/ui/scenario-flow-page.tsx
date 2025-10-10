import { useLoaderData } from "react-router-dom";
import type { ScenarioItemLoaderData } from "@/entities/scenario";
import FlowEditor from "@/widgets/flow-editor/ui/flow/flow-editor";

const ScenarioFlowPage = () => {
  const { flow } = useLoaderData() as ScenarioItemLoaderData;

  return (
    <div className="w-full relative overflow-x-hidden">
      <FlowEditor
        nodes={flow.blocks}
        edges={flow.edges}
        showAside={false}
      />
    </div>
  );
};

export default ScenarioFlowPage;
