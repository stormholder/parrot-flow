import { useState, useCallback } from "react";
import { scenarioApi } from "@/entities/scenario";
import type { Node, Edge } from "@/shared/api-client";

export interface FlowData {
  nodes: Node[];
  edges: Edge[];
}

export function useSaveFlow(scenarioId: number) {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const [lastSaved, setLastSaved] = useState<Date | null>(null);

  const saveFlow = useCallback(async (flowData: FlowData) => {
    try {
      setLoading(true);
      setError(null);

      await scenarioApi.update(scenarioId, {
        context: {
          blocks: flowData.nodes,
          edges: flowData.edges,
        },
      });

      setLastSaved(new Date());
    } catch (err) {
      const errorObj = err instanceof Error ? err : new Error("Failed to save flow");
      setError(errorObj);
      throw errorObj;
    } finally {
      setLoading(false);
    }
  }, [scenarioId]);

  return {
    saveFlow,
    loading,
    error,
    lastSaved,
  };
}
