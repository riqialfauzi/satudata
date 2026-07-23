import { create } from "zustand";
import type { User } from "@/types";

interface AuthState {
  user: User | null;
  isAuthenticated: boolean;
  isAdmin: boolean;
  isLoading: boolean;
  setUser: (user: User | null) => void;
  setLoading: (loading: boolean) => void;
  login: (accessToken: string, refreshToken: string, user: User) => void;
  logout: () => void;
  hydrate: () => void;
}

export const useAuthStore = create<AuthState>((set, get) => ({
  user: null,
  isAuthenticated: false,
  isAdmin: false,
  isLoading: true,

  setUser: (user) =>
    set({
      user,
      isAuthenticated: !!user,
      isAdmin: user?.role === "admin",
    }),

  setLoading: (isLoading) => set({ isLoading }),

  login: (accessToken, refreshToken, user) => {
    localStorage.setItem("access_token", accessToken);
    localStorage.setItem("refresh_token", refreshToken);
    localStorage.setItem("user", JSON.stringify(user));
    set({
      user,
      isAuthenticated: true,
      isAdmin: user.role === "admin",
      isLoading: false,
    });
  },

  logout: () => {
    localStorage.removeItem("access_token");
    localStorage.removeItem("refresh_token");
    localStorage.removeItem("user");
    set({
      user: null,
      isAuthenticated: false,
      isAdmin: false,
      isLoading: false,
    });
  },

  hydrate: () => {
    try {
      const storedUser = localStorage.getItem("user");
      const token = localStorage.getItem("access_token");
      if (storedUser && token) {
        const user = JSON.parse(storedUser) as User;
        set({
          user,
          isAuthenticated: true,
          isAdmin: user.role === "admin",
          isLoading: false,
        });
      } else {
        set({ isLoading: false });
      }
    } catch {
      set({ isLoading: false });
    }
  },
}));
