import { scenario } from "@/../mock/scenario.mock";
import FlowEditor from "@/widgets/flow-editor/ui/flow/flow-editor";

const ScenarioFlowPage = () => {
  return (
    <div className="w-full relative overflow-x-hidden">
      <FlowEditor
        nodes={scenario.blocks}
        edges={scenario.edges}
        showAside={false}
      />
    </div>
  );
};

export default ScenarioFlowPage;
