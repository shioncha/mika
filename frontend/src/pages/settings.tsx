import "./App.css";

import { useContext } from "react";
import { Link } from "react-router";

import { AuthContext } from "../hooks/auth_context";
import { UserContext } from "../hooks/user_context";

function SettingsPage() {
  const { isAuthenticated } = useContext(AuthContext);
  const { user } = useContext(UserContext);

  return (
    <>
      {isAuthenticated ? (
        <p>Hello, {user?.name}!</p>
      ) : (
        <>
          <Link to="/signin">Sign In</Link>
          <Link to="/signup">Sign Up</Link>
        </>
      )}
    </>
  );
}

export default SettingsPage;
