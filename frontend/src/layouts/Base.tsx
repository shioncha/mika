import { Outlet } from "react-router";

import MobileNavBar from "../components/layouts/MobileNavBar";
import NavBar from "../components/layouts/NavBar";
import useMobile from "../libs/useMobile";
import style from "../styles/layouts/Base.module.css";

function Base() {
  const isMobile = useMobile();

  return (
    <div className={style.container}>
      {isMobile ? <MobileNavBar /> : <NavBar />}
      <div className={style.content}>
        <div className={style.spacer}>
          <Outlet />
        </div>
      </div>
    </div>
  );
}

export default Base;
