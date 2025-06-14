import { createContext, useContext, useEffect, useState } from "react";

import { AuthContext } from "./auth_context";

type User = {
  id: string;
  name: string;
  email: string;
};

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
  const { isAuthenticated, token } = useContext(AuthContext);
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (isAuthenticated && token) {
      fetch("/api/users/me", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }).then(async (response) => {
        const data: User = await response.json();
        if (response.status === 200) {
          setUser(data);
        } else {
          setError("Failed to fetch user data");
        }
      });
    }
  }, [isAuthenticated, token]);

  return (
    <UserContext.Provider
      value={{ user, setUser, loading, setLoading, error, setError }}
    >
      {children}
    </UserContext.Provider>
  );
}

export { UserContext, UserProvider };
export default UserProvider;
