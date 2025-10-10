import { useLoaderData } from "react-router-dom";
import type { ScenarioItemLoaderData } from "@/entities/scenario";
import FlowView from "@/widgets/flow-editor/ui/flow/flow-view";

const ScenarioItemPage = () => {
  const { scenario, flow } = useLoaderData() as ScenarioItemLoaderData;

  return (
    <div className="container pt-8 mb-28">
      <h1 className="text-2xl line-clamp font-semibold px-8">{scenario.name}</h1>
      <p className="text-sm text-gray-600 px-8 mt-1">{scenario.description}</p>
      <div
        className="w-full overflow-hidden relative rounded mt-4 py-2 px-8"
        style={{ height: "450px" }}
      >
        <FlowView nodes={flow.blocks} edges={flow.edges} />
      </div>
    </div>
  );
};

export default ScenarioItemPage;
