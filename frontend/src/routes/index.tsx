import { useContext } from "react";
import { Route, Routes, useLocation } from "react-router";

import Modal from "../components/elements/modal";
import TimelineLayout from "../components/layouts/timeline";
import { AuthContext } from "../hooks/auth_context";
import Base from "../layouts/Base";
import HomePage from "../pages/home";
import SettingsPage from "../pages/settings";
import SignInPage from "../pages/signin";
import SignUpPage from "../pages/signup";
import PrivateRoute from "./private";
import PublicRoute from "./public";

function AppRoutes() {
  const location = useLocation();
  const background = location.state?.background;

  const { isAuthenticated } = useContext(AuthContext);

  const mainRoutes = (
    <Routes location={background || location}>
      <Route element={<Base />}>
        <Route
          path="/"
          element={isAuthenticated ? <TimelineLayout /> : <HomePage />}
        />
        <Route element={<PrivateRoute />}>
          <Route path="/settings" element={<SettingsPage />} />
          <Route path="/tags/:tag" element={<TimelineLayout />} />
        </Route>
        <Route element={<PublicRoute />}>
          <Route path="/signin" element={<SignInPage />} />
          <Route path="/signup" element={<SignUpPage />} />
        </Route>
      </Route>
    </Routes>
  );

  const modalRoutes = background && (
    <Routes>
      <Route path="/posts/:id" element={<Modal />} />
    </Routes>
  );

  return (
    <>
      {mainRoutes}
      {modalRoutes}
    </>
  );
}

export default AppRoutes;
