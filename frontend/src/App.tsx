import { RouterProvider } from "react-router-dom";
import "./index.css";
import { router } from "./routes/router";
import { AppLayout } from "./layout/app";

function App() {
  return (
    <AppLayout>
      <RouterProvider router={router} />
    </AppLayout>
  );
}

export default App;
