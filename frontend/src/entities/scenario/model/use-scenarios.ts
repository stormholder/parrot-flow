import { useState, useEffect } from "react";
import { scenarioApi } from "../api/scenario-api";
import type { ScenarioList, ScenarioFilters } from "./types";

export function useScenarios(filters?: ScenarioFilters) {
  const [scenarios, setScenarios] = useState<ScenarioList | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    const fetchScenarios = async () => {
      try {
        setLoading(true);
        const data = await scenarioApi.getList(filters);
        setScenarios(data);
        setError(null);
      } catch (err) {
        setError(
          err instanceof Error ? err : new Error("Failed to fetch scenarios")
        );
        setScenarios(null);
      } finally {
        setLoading(false);
      }
    };

    fetchScenarios();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [
    filters?.name,
    filters?.tags,
    filters?.page,
    filters?.rpp,
    filters?.order,
    filters?.dir,
  ]);

  return { scenarios, loading, error };
}

export async function scenariosLoader(): Promise<ScenarioList | null> {
  try {
    return await scenarioApi.getList();
  } catch (error) {
    console.error("Failed to load scenarios:", error);
    return null;
  }
}
