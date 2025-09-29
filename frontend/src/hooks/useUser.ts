import { useContext } from "react";

import { UserContext } from "./userContext";

function useUser() {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}

export { useUser };
