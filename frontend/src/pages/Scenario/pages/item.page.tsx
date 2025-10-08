import FlowView from "@/widgets/flow-editor/ui/FlowView";

import { scenario } from "@/../mock/scenario.mock";

const ScenarioItemPage = () => {
  return (
    <div className="container pt-8 mb-28">
      <h1 className="text-2xl line-clamp font-semibold px-8">Test scenario</h1>
      <div
        className="w-full overflow-hidden relative rounded mt-1 py-2 px-8"
        style={{ height: "450px" }}
      >
        <FlowView nodes={scenario.blocks} edges={scenario.edges} />
      </div>
    </div>
  );
};

export default ScenarioItemPage;
