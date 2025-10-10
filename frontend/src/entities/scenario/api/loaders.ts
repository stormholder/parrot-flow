import { ScenariosService } from "@/shared/api-client";
import type { ScenarioResponseItem, ListScenariosResponseBody, ErrorModel } from "@/shared/api-client";
import { getFlowByScenarioId } from "@/shared/mocks/data/flows";

interface FlowData {
  blocks: Array<{
    id: string;
    type: string;
    position: { x: number; y: number };
  }>;
  edges: Array<{
    id: string;
    source: string;
    target: string;
    sourceHandle: string;
    targetHandle: string;
  }>;
}

export interface ScenarioListLoaderData {
  scenarios: ListScenariosResponseBody;
  query: {
    name?: string;
    tag?: string;
  };
}

export interface ScenarioItemLoaderData {
  scenario: ScenarioResponseItem;
  flow: FlowData;
}

function isErrorModel(response: unknown): response is ErrorModel {
  return (
    typeof response === "object" &&
    response !== null &&
    "status" in response &&
    typeof (response as ErrorModel).status === "number"
  );
}

export async function scenarioListLoader({ request }: { request: Request }): Promise<ScenarioListLoaderData> {
  const url = new URL(request.url);
  const queryName = url.searchParams.get('name');
  const queryTag = url.searchParams.get('tag');
  const page = url.searchParams.get('page');
  const rpp = url.searchParams.get('rpp');

  try {
    const response = await ScenariosService.listScenarios(
      queryName || undefined,
      queryTag || undefined,
      page ? parseInt(page) : undefined,
      rpp ? parseInt(rpp) : undefined,
    );

    if (isErrorModel(response)) {
      throw new Error(response.detail || "Failed to load scenarios");
    }

    return {
      scenarios: response,
      query: {
        name: queryName || undefined,
        tag: queryTag || undefined,
      },
    };
  } catch (error) {
    console.error("Failed to load scenarios:", error);
    return {
      scenarios: {
        total: 0,
        data: [],
        page: 1,
        rpp: 10,
      },
      query: {
        name: queryName || undefined,
        tag: queryTag || undefined,
      },
    };
  }
}

export async function scenarioItemLoader({ params }: { params: { scenarioId: string } }): Promise<ScenarioItemLoaderData> {
  try {
    const response = await ScenariosService.getScenario(params.scenarioId);

    if (isErrorModel(response)) {
      throw new Error(response.detail || "Failed to load scenario");
    }

    // Get flow data - in production this would come from a separate API endpoint
    // For now, we use the mock data
    const flow = getFlowByScenarioId(params.scenarioId);

    return {
      scenario: response,
      flow,
    };
  } catch (error) {
    console.error("Failed to load scenario:", error);
    throw new Response("Scenario not found", { status: 404 });
  }
}
