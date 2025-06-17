import { useContext, useEffect, useState } from "react";
import type { IconType } from "react-icons";
import { CiHome, CiSettings } from "react-icons/ci";
import { IoIosAdd } from "react-icons/io";
import { Link, useLocation } from "react-router";

import { UserContext } from "../../hooks/user_context";
import style from "../../styles/components/layouts/NavMenu.module.css";

interface NavItem {
  label: string;
  icon: IconType;
  link: string;
}

const navItems: NavItem[] = [
  {
    label: "Timeline",
    icon: CiHome,
    link: "/",
  },
  {
    label: "New",
    icon: IoIosAdd,
    link: "/new",
  },
  {
    label: "Settings",
    icon: CiSettings,
    link: "/settings",
  },
];

function NavMenu() {
  const { user } = useContext(UserContext);
  const [active, setActive] = useState<string>("");
  const location = useLocation();

  useEffect(() => {
    const path = location.pathname;
    switch (path) {
      case "/":
        setActive("Timeline");
        break;
      case "/new":
        setActive("New");
        break;
      case "/settings":
        setActive("Settings");
        break;
      default:
        break;
    }
  }, [location]);

  return (
    <div className={style.container}>
      <ul className={style.nav}>
        {user ? (
          <>
            {navItems.map((item) => (
              <li key={item.label}>
                <Link
                  to={item.link}
                  className={`${style.link} ${
                    active == item.label ? style.active : undefined
                  }`}
                >
                  <item.icon size={24} />
                  <span className={style.label}>{item.label}</span>
                </Link>
              </li>
            ))}
          </>
        ) : (
          <>
            <li>
              <Link to="/signin" className={style.link}>
                Sign in
              </Link>
            </li>
            <li>
              <Link to="/signup" className={style.link}>
                Sign up
              </Link>
            </li>
          </>
        )}
      </ul>
    </div>
  );
}

export default NavMenu;
