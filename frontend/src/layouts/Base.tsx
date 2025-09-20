import { Outlet } from "react-router";

import MobileNavBar from "../components/layouts/MobileNavBar";
import NavBar from "../components/layouts/NavBar";
import useMobile from "../libs/useMobile";
import styles from "../styles/layouts/Base.module.css";

function Base() {
  const isMobile = useMobile();

  return (
    <div className={styles.container}>
      {isMobile ? <MobileNavBar /> : <NavBar />}
      <div className={styles.content}>
        <div className={styles.spacer}>
          <Outlet />
        </div>
      </div>
    </div>
  );
}

export default Base;
