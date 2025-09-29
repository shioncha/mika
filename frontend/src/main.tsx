import "./styles/globals.css";
import "./styles/variables.css";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router";

import { AuthProvider } from "./hooks/authContext";
import { UserProvider } from "./hooks/userContext";
import AppRoutes from "./routes/index.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <AuthProvider>
        <UserProvider>
          <AppRoutes />
        </UserProvider>
      </AuthProvider>
    </BrowserRouter>
  </StrictMode>
);
