import { FaChevronLeft } from "react-icons/fa6";
import { Link, Outlet } from "react-router";

import style from "../../styles/components/layouts/Settings.module.css";

function SettingsLayout() {
  return (
    <>
      <Link className={style.backButton} to="/settings">
        <FaChevronLeft size={24} className={style.arrow} />
        <span>Back</span>
      </Link>
      <Outlet />
    </>
  );
}

export default SettingsLayout;
