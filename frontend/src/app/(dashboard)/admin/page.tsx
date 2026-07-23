"use client";

import { useReleaseStats } from "@/hooks/useReleases";
import { useAuthStore } from "@/store/authStore";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Database, FileText, BarChart3 } from "lucide-react";

export default function AdminDashboardPage() {
  const { data: stats, isLoading } = useReleaseStats();
  const user = useAuthStore((s) => s.user);

  return (
    <div>
      <h1 className="text-2xl font-bold">Selamat datang, {user?.full_name}</h1>
      <p className="mt-1 text-muted-foreground">
        Dashboard admin Satudata
      </p>

      {/* Stats */}
      {isLoading ? (
        <div className="mt-6 grid gap-4 md:grid-cols-4">
          {[...Array(4)].map((_, i) => (
            <Skeleton key={i} className="h-28 rounded-xl" />
          ))}
        </div>
      ) : stats ? (
        <div className="mt-6 grid gap-4 md:grid-cols-4">
          <Card>
            <CardHeader className="flex flex-row items-center justify-between pb-2">
              <CardTitle className="text-sm font-medium">Total Releases</CardTitle>
              <Database className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-3xl font-bold">{stats.total}</div>
            </CardContent>
          </Card>
          {Object.entries(stats.by_type).map(([type, count]) => (
            <Card key={type}>
              <CardHeader className="flex flex-row items-center justify-between pb-2">
                <CardTitle className="text-sm font-medium capitalize">{type}</CardTitle>
                {type === "dataset" ? (
                  <BarChart3 className="h-4 w-4 text-muted-foreground" />
                ) : (
                  <FileText className="h-4 w-4 text-muted-foreground" />
                )}
              </CardHeader>
              <CardContent>
                <div className="text-3xl font-bold">{count}</div>
              </CardContent>
            </Card>
          ))}
        </div>
      ) : null}
    </div>
  );
}
