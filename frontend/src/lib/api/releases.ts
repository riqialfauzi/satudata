import { apiClient, unwrap } from "./client";
import type {
  Release,
  ReleaseStats,
  ReleaseFilter,
  CreateReleaseRequest,
  UpdateReleaseRequest,
  APIResponse,
} from "@/types";

export const releasesApi = {
  list: async (filters?: ReleaseFilter): Promise<{ data: Release[]; total: number }> => {
    const response = await apiClient.get("/api/v1/public/releases", {
      params: filters,
    });
    const apiResp = response.data as APIResponse<Release[]>;
    return {
      data: apiResp.data || [],
      total: apiResp.meta?.total || 0,
    };
  },

  getById: async (id: string): Promise<Release> => {
    const response = await apiClient.get(`/api/v1/public/releases/${id}`);
    return unwrap<Release>(response);
  },

  getBySlug: async (slug: string): Promise<Release> => {
    const response = await apiClient.get(`/api/v1/public/releases/slug/${slug}`);
    return unwrap<Release>(response);
  },

  getStats: async (): Promise<ReleaseStats> => {
    const response = await apiClient.get("/api/v1/public/releases/stats");
    return unwrap<ReleaseStats>(response);
  },

  create: async (data: CreateReleaseRequest): Promise<Release> => {
    const response = await apiClient.post("/api/v1/protected/releases", data);
    return unwrap<Release>(response);
  },

  update: async (id: string, data: UpdateReleaseRequest): Promise<Release> => {
    const response = await apiClient.put(`/api/v1/protected/releases/${id}`, data);
    return unwrap<Release>(response);
  },

  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/v1/admin/releases/${id}`);
  },
};
