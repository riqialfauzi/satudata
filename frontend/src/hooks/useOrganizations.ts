import { useQuery } from "@tanstack/react-query";
import { organizationsApi } from "@/lib/api/organizations";

export function useOrganizations() {
  return useQuery({
    queryKey: ["organizations"],
    queryFn: () => organizationsApi.list(),
    staleTime: 30 * 60 * 1000, // 30 minutes - orgs don't change often
  });
}

export function useOrganizationBySlug(slug: string) {
  return useQuery({
    queryKey: ["organization", slug],
    queryFn: () => organizationsApi.getBySlug(slug),
    enabled: !!slug,
  });
}
