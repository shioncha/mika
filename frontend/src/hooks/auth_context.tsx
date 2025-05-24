import { createContext, useEffect, useState } from "react";

type AuthContextType = {
  isAuthenticated: boolean;
  token: string | null;
  refreshToken: string | null;
  login: (token: string, refreshToken: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
  token: null,
  refreshToken: null,
  login: () => {},
  logout: () => {},
});

function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [token, setToken] = useState<string | null>(null);
  const [refreshToken, setRefreshToken] = useState<string | null>(null);

  useEffect(() => {
    const storedToken = window.localStorage.getItem("token");
    const storedRefreshToken = window.localStorage.getItem("refresh_token");
    if (storedToken && storedRefreshToken) {
      setIsAuthenticated(true);
      setToken(storedToken);
      setRefreshToken(storedRefreshToken);
    }
  }, []);

  const login = (newToken: string, newRefreshToken: string) => {
    setIsAuthenticated(true);
    setToken(newToken);
    setRefreshToken(newRefreshToken);
  };

  const logout = () => {
    setIsAuthenticated(false);
    setToken(null);
    setRefreshToken(null);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, token, refreshToken, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export { AuthContext, AuthProvider };
export default AuthProvider;
