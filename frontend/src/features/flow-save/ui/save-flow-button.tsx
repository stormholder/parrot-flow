import { Button } from "@heroui/react";
import { useSaveFlow, type FlowData } from "../model";

interface SaveFlowButtonProps {
  scenarioId: number;
  flowData: FlowData;
  onSuccess?: () => void;
  children?: React.ReactNode;
}

export const SaveFlowButton = ({ scenarioId, flowData, onSuccess, children }: SaveFlowButtonProps) => {
  const { saveFlow, loading, lastSaved } = useSaveFlow(scenarioId);

  const handleSave = async () => {
    try {
      await saveFlow(flowData);
      if (onSuccess) {
        onSuccess();
      }
    } catch (error) {
      // Error is handled by the hook
      console.error("Failed to save flow:", error);
    }
  };

  return (
    <div className="flex items-center gap-2">
      <Button onClick={handleSave} isLoading={loading} color="primary">
        {children || "Save Flow"}
      </Button>
      {lastSaved && (
        <span className="text-xs text-gray-500">
          Last saved: {lastSaved.toLocaleTimeString()}
        </span>
      )}
    </div>
  );
};
