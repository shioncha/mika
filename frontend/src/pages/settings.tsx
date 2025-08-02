import { useContext } from "react";

import List from "../components/elements/List";
import ListElementSetting from "../components/elements/ListElementSetting";
import { useAuth } from "../hooks/auth_context";
import { UserContext } from "../hooks/user_context";
import style from "../styles/pages/settings.module.css";

function SettingsPage() {
  const { signOut } = useAuth();
  const { user } = useContext(UserContext);

  return (
    <>
      <p className={style.name}>Hello, {user?.name}!</p>
      <List className={style.list}>
        <ListElementSetting to="name" name="Name">
          {user?.name || "-"}
        </ListElementSetting>
        <ListElementSetting to="email" name="Email">
          {user?.email || "-"}
        </ListElementSetting>
        <ListElementSetting to="password" name="Password">
          -
        </ListElementSetting>
      </List>
      <List className={style.list}>
        <button
          className={style.logout}
          onClick={() => {
            signOut();
          }}
        >
          Sign Out
        </button>
      </List>
      <List className={style.list}>
        <ListElementSetting to="about" name="About Mika"></ListElementSetting>
      </List>
    </>
  );
}

export default SettingsPage;
