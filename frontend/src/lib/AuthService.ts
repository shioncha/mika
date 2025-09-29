import { AxiosError } from "axios";

import apiClient from "./api";

export interface SignUpCredentials {
  name: string;
  email: string;
  password: string;
  password_confirm: string;
}

export interface SignInCredentials {
  email: string;
  password: string;
}

export interface AuthResponse {
  token: string;
}

export const authService = {
  /**
   * サインアップ
   * @param credentials ユーザー名、メールアドレス、パスワード
   * @return サインアップ後の認証レスポンスを解決するPromise
   */
  async signUp(credentials: SignUpCredentials): Promise<AuthResponse> {
    const { data } = await apiClient.post<AuthResponse>(
      "/sign-up",
      credentials
    );
    return data;
  },

  /**
   * サインイン
   * @param credentials ユーザーのメールアドレスとパスワード
   * @returns アクセストークンを含む認証レスポンスを解決するPromise
   */
  async signIn(credentials: SignInCredentials): Promise<AuthResponse> {
    const { data } = await apiClient.post<AuthResponse>(
      "/sign-in",
      credentials
    );
    return data;
  },

  /**
   * サインアウト
   */
  async signOut(): Promise<void> {
    try {
      await apiClient.post("/sign-out");
    } catch (error) {
      if (error instanceof AxiosError && error.response?.status !== 401) {
        console.error("Sign out request failed:", error);
      }
    }
  },
};
