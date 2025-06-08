import { useContext } from "react";
import { Navigate,Outlet } from "react-router";

import { AuthContext } from "../hooks/auth_context";

function PublicRoute() {
  const { isAuthenticated } = useContext(AuthContext);

  if (isAuthenticated) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
}

export default PublicRoute;
