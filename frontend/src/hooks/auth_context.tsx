import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from "react";
import { useNavigate } from "react-router";

import apiClient from "../libs/api";
import { setupAuthHeader } from "../libs/api";
import { authService } from "../libs/AuthService";

type AuthContextType = {
  isAuthenticated: boolean;
  isLoading: boolean;
  signIn: (accessToken: string) => void;
  signOut: () => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

function AuthProvider({ children }: { children: React.ReactNode }) {
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const navigate = useNavigate();

  const isAuthenticated = !!accessToken;

  useEffect(() => {
    const checkAuthStatus = async () => {
      try {
        const response = await apiClient.post("/refresh-token", {
          withCredentials: true,
        });
        setAccessToken(response.data.access_token);
        apiClient.defaults.headers.common[
          "Authorization"
        ] = `Bearer ${response.data.access_token}`;
      } catch {
        console.log("No active session found.");
        setAccessToken(null);
      } finally {
        setIsLoading(false);
      }
    };

    checkAuthStatus();
  }, []);

  useEffect(() => {
    const handleTokenRefreshed = (event: CustomEvent<string>) => {
      console.log("Token refreshed by interceptor.");
      setAccessToken(event.detail);
    };

    const handleSessionExpired = () => {
      console.log("Session expired, signing out.");
      signOut(true); // 強制ログアウトフラグ
    };

    window.addEventListener(
      "tokenRefreshed",
      handleTokenRefreshed as EventListener
    );
    window.addEventListener("sessionExpired", handleSessionExpired);

    return () => {
      window.removeEventListener(
        "tokenRefreshed",
        handleTokenRefreshed as EventListener
      );
      window.removeEventListener("sessionExpired", handleSessionExpired);
    };
  }, []);

  useEffect(() => {
    setupAuthHeader(accessToken);
  }, [accessToken]);

  const signIn = useCallback((newAccessToken: string) => {
    setAccessToken(newAccessToken);
    apiClient.defaults.headers.common[
      "Authorization"
    ] = `Bearer ${newAccessToken}`;
  }, []);

  const signOut = useCallback(
    async (isForced = false) => {
      if (isForced) {
        setAccessToken(null);
        navigate("/signin");
        return;
      }

      try {
        await authService.signOut();
      } catch (error) {
        console.error("Error during sign out:", error);
      } finally {
        setAccessToken(null);
        delete apiClient.defaults.headers.common["Authorization"];
        navigate("/signin");
      }
    },
    [navigate]
  );

  return (
    <AuthContext.Provider
      value={{ isAuthenticated, isLoading, signIn, signOut }}
    >
      {children}
    </AuthContext.Provider>
  );
}

const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

export { AuthProvider, useAuth };
