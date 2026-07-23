"use client";

import { useSearchParams, useRouter, usePathname } from "next/navigation";
import { useReleases } from "@/hooks/useReleases";
import { useOrganizations } from "@/hooks/useOrganizations";
import { useGroups } from "@/hooks/useGroups";
import Link from "next/link";
import { TrendingUp, ChevronLeft, ChevronRight, SlidersHorizontal, X } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { SearchBar } from "@/components/ui/search-bar";
import { useCallback, useState } from "react";

const RELEASE_TYPES = [
  { value: "", label: "Semua" },
  { value: "dataset", label: "Dataset" },
  { value: "article", label: "Artikel" },
  { value: "infographic", label: "Infografis" },
];

const YEARS = [2025, 2024, 2023, 2022];

export default function ReleasesPage() {
  const searchParams = useSearchParams();
  const router = useRouter();
  const pathname = usePathname();

  const type = searchParams.get("type") || "";
  const year = searchParams.get("year") || "";
  const q = searchParams.get("q") || "";
  const org = searchParams.get("org") || "";
  const group = searchParams.get("group") || "";
  const page = parseInt(searchParams.get("page") || "1");
  const [showFilters, setShowFilters] = useState(false);

  const { data, isLoading, error } = useReleases({
    type: type || undefined,
    year: year ? parseInt(year) : undefined,
    q: q || undefined,
    page,
    limit: 10,
  });
  const { data: orgs } = useOrganizations();
  const { data: groups } = useGroups();

  const setFilter = useCallback(
    (key: string, value: string) => {
      const params = new URLSearchParams(searchParams.toString());
      if (value) {
        params.set(key, value);
      } else {
        params.delete(key);
      }
      if (key !== "page") params.set("page", "1");
      router.push(`${pathname}?${params.toString()}`);
    },
    [searchParams, router, pathname]
  );

  const activeFilterCount = [type, year, org, group].filter(Boolean).length;

  return (
    <div className="mx-auto max-w-7xl px-4 py-8">
      <h1 className="text-3xl font-bold">Data & Artikel</h1>
      <p className="mt-1 text-muted-foreground">
        Jelajahi dataset, artikel, dan infografis statistik.
      </p>

      {/* Filters */}
      <div className="mt-6 space-y-3">
        {/* Search & Quick Filters Row */}
        <div className="flex flex-wrap items-center gap-3">
          <div className="flex-1 min-w-[200px]">
            <SearchBar
              value={q}
              onChange={(val) => {
                if (!val) setFilter("q", "");
              }}
              onSearch={(val) => setFilter("q", val)}
              placeholder="Cari dataset, artikel..."
            />
          </div>

          {/* Type filter */}
          <select
            value={type}
            onChange={(e) => setFilter("type", e.target.value)}
            className="rounded-lg border bg-background px-3 py-2 text-sm"
          >
            {RELEASE_TYPES.map((t) => (
              <option key={t.value} value={t.value}>
                {t.label}
              </option>
            ))}
          </select>

          {/* Year filter */}
          <select
            value={year}
            onChange={(e) => setFilter("year", e.target.value)}
            className="rounded-lg border bg-background px-3 py-2 text-sm"
          >
            <option value="">Semua Tahun</option>
            {YEARS.map((y) => (
              <option key={y} value={y}>
                {y}
              </option>
            ))}
          </select>

          {/* Toggle facet filter */}
          <button
            onClick={() => setShowFilters(!showFilters)}
            className={`inline-flex items-center gap-2 rounded-lg border px-3 py-2 text-sm transition-colors ${
              showFilters || activeFilterCount > 0
                ? "border-primary bg-primary/10 text-primary"
                : "hover:bg-muted"
            }`}
          >
            <SlidersHorizontal className="h-4 w-4" />
            Filter
            {activeFilterCount > 0 && (
              <Badge variant="default" className="h-5 px-1.5 text-xs">
                {activeFilterCount}
              </Badge>
            )}
          </button>
        </div>

        {/* Facet Filters (collapsible) */}
        {showFilters && (
          <div className="rounded-lg border bg-card p-4">
            <div className="flex items-center justify-between mb-3">
              <h3 className="text-sm font-semibold">Filter Lanjutan</h3>
              <button
                onClick={() => {
                  setFilter("org", "");
                  setFilter("group", "");
                }}
                className="text-xs text-muted-foreground hover:text-foreground"
              >
                Reset filter
              </button>
            </div>
            <div className="grid gap-4 md:grid-cols-2">
              {/* Organization filter */}
              <div>
                <label className="text-xs font-medium text-muted-foreground">
                  Organisasi
                </label>
                <select
                  value={org}
                  onChange={(e) => setFilter("org", e.target.value)}
                  className="mt-1 w-full rounded-lg border bg-background px-3 py-2 text-sm"
                >
                  <option value="">Semua Organisasi</option>
                  {orgs?.map((o) => (
                    <option key={o.id} value={o.slug}>
                      {o.name}
                    </option>
                  ))}
                </select>
              </div>

              {/* Group/Category filter */}
              <div>
                <label className="text-xs font-medium text-muted-foreground">
                  Kategori
                </label>
                <select
                  value={group}
                  onChange={(e) => setFilter("group", e.target.value)}
                  className="mt-1 w-full rounded-lg border bg-background px-3 py-2 text-sm"
                >
                  <option value="">Semua Kategori</option>
                  {groups?.map((g) => (
                    <option key={g.id} value={g.slug}>
                      {g.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          </div>
        )}

        {/* Active filter badges */}
        {activeFilterCount > 0 && (
          <div className="flex flex-wrap gap-2">
            {type && (
              <Badge variant="secondary" className="gap-1">
                Tipe: {type}
                <button onClick={() => setFilter("type", "")}>
                  <X className="h-3 w-3" />
                </button>
              </Badge>
            )}
            {year && (
              <Badge variant="secondary" className="gap-1">
                Tahun: {year}
                <button onClick={() => setFilter("year", "")}>
                  <X className="h-3 w-3" />
                </button>
              </Badge>
            )}
            {org && (
              <Badge variant="secondary" className="gap-1">
                Organisasi: {org}
                <button onClick={() => setFilter("org", "")}>
                  <X className="h-3 w-3" />
                </button>
              </Badge>
            )}
            {group && (
              <Badge variant="secondary" className="gap-1">
                Kategori: {group}
                <button onClick={() => setFilter("group", "")}>
                  <X className="h-3 w-3" />
                </button>
              </Badge>
            )}
          </div>
        )}
      </div>

      {/* Content */}

      {/* Loading */}
      {isLoading && (
        <div className="mt-6 grid gap-4 md:grid-cols-2">
          {[...Array(4)].map((_, i) => (
            <Skeleton key={i} className="h-40 rounded-xl" />
          ))}
        </div>
      )}

      {/* Error */}
      {error && (
        <div className="mt-6 rounded-lg border border-destructive/50 bg-destructive/10 p-4 text-sm text-destructive">
          Gagal memuat data. Silakan coba lagi.
        </div>
      )}

      {/* Empty */}
      {!isLoading && !error && data?.data?.length === 0 && (
        <div className="mt-20 text-center text-muted-foreground">
          <p className="text-lg">Tidak ada release ditemukan</p>
          <p className="text-sm">Coba ubah filter pencarian.</p>
        </div>
      )}

      {/* List */}
      {!isLoading && data?.data && data.data.length > 0 && (
        <>
          <div className="mt-6 grid gap-4 md:grid-cols-2">
            {data.data.map((release) => (
              <Link key={release.id} href={`/releases/${release.slug}`}>
                <Card className="h-full transition-shadow hover:shadow-md">
                  <CardContent className="p-5">
                    <div className="flex items-start justify-between">
                      <Badge
                        variant={
                          release.release_type === "dataset"
                            ? "default"
                            : release.release_type === "article"
                            ? "secondary"
                            : "outline"
                        }
                      >
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
                    <div className="mt-3 flex items-center gap-3 text-xs text-muted-foreground">
                      <span className="flex items-center gap-1">
                        <TrendingUp className="h-3 w-3" />
                        {release.view_count}
                      </span>
                      <span>
                        {new Date(release.created_at).toLocaleDateString(
                          "id-ID"
                        )}
                      </span>
                    </div>
                  </CardContent>
                </Card>
              </Link>
            ))}
          </div>

          {/* Pagination */}
          {data.total > 0 && (
            <div className="mt-8 flex items-center justify-center gap-4">
              <button
                onClick={() => setFilter("page", String(page - 1))}
                disabled={page <= 1}
                className="flex items-center gap-1 rounded-lg border px-3 py-2 text-sm disabled:opacity-50"
              >
                <ChevronLeft className="h-4 w-4" /> Sebelumnya
              </button>
              <span className="text-sm text-muted-foreground">
                Halaman {page}
              </span>
              <button
                onClick={() => setFilter("page", String(page + 1))}
                disabled={data.data.length < 10}
                className="flex items-center gap-1 rounded-lg border px-3 py-2 text-sm disabled:opacity-50"
              >
                Selanjutnya <ChevronRight className="h-4 w-4" />
              </button>
            </div>
          )}
        </>
      )}
    </div>
  );
}
