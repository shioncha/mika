import { Outlet } from "react-router";

import NavMenu from "../components/layouts/NavMenu";
import styles from "../styles/layouts/Base.module.css";

function Base() {
  return (
    <div className={styles.container}>
      <Outlet />
      <NavMenu />
    </div>
  )
}

export default Base;
