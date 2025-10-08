import { RouterProvider } from "react-router-dom";
import "./index.css";
import { router } from "./app/router/router";
import { AppLayout } from "./shared/ui/layout/app";

function App() {
  return (
    <AppLayout>
      <RouterProvider router={router} />
    </AppLayout>
  );
}

export default App;
