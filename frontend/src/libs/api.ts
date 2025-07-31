import axios, { AxiosError } from "axios";

const apiClient = axios.create({
  baseURL: "/api",
  withCredentials: true,
});

apiClient.interceptors.response.use(
  // 成功したレスポンスはそのまま返す
  (response) => response,

  // エラーレスポンスを処理
  async (error: AxiosError) => {
    const originalRequest = error.config;

    if (
      error.response?.status === 401 &&
      originalRequest &&
      !(originalRequest as any)._retry &&
      originalRequest.url !== "/refresh-token" &&
      originalRequest.url !== "/sign-in"
    ) {
      (originalRequest as any)._retry = true;

      try {
        const { data } = await axios.post(
          "/api/refresh-token",
          {},
          { withCredentials: true }
        );

        const event = new CustomEvent("tokenRefreshed", {
          detail: data.token,
        });
        window.dispatchEvent(event);

        originalRequest.headers["Authorization"] = `Bearer ${data.token}`;

        return apiClient(originalRequest);
      } catch (refreshError) {
        const event = new Event("sessionExpired");
        window.dispatchEvent(event);

        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

// リクエストインターセプター（アクセストークンをヘッダーに付与）
// これをAuthContextの外で管理することで、トークンが常に最新に保たれる
export const setupAuthHeader = (token: string | null) => {
  if (token) {
    apiClient.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  } else {
    delete apiClient.defaults.headers.common["Authorization"];
  }
};

export default apiClient;
