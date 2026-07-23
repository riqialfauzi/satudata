import { apiClient, unwrap } from "./client";
import type { Standard, CreateStandardRequest, UpdateStandardRequest } from "@/types";

export const standardsApi = {
  list: async (): Promise<Standard[]> => {
    const response = await apiClient.get("/api/v1/public/standards");
    return unwrap<Standard[]>(response);
  },

  create: async (data: CreateStandardRequest): Promise<Standard> => {
    const response = await apiClient.post("/api/v1/protected/standards", data);
    return unwrap<Standard>(response);
  },

  update: async (id: string, data: UpdateStandardRequest): Promise<Standard> => {
    const response = await apiClient.put(`/api/v1/protected/standards/${id}`, data);
    return unwrap<Standard>(response);
  },
};
