"use client";

import { useSearchParams, useRouter, usePathname } from "next/navigation";
import { useReleases } from "@/hooks/useReleases";
import Link from "next/link";
import { Search, TrendingUp, ChevronLeft, ChevronRight } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useCallback } from "react";

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
  const page = parseInt(searchParams.get("page") || "1");

  const { data, isLoading, error } = useReleases({
    type: type || undefined,
    year: year ? parseInt(year) : undefined,
    q: q || undefined,
    page,
    limit: 10,
  });

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

  return (
    <div className="mx-auto max-w-7xl px-4 py-8">
      <h1 className="text-3xl font-bold">Data & Artikel</h1>
      <p className="mt-1 text-muted-foreground">
        Jelajahi dataset, artikel, dan infografis statistik.
      </p>

      {/* Filters */}
      <div className="mt-6 flex flex-wrap gap-3">
        {/* Search */}
        <div className="relative flex-1 min-w-[200px]">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <input
            type="text"
            placeholder="Cari..."
            defaultValue={q}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                setFilter("q", (e.target as HTMLInputElement).value);
              }
            }}
            className="w-full rounded-lg border bg-background py-2 pl-10 pr-3 text-sm"
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
