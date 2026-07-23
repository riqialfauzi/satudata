"use client";

import { useQuery } from "@tanstack/react-query";
import { releasesApi } from "@/lib/api/releases";
import type { ReleaseFilter } from "@/types";

export function useReleases(filters?: ReleaseFilter) {
  return useQuery({
    queryKey: ["releases", filters],
    queryFn: () => releasesApi.list(filters),
  });
}

export function useReleaseBySlug(slug: string) {
  return useQuery({
    queryKey: ["release", slug],
    queryFn: () => releasesApi.getBySlug(slug),
    enabled: !!slug,
  });
}

export function useReleaseStats() {
  return useQuery({
    queryKey: ["releases", "stats"],
    queryFn: () => releasesApi.getStats(),
  });
}

export function useRelatedReleases(id: string, limit: number = 5) {
  return useQuery({
    queryKey: ["releases", "related", id],
    queryFn: () => releasesApi.getRelated(id, limit),
    enabled: !!id,
  });
}
