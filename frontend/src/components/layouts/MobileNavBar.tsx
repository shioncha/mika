import { useEffect, useState } from "react";
import type { IconType } from "react-icons";
import { CiCirclePlus, CiHome, CiLock, CiSettings } from "react-icons/ci";
import { IoIosAdd } from "react-icons/io";
import { Link, useLocation } from "react-router";

import { useUser } from "../../hooks/useUser";
import style from "../../styles/components/layouts/MobileNavBar.module.css";

interface NavItem {
  label: string;
  icon: IconType;
  link: string;
}

const signedNavItems: NavItem[] = [
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

const guestNavItems: NavItem[] = [
  {
    label: "Top",
    icon: CiHome,
    link: "/",
  },
  {
    label: "Sign In",
    icon: CiLock,
    link: "/signin",
  },
  {
    label: "Sign Up",
    icon: CiCirclePlus,
    link: "/signup",
  },
];

function MobileNavBar() {
  const { user } = useUser();
  const [active, setActive] = useState<string>("");
  const location = useLocation();

  const navItems = user ? signedNavItems : guestNavItems;

  useEffect(() => {
    setActive(location.pathname);
  }, [location]);

  return (
    <div className={style.container}>
      <ul className={style.nav}>
        {navItems.map((item) => (
          <li key={item.label}>
            <Link
              to={item.link}
              className={`${style.link} ${
                active == item.link ? style.active : undefined
              }`}
            >
              <item.icon size={24} />
              <span className={style.label}>{item.label}</span>
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default MobileNavBar;
