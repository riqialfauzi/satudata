"use client";

import { useQuery } from "@tanstack/react-query";
import { standardsApi } from "@/lib/api/standards";

export function useStandards() {
  return useQuery({
    queryKey: ["standards"],
    queryFn: () => standardsApi.list(),
  });
}
