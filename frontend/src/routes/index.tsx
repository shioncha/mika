import { Route, Routes, useLocation } from "react-router";

import Modal from "../components/elements/modal";
import TimelineLayout from "../components/layouts/timeline"
import SettingsPage from "../pages/settings";
import SignInPage from "../pages/signin";

function AppRoutes() {
  const location = useLocation();
  const background = location.state?.background;

  const mainRoutes = (
    <Routes location={background || location}>
      <Route path="/settings" element={<SettingsPage />} />
      <Route path="/signin" element={<SignInPage />} />
      <Route element={<TimelineLayout />}>
        <Route path="/" element={null} />
        <Route path="/tags/:tag" element={null} />
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
