import List from "../../components/ui/List";
import ListElementSetting from "../../features/settings/components/ListElementSetting";
import { useAuth } from "../../hooks/useAuth";
import { useUser } from "../../hooks/useUser";
import style from "./Settings.module.css";

function SettingsPage() {
  const { signOut } = useAuth();
  const { user } = useUser();

  return (
    <>
      <h1>Settings</h1>
      <List className={style.list}>
        <ListElementSetting to="name" name="Name">
          {user?.name || "-"}
        </ListElementSetting>
        <ListElementSetting to="email" name="Email">
          {user?.email || "-"}
        </ListElementSetting>
        <ListElementSetting to="password" name="Password">
          ********
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
