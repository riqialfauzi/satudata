"use client";

import { useAuthStore } from "@/store/authStore";
import { authApi } from "@/lib/api/auth";
import { useRouter } from "next/navigation";
import { useCallback } from "react";
import type { LoginRequest, RegisterRequest } from "@/types";

export function useAuth() {
  const {
    user,
    isAuthenticated,
    isAdmin,
    isLoading,
    login: storeLogin,
    logout: storeLogout,
    setLoading,
  } = useAuthStore();

  const router = useRouter();

  const login = useCallback(
    async (data: LoginRequest) => {
      setLoading(true);
      try {
        const result = await authApi.login(data);
        storeLogin(result.access_token, result.refresh_token, result.user);
        router.push("/admin");
        return { success: true, error: null };
      } catch (err: unknown) {
        const message =
          err instanceof Error ? err.message : "Login gagal";
        return { success: false, error: message };
      } finally {
        setLoading(false);
      }
    },
    [storeLogin, setLoading, router]
  );

  const register = useCallback(
    async (data: RegisterRequest) => {
      setLoading(true);
      try {
        await authApi.register(data);
        return { success: true, error: null };
      } catch (err: unknown) {
        const message =
          err instanceof Error ? err.message : "Registrasi gagal";
        return { success: false, error: message };
      } finally {
        setLoading(false);
      }
    },
    [setLoading]
  );

  const logout = useCallback(async () => {
    try {
      await authApi.logout();
    } catch {
      // ignore
    }
    storeLogout();
    router.push("/");
  }, [storeLogout, router]);

  return {
    user,
    isAuthenticated,
    isAdmin,
    isLoading,
    login,
    register,
    logout,
  };
}
