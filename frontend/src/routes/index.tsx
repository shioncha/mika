import { Route, Routes, useLocation } from "react-router";

import Modal from "../components/elements/Modal";
import SettingsLayout from "../components/layouts/Settings";
import TimelineLayout from "../components/layouts/Timeline";
import { useAuth } from "../hooks/useAuth";
import Base from "../layouts/Base";
import HomePage from "../pages/Home";
import NewPage from "../pages/New";
import SettingsPage from "../pages/Settings";
import AboutPage from "../pages/Settings/About";
import EmailPage from "../pages/Settings/Email";
import PasswordPage from "../pages/Settings/Password";
import UsernamePage from "../pages/Settings/Username";
import SignInPage from "../pages/SignIn";
import SignUpPage from "../pages/SignUp";
import PrivateRoute from "./private";
import PublicRoute from "./public";

function AppRoutes() {
  const location = useLocation();
  const background = location.state?.background;

  const { isAuthenticated } = useAuth();

  const mainRoutes = (
    <Routes location={background || location}>
      <Route element={<Base />}>
        <Route
          path="/"
          element={isAuthenticated ? <TimelineLayout /> : <HomePage />}
        />
        <Route element={<PrivateRoute />}>
          <Route path="/new" element={<NewPage />} />
          <Route path="/settings" element={<SettingsPage />} />
          <Route path="/settings" element={<SettingsLayout />}>
            <Route path="name" element={<UsernamePage />} />
            <Route path="email" element={<EmailPage />} />
            <Route path="password" element={<PasswordPage />} />
            <Route path="about" element={<AboutPage />} />
          </Route>
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
