import { Button } from "@heroui/react";
import { useDeleteScenario } from "../model";

interface DeleteScenarioButtonProps {
  scenarioId: number;
  onSuccess?: () => void;
  children?: React.ReactNode;
}

export const DeleteScenarioButton = ({ scenarioId, onSuccess, children }: DeleteScenarioButtonProps) => {
  const { deleteScenario, loading } = useDeleteScenario();

  const handleDelete = async () => {
    try {
      await deleteScenario(scenarioId);
      if (onSuccess) {
        onSuccess();
      }
    } catch (error) {
      // Error is handled by the hook
      console.error("Failed to delete scenario:", error);
    }
  };

  return (
    <Button onClick={handleDelete} isLoading={loading} color="danger">
      {children || "Delete"}
    </Button>
  );
};
