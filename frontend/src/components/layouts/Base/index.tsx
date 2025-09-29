import { Outlet } from "react-router";

import MobileNavBar from "../MobileNavBar";
import NavBar from "../NavBar";
import SideBar from "../SideBar";
import { useUser } from "../../../hooks/useUser";
import useMobile from "../../../lib/useMobile";
import style from "./Base.module.css";

function Base() {
  const { user } = useUser();
  const isMobile = useMobile();

  return (
    <div className={style.container}>
      {isMobile ? <MobileNavBar /> : <NavBar />}
      <div className={style.content}>
        <div className={style.spacer}>
          <Outlet />
        </div>
      </div>
      {!isMobile && user && <SideBar />}
    </div>
  );
}

export default Base;
