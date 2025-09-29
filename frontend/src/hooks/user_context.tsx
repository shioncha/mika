import { createContext, useContext, useEffect, useState } from "react";

import { userService } from "../libs/ContentService";
import type { User } from "../type/user";
import { useAuth } from "./auth_context";

type UserContextType = {
  user: User | null;
  setUser: (user: User | null) => void;
  loading: boolean;
  setLoading: (loading: boolean) => void;
  error: string | null;
  setError: (error: string | null) => void;
};

const UserContext = createContext<UserContextType>({
  user: null,
  setUser: () => {},
  loading: false,
  setLoading: () => {},
  error: null,
  setError: () => {},
});

function UserProvider({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuth();
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!isAuthenticated) {
      setUser(null);
      setLoading(false);
      return;
    }
    userService.fetchMe().then((data) => {
      setUser(data);
      setLoading(false);
    });
  }, [isAuthenticated]);

  return (
    <UserContext.Provider
      value={{ user, setUser, loading, setLoading, error, setError }}
    >
      {children}
    </UserContext.Provider>
  );
}

function useUser() {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}

export { UserProvider, useUser };
