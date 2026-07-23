import { apiClient, unwrap } from "./client";
import type { Organization } from "@/types";

export const organizationsApi = {
  list: async (): Promise<Organization[]> => {
    const response = await apiClient.get("/api/v1/public/organizations");
    return unwrap<Organization[]>(response);
  },

  getBySlug: async (slug: string): Promise<Organization> => {
    const response = await apiClient.get(`/api/v1/public/organizations/${slug}`);
    return unwrap<Organization>(response);
  },
};
