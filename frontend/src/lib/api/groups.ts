import { apiClient, unwrap } from "./client";
import type { Group } from "@/types";

export const groupsApi = {
  list: async (): Promise<Group[]> => {
    const response = await apiClient.get("/api/v1/public/groups");
    return unwrap<Group[]>(response);
  },

  getBySlug: async (slug: string): Promise<Group> => {
    const response = await apiClient.get(`/api/v1/public/groups/${slug}`);
    return unwrap<Group>(response);
  },
};
