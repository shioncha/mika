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
        <ListElementSetting to="/settings/name">Change Name</ListElementSetting>
        <ListElementSetting to="/settings/email">
          Change Email
        </ListElementSetting>
        <ListElementSetting to="/settings/password">
          Change Password
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
        <ListElementSetting to="/settings/about">About Mika</ListElementSetting>
      </List>
    </>
  );
}

export default SettingsPage;
