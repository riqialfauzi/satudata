import { useQuery } from "@tanstack/react-query";
import { groupsApi } from "@/lib/api/groups";

export function useGroups() {
  return useQuery({
    queryKey: ["groups"],
    queryFn: () => groupsApi.list(),
    staleTime: 30 * 60 * 1000, // 30 minutes
  });
}

export function useGroupBySlug(slug: string) {
  return useQuery({
    queryKey: ["group", slug],
    queryFn: () => groupsApi.getBySlug(slug),
    enabled: !!slug,
  });
}
