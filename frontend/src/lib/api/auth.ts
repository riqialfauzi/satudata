import { apiClient, unwrap } from "./client";
import type {
  LoginRequest,
  RegisterRequest,
  TokenResponse,
  User,
} from "@/types";

export const authApi = {
  login: async (data: LoginRequest): Promise<TokenResponse> => {
    const response = await apiClient.post("/api/v1/auth/login", data);
    return unwrap<TokenResponse>(response);
  },

  register: async (data: RegisterRequest): Promise<User> => {
    const response = await apiClient.post("/api/v1/auth/register", data);
    return unwrap<User>(response);
  },

  refresh: async (refreshToken: string): Promise<TokenResponse> => {
    const response = await apiClient.post("/api/v1/auth/refresh", {
      refresh_token: refreshToken,
    });
    return unwrap<TokenResponse>(response);
  },

  logout: async (): Promise<void> => {
    await apiClient.post("/api/v1/auth/logout");
  },

  getProfile: async (): Promise<User> => {
    const response = await apiClient.get("/api/v1/protected/profile");
    return unwrap<User>(response);
  },
};
