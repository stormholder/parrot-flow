import { useState, useEffect } from "react";
import { scenarioApi } from "../api/scenario-api";
import type { Scenario } from "./types";

export function useScenario(id: number | string) {
  const [scenario, setScenario] = useState<Scenario | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    const fetchScenario = async () => {
      try {
        setLoading(true);
        const numericId = typeof id === "string" ? parseInt(id, 10) : id;
        const data = await scenarioApi.getById(numericId);
        setScenario(data);
        setError(null);
      } catch (err) {
        setError(
          err instanceof Error ? err : new Error("Failed to fetch scenario")
        );
        setScenario(null);
      } finally {
        setLoading(false);
      }
    };

    fetchScenario();
  }, [id]);

  return { scenario, loading, error };
}

export async function scenarioLoader({
  params,
}: {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  params: any;
}): Promise<Scenario | null> {
  try {
    if (!params.scenarioId) {
      return null;
    }
    const id = parseInt(params.scenarioId, 10);
    return await scenarioApi.getById(id);
  } catch (error) {
    console.error("Failed to load scenario:", error);
    return null;
  }
}
