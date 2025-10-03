import { RootPage } from "@/pages/RootPage/root.page";
import { createBrowserRouter } from "react-router-dom";

import ErrorPage from "@/pages/ErrorPage/error.page";
import DashboardPage from "@/pages/Dashboard/dashboard.page";
import { ScenarioItemLoader, ScenarioListLoader } from "@/pages/Scenario/query";
import ScenarioListPage from "@/pages/Scenario/pages/list.page";
import ScenarioItemPage from "@/pages/Scenario/pages/item.page";
import ScenarioFlowPage from "@/pages/Scenario/pages/flow.page";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootPage />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <DashboardPage />,
      },
      {
        path: "scenarios",
        loader: ScenarioListLoader,
        element: <ScenarioListPage />,
      },
      {
        path: "scenarios/:scenarioId",
        loader: ScenarioItemLoader,
        element: <ScenarioItemPage />,
      },
      {
        path: "scenarios/:scenarioId/flow",
        loader: ScenarioItemLoader,
        element: <ScenarioFlowPage />,
      },
    ],
  },
]);
