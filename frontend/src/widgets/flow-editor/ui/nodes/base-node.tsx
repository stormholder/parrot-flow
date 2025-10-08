import { useMemo } from "react";
import { Handle, Position, useEdges } from "reactflow";
import type {
  BaseNodeBodyProps,
  BaseNodeBodySectionProps,
  BaseNodeParamProps,
  BaseNodeProps,
} from "../../model";
import { clsx } from "@/shared/lib/utils";

const BaseNode = ({
  label,
  title,
  icon: Icon,
  color,
  children,
  leftHandle = true,
  rightHandle = true,
}: Readonly<BaseNodeProps>) => {
  const style = color
    ? {
        backgroundColor: color,
      }
    : {};
  return (
    <div
      className={clsx(
        "base-node",
        "min-w-40 rounded-lg shadow-md bg-white dark:bg-neutral-800"
      )}
    >
      <div className="relative flex items-center justify-between p-2">
        {leftHandle ? (
          <Handle
            type="target"
            id="node-source"
            position={Position.Left}
            className="!-left-2 text-black dark:text-white"
          />
        ) : null}
        <div className="flex items-center space-x-2">
          <div
            className="flex h-10 w-10 items-center justify-center rounded-md bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-white"
            style={style}
          >
            <Icon size={20} />
          </div>
          <div className="flex flex-col pr-1 capitalize">
            <p className="text-[10px] opacity-50">{label}</p>
            <p className="text-xs font-bold">{title}</p>
          </div>
        </div>
        {rightHandle ? (
          <Handle
            type="source"
            id="node-target"
            position={Position.Right}
            className="!-right-2 text-black dark:text-white"
          />
        ) : null}
      </div>
      {children}
    </div>
  );
};

const BaseNodeBody = ({ children }: Readonly<BaseNodeBodyProps>) => (
  <div className="border-b-editor-node-border flex justify-between space=x=6 border-t px-2 py-3">
    {children}
  </div>
);

const BaseNodeBodySection = ({
  children,
}: Readonly<BaseNodeBodySectionProps>) => (
  <div className="min-w-1/2 flex flex-col space-y-3">{children}</div>
);

const BaseNodeParam = ({
  direction = "left",
  id,
  label,
  type,
}: BaseNodeParamProps) => {
  const edges = useEdges();

  const isConnected = useMemo(() => {
    const connection = edges.find((e) => e.id.includes(id));
    return !!connection;
  }, [edges, id]);

  const isLeft = direction === "left";

  return (
    <div className="relative flex h-6 items-center justify-between text-xs font-bold">
      <Handle
        type={isLeft ? "target" : "source"}
        id={id}
        position={isLeft ? Position.Left : Position.Right}
        className={clsx(
          "!h-3 !w-1 !min-w-0 !rounded-none !border-none",
          isLeft ? "!-left-4 !rounded-l-sm" : "!-right-4 !rounded-r-sm",
          isConnected ? "!bg-black" : "!bg-black/25"
        )}
      />
      <div
        className={clsx(
          "flex flex-1 items-center justify-between space-x-2 capitalize gap-1",
          isLeft ? "" : "flex-row-reverse space-x-reverse text-right"
        )}
      >
        <span
          className={clsx(
            "mr-1 flex min-h-4 min-w-4 items-center justify-center rounded font-bold uppercase",
            isConnected
              ? "bg-black text-white dark:bg-gray-400 dark:text-black"
              : "bg-gray-100 text-black/50"
          )}
        >
          {type}
        </span>
        <p
          className={clsx(
            "w-full leading-[10px]",
            isConnected
              ? " text-black dark:text-white"
              : "text-black/50 dark:text-white/50"
          )}
        >
          {label}
        </p>
      </div>
    </div>
  );
};

BaseNode.Body = BaseNodeBody;
BaseNode.Section = BaseNodeBodySection;
BaseNode.Param = BaseNodeParam;

export default BaseNode;
