import { apiClient, unwrap } from "./client";

interface AuditLogEntry {
  id: string;
  action: string;
  resource: string;
  resource_id?: string;
  ip_address?: string;
  user_agent?: string;
  created_at: string;
}

export const auditLogsApi = {
  list: async (params?: { page?: number; limit?: number }): Promise<AuditLogEntry[]> => {
    const response = await apiClient.get("/api/v1/admin/audit-logs", { params });
    return unwrap<AuditLogEntry[]>(response);
  },
};
