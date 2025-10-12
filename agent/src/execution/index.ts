/**
 * Execution Engine
 *
 * Main entry point for the scenario execution engine.
 * Exports all necessary components for executing browser automation scenarios.
 */

// Main orchestrator
export { ScenarioExecutor } from './ScenarioExecutor.js';
export type { ScenarioExecutorOptions } from './ScenarioExecutor.js';

// Context
export { ExecutionContext } from './context/ExecutionContext.js';
export type { ExecutionContextConfig } from './context/ExecutionContext.js';

// Factory
export { NodeExecutorFactory } from './factory/NodeExecutorFactory.js';

// Base interfaces
export type { INodeExecutor, NodeExecutionResult } from './executors/base/INodeExecutor.js';
export { BaseNodeExecutor } from './executors/base/BaseNodeExecutor.js';

// Individual executors (for custom usage or testing)
export { StartNodeExecutor } from './executors/StartNodeExecutor.js';
export { GotoNodeExecutor } from './executors/GotoNodeExecutor.js';
export { WaitDurationNodeExecutor } from './executors/WaitDurationNodeExecutor.js';
export { FindElementNodeExecutor } from './executors/FindElementNodeExecutor.js';
export { ClickNodeExecutor } from './executors/ClickNodeExecutor.js';
export { InputDataNodeExecutor } from './executors/InputDataNodeExecutor.js';
export { KeyPressNodeExecutor } from './executors/KeyPressNodeExecutor.js';
export { ScreenshotNodeExecutor } from './executors/ScreenshotNodeExecutor.js';
export { LoadDataNodeExecutor } from './executors/LoadDataNodeExecutor.js';
export { ExtractDataNodeExecutor } from './executors/ExtractDataNodeExecutor.js';
