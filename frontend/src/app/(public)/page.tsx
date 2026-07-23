"use client";

import Link from "next/link";
import { useReleases, useReleaseStats } from "@/hooks/useReleases";
import { useOrganizations } from "@/hooks/useOrganizations";
import { useGroups } from "@/hooks/useGroups";
import { Search, BarChart3, FileText, Database, TrendingUp, Building2, Tags } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";

export default function HomePage() {
  const { data: stats, isLoading: statsLoading } = useReleaseStats();
  const { data: releases, isLoading: releasesLoading } = useReleases({
    limit: 4,
    sort_by: "published_at",
    sort_dir: "desc",
  });
  const { data: orgs, isLoading: orgsLoading } = useOrganizations();
  const { data: groups, isLoading: groupsLoading } = useGroups();

  return (
    <div className="flex flex-col">
      {/* Hero Section */}
      <section className="bg-gradient-to-br from-primary/5 via-background to-primary/10 py-20 md:py-28">
        <div className="mx-auto max-w-7xl px-4 text-center">
          <h1 className="text-4xl font-bold tracking-tight md:text-6xl">
            Portal Data{" "}
            <span className="text-primary">Terbuka</span> Indonesia
          </h1>
          <p className="mx-auto mt-4 max-w-2xl text-lg text-muted-foreground">
            Akses data statistik, artikel analitis, infografis, dan standar data
            nasional secara terbuka dan transparan.
          </p>

          {/* Search CTA */}
          <div className="mx-auto mt-8 flex max-w-md items-center gap-2 rounded-xl border bg-background p-2 shadow-sm">
            <Search className="ml-2 h-5 w-5 text-muted-foreground" />
            <Link
              href="/releases"
              className="flex-1 py-2 text-sm text-muted-foreground hover:text-foreground"
            >
              Cari data atau artikel...
            </Link>
            <Link
              href="/releases"
              className="rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
            >
              Cari
            </Link>
          </div>
        </div>
      </section>

      {/* Stats Section */}
      <section className="border-t bg-muted/30 py-12">
        <div className="mx-auto max-w-7xl px-4">
          <h2 className="mb-8 text-center text-2xl font-bold">Ringkasan Data</h2>
          {statsLoading ? (
            <div className="grid gap-4 md:grid-cols-4">
              {[...Array(4)].map((_, i) => (
                <Skeleton key={i} className="h-28 rounded-xl" />
              ))}
            </div>
          ) : stats ? (
            <div className="grid gap-4 md:grid-cols-4">
              <Card>
                <CardHeader className="flex flex-row items-center justify-between pb-2">
                  <CardTitle className="text-sm font-medium">
                    Total Release
                  </CardTitle>
                  <Database className="h-4 w-4 text-muted-foreground" />
                </CardHeader>
                <CardContent>
                  <div className="text-3xl font-bold">{stats.total}</div>
                </CardContent>
              </Card>
              {Object.entries(stats.by_type).map(([type, count]) => (
                <Card key={type}>
                  <CardHeader className="flex flex-row items-center justify-between pb-2">
                    <CardTitle className="text-sm font-medium capitalize">
                      {type}
                    </CardTitle>
                    {type === "dataset" ? (
                      <Database className="h-4 w-4 text-muted-foreground" />
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
      </section>

      {/* Latest Releases */}
      <section className="py-16">
        <div className="mx-auto max-w-7xl px-4">
          <div className="mb-8 flex items-center justify-between">
            <h2 className="text-2xl font-bold">Release Terbaru</h2>
            <Link
              href="/releases"
              className="text-sm font-medium text-primary hover:underline"
            >
              Lihat Semua &rarr;
            </Link>
          </div>

          {releasesLoading ? (
            <div className="grid gap-6 md:grid-cols-2">
              {[...Array(4)].map((_, i) => (
                <Skeleton key={i} className="h-40 rounded-xl" />
              ))}
            </div>
          ) : releases?.data?.length ? (
            <div className="grid gap-6 md:grid-cols-2">
              {releases.data.map((release) => (
                <Link key={release.id} href={`/releases/${release.slug}`}>
                  <Card className="h-full transition-shadow hover:shadow-md">
                    <CardContent className="p-6">
                      <div className="flex items-start justify-between">
                        <Badge variant={release.release_type === "dataset" ? "default" : "secondary"}>
                          {release.release_type}
                        </Badge>
                        <span className="text-xs text-muted-foreground">
                          {release.year}
                        </span>
                      </div>
                      <h3 className="mt-3 font-semibold line-clamp-2">
                        {release.title}
                      </h3>
                      {release.description && (
                        <p className="mt-1 text-sm text-muted-foreground line-clamp-2">
                          {release.description}
                        </p>
                      )}
                      <div className="mt-4 flex items-center gap-4 text-xs text-muted-foreground">
                        <span className="flex items-center gap-1">
                          <TrendingUp className="h-3 w-3" />
                          {release.view_count} dilihat
                        </span>
                        <span>{new Date(release.created_at).toLocaleDateString("id-ID")}</span>
                      </div>
                    </CardContent>
                  </Card>
                </Link>
              ))}
            </div>
          ) : (
            <p className="text-center text-muted-foreground">
              Belum ada release tersedia.
            </p>
          )}
        </div>
      </section>

      {/* Featured Organizations */}
      <section className="border-t bg-muted/30 py-16">
        <div className="mx-auto max-w-7xl px-4">
          <div className="mb-8 flex items-center justify-between">
            <h2 className="text-2xl font-bold">Organisasi</h2>
            <Link
              href="/releases"
              className="text-sm font-medium text-primary hover:underline"
            >
              Lihat Semua &rarr;
            </Link>
          </div>

          {orgsLoading ? (
            <div className="grid gap-4 md:grid-cols-3">
              {[...Array(3)].map((_, i) => (
                <Skeleton key={i} className="h-24 rounded-xl" />
              ))}
            </div>
          ) : orgs && orgs.length > 0 ? (
            <div className="grid gap-4 md:grid-cols-3">
              {orgs.slice(0, 6).map((org) => (
                <Card key={org.id} className="transition-shadow hover:shadow-md">
                  <CardContent className="flex items-center gap-4 p-5">
                    <div className="flex h-12 w-12 items-center justify-center rounded-lg bg-primary/10">
                      <Building2 className="h-6 w-6 text-primary" />
                    </div>
                    <div className="min-w-0 flex-1">
                      <h3 className="font-semibold truncate">{org.name}</h3>
                      {org.dataset_count !== undefined && (
                        <p className="text-xs text-muted-foreground">
                          {org.dataset_count} dataset
                        </p>
                      )}
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          ) : null}
        </div>
      </section>

      {/* Featured Groups/Categories */}
      <section className="py-16">
        <div className="mx-auto max-w-7xl px-4">
          <div className="mb-8 flex items-center justify-between">
            <h2 className="text-2xl font-bold">Kategori</h2>
            <Link
              href="/releases"
              className="text-sm font-medium text-primary hover:underline"
            >
              Lihat Semua &rarr;
            </Link>
          </div>

          {groupsLoading ? (
            <div className="flex gap-3">
              {[...Array(4)].map((_, i) => (
                <Skeleton key={i} className="h-20 w-40 rounded-full" />
              ))}
            </div>
          ) : groups && groups.length > 0 ? (
            <div className="flex flex-wrap gap-3">
              {groups.map((group) => (
                <Link key={group.id} href={`/releases?group=${group.slug}`}>
                  <div className="inline-flex items-center gap-2 rounded-full border bg-background px-5 py-2.5 text-sm font-medium transition-colors hover:bg-primary hover:text-primary-foreground hover:border-primary">
                    <Tags className="h-4 w-4" />
                    {group.name}
                  </div>
                </Link>
              ))}
            </div>
          ) : null}
        </div>
      </section>
    </div>
  );
}
