import { Outlet } from "react-router";

import useMobile from "../../../hooks/useMobile";
import { useUser } from "../../../hooks/useUser";
import MobileNavBar from "../MobileNavBar";
import NavBar from "../NavBar";
import SideBar from "../SideBar";
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
