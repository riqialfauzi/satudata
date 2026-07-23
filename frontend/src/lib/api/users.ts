import { apiClient, unwrap } from "./client";
import type { User } from "@/types";

export const usersApi = {
  list: async (): Promise<User[]> => {
    const response = await apiClient.get("/api/v1/admin/users");
    return unwrap<User[]>(response);
  },

  updateRole: async (id: string, role: string): Promise<void> => {
    await apiClient.put(`/api/v1/admin/users/${id}/role`, { role });
  },
};
