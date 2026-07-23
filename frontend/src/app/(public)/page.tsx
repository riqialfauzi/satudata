"use client";

import Link from "next/link";
import { useReleases, useReleaseStats } from "@/hooks/useReleases";
import { useOrganizations } from "@/hooks/useOrganizations";
import { useGroups } from "@/hooks/useGroups";
import {
  Search,
  BarChart3,
  FileText,
  Database,
  TrendingUp,
  Building2,
  Tags,
  ArrowRight,
  BookOpen,
  Download,
  ShieldCheck,
} from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";

// ─── Stat Card ───────────────────────────────────────────
function StatCard({
  label,
  value,
  icon: Icon,
  gradient,
}: {
  label: string;
  value: string | number;
  icon: React.ElementType;
  gradient: string;
}) {
  return (
    <div className="group relative animate-fade-in-up overflow-hidden rounded-2xl border border-border/50 bg-card p-6 shadow-sm transition-all duration-300 hover:-translate-y-1 hover:shadow-lg">
      {/* Gradient accent */}
      <div
        className={`absolute inset-0 opacity-0 transition-opacity duration-300 group-hover:opacity-5 ${gradient}`}
      />
      <div className="relative flex items-start justify-between">
        <div>
          <p className="text-sm font-medium text-muted-foreground">{label}</p>
          <p className="mt-2 text-3xl font-bold tracking-tight">{value}</p>
        </div>
        <div className={`rounded-xl p-3 ${gradient} shadow-sm`}>
          <Icon className="h-5 w-5 text-white" />
        </div>
      </div>
    </div>
  );
}

// ─── Release Card ─────────────────────────────────────────
function ReleaseCard({
  slug,
  type,
  year,
  title,
  description,
  viewCount,
  createdAt,
}: {
  slug: string;
  type: string;
  year: number;
  title: string;
  description?: string;
  viewCount: number;
  createdAt: string;
}) {
  const isDataset = type === "dataset";
  return (
    <Link href={`/releases/${slug}`} className="group block">
      <Card className="h-full overflow-hidden border-border/50 transition-all duration-300 hover:-translate-y-1 hover:border-primary/30 hover:shadow-lg">
        {/* Color top bar */}
        <div
          className={`h-1 w-full bg-gradient-to-r ${
            isDataset
              ? "from-primary to-primary/60"
              : type === "article"
              ? "from-emerald-500 to-emerald-400"
              : "from-violet-500 to-violet-400"
          }`}
        />
        <CardContent className="p-5">
          <div className="flex items-start justify-between gap-3">
            <Badge
              variant={isDataset ? "default" : "secondary"}
              className={`shrink-0 border-0 ${
                isDataset
                  ? "bg-primary/10 text-primary hover:bg-primary/15"
                  : type === "article"
                  ? "bg-emerald-50 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300"
                  : "bg-violet-50 text-violet-700 dark:bg-violet-950 dark:text-violet-300"
              }`}
            >
              {type === "dataset" ? "Dataset" : type === "article" ? "Artikel" : "Infografis"}
            </Badge>
            <span className="shrink-0 text-xs text-muted-foreground">{year}</span>
          </div>
          <h3 className="mt-3 font-semibold leading-snug transition-colors group-hover:text-primary line-clamp-2">
            {title}
          </h3>
          {description && (
            <p className="mt-1.5 text-sm leading-relaxed text-muted-foreground line-clamp-2">
              {description}
            </p>
          )}
          <div className="mt-4 flex items-center gap-4 text-xs text-muted-foreground">
            <span className="flex items-center gap-1">
              <TrendingUp className="h-3.5 w-3.5" />
              {viewCount} dilihat
            </span>
            <span>{new Date(createdAt).toLocaleDateString("id-ID", { year: "numeric", month: "short", day: "numeric" })}</span>
            <span className="ml-auto flex items-center gap-1 text-primary opacity-0 transition-opacity group-hover:opacity-100">
              Detail <ArrowRight className="h-3 w-3" />
            </span>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}

// ─── Home Page ────────────────────────────────────────────
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
      {/* ═══════════════════ HERO ═══════════════════ */}
      <section className="relative overflow-hidden py-20 md:py-32">
        {/* Animated background blobs */}
        <div className="pointer-events-none absolute inset-0 overflow-hidden">
          <div className="absolute -left-40 -top-40 h-[500px] w-[500px] animate-float rounded-full bg-gradient-to-br from-primary/10 to-primary/5 blur-3xl" />
          <div className="absolute -bottom-40 -right-40 h-[600px] w-[600px] animate-float rounded-full bg-gradient-to-br from-primary/5 to-primary/10 blur-3xl" style={{ animationDelay: "-3s" }} />
          <div className="absolute left-1/2 top-1/2 h-[300px] w-[300px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-gradient-to-r from-primary/5 to-transparent blur-2xl" />
        </div>

        <div className="relative mx-auto max-w-7xl px-4 text-center">
          {/* Badge */}
          <div className="animate-fade-in-up mb-6 inline-flex items-center gap-1.5 rounded-full border border-primary/20 bg-primary/5 px-4 py-1.5 text-xs font-medium text-primary">
            <ShieldCheck className="h-3.5 w-3.5" />
            Portal Resmi Pemerintah Kabupaten Aceh Besar
          </div>

          <h1 className="animate-fade-in-up text-4xl font-bold tracking-tight md:text-6xl lg:text-7xl">
            Portal Data{" "}
            <span className="gradient-text">Aceh Besar</span>
          </h1>
          <p className="animate-fade-in-up-1 mx-auto mt-5 max-w-2xl text-base leading-relaxed text-muted-foreground md:text-lg">
            Portal Satu Data Aceh Besar merupakan pusat integrasi dan penyebarluasan data Pemerintah Kabupaten Aceh Besar yang menjamin data akurat, mutakhir, terstandar, dan dapat dipertanggungjawabkan untuk mewujudkan transparansi serta akuntabilitas.
          </p>

          {/* Search CTA */}
          <div className="animate-fade-in-up-2 mx-auto mt-10 flex max-w-lg items-center gap-2 rounded-2xl border border-border/50 bg-card/80 p-1.5 shadow-lg shadow-primary/5 backdrop-blur-sm transition-all duration-300 hover:shadow-xl hover:shadow-primary/10">
            <Search className="ml-3 h-5 w-5 shrink-0 text-muted-foreground" />
            <Link
              href="/releases"
              className="flex-1 py-2.5 text-left text-sm text-muted-foreground transition-colors hover:text-foreground"
            >
              Cari dataset, artikel, atau informasi publik...
            </Link>
            <Link
              href="/releases"
              className="inline-flex items-center gap-1.5 rounded-xl bg-gradient-to-r from-primary to-primary/80 px-5 py-2.5 text-sm font-medium text-primary-foreground shadow-sm transition-all duration-200 hover:shadow-md hover:brightness-110"
            >
              Cari <ArrowRight className="h-4 w-4" />
            </Link>
          </div>

          {/* Quick actions */}
          <div className="animate-fade-in-up-3 mt-10 flex flex-wrap justify-center gap-3">
            {[
              { href: "/releases?type=dataset", label: "Dataset", icon: Database },
              { href: "/releases?type=article", label: "Artikel", icon: BookOpen },
              { href: "/standards", label: "Standar Data", icon: Download },
            ].map((item) => (
              <Link
                key={item.label}
                href={item.href}
                className="inline-flex items-center gap-2 rounded-xl border border-border/50 bg-card/60 px-4 py-2.5 text-sm font-medium text-muted-foreground shadow-sm backdrop-blur-sm transition-all duration-200 hover:border-primary/30 hover:bg-primary/5 hover:text-primary hover:shadow-md"
              >
                <item.icon className="h-4 w-4" />
                {item.label}
              </Link>
            ))}
          </div>
        </div>
      </section>

      {/* ═══════════════════ STATS ═══════════════════ */}
      <section className="relative border-y border-border/30 bg-gradient-to-b from-muted/30 to-background py-16">
        {/* Subtle top border accent */}
        <div className="absolute inset-x-0 top-0 h-px bg-gradient-to-r from-transparent via-primary/20 to-transparent" />

        <div className="mx-auto max-w-7xl px-4">
          <div className="mb-10 text-center">
            <h2 className="text-2xl font-bold tracking-tight">Ringkasan Data</h2>
            <p className="mt-1.5 text-sm text-muted-foreground">
              Statistik terkini portal data Aceh Besar
            </p>
          </div>

          {statsLoading ? (
            <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-4">
              {[...Array(4)].map((_, i) => (
                <Skeleton key={i} className="h-32 rounded-2xl" />
              ))}
            </div>
          ) : stats ? (
            <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-4">
              <StatCard
                label="Total Data"
                value={stats.total}
                icon={Database}
                gradient="bg-gradient-to-br from-primary to-blue-600"
              />
              {Object.entries(stats.by_type).map(([type, count], idx) => {
                const gradients = [
                  "bg-gradient-to-br from-emerald-500 to-teal-600",
                  "bg-gradient-to-br from-violet-500 to-purple-600",
                  "bg-gradient-to-br from-amber-500 to-orange-600",
                ];
                const icons: Record<string, React.ElementType> = {
                  dataset: Database,
                  article: FileText,
                  infographic: TrendingUp,
                };
                return (
                  <StatCard
                    key={type}
                    label={
                      type === "dataset"
                        ? "Dataset"
                        : type === "article"
                        ? "Artikel"
                        : "Infografis"
                    }
                    value={count}
                    icon={icons[type] || FileText}
                    gradient={gradients[idx] || gradients[0]}
                  />
                );
              })}
            </div>
          ) : null}
        </div>
      </section>

      {/* ═══════════════════ LATEST RELEASES ═══════════════════ */}
      <section className="py-20">
        <div className="mx-auto max-w-7xl px-4">
          <div className="mb-10 flex flex-col gap-2 sm:flex-row sm:items-end sm:justify-between">
            <div>
              <h2 className="text-2xl font-bold tracking-tight">Release Terbaru</h2>
              <p className="mt-1 text-sm text-muted-foreground">
                Dataset, artikel, dan infografis terkini
              </p>
            </div>
            <Link
              href="/releases"
              className="group inline-flex items-center gap-1.5 text-sm font-medium text-primary transition-colors hover:text-primary/80"
            >
              Lihat Semua
              <ArrowRight className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
            </Link>
          </div>

          {releasesLoading ? (
            <div className="grid gap-6 md:grid-cols-2">
              {[...Array(4)].map((_, i) => (
                <Skeleton key={i} className="h-44 rounded-2xl" />
              ))}
            </div>
          ) : releases?.data?.length ? (
            <div className="grid gap-6 md:grid-cols-2">
              {releases.data.map((release, i) => (
                <div key={release.id} className="animate-fade-in-up" style={{ animationDelay: `${i * 0.1}s` }}>
                  <ReleaseCard
                    slug={release.slug}
                    type={release.release_type}
                    year={release.year}
                    title={release.title}
                    description={release.description}
                    viewCount={release.view_count}
                    createdAt={release.created_at}
                  />
                </div>
              ))}
            </div>
          ) : (
            <div className="rounded-2xl border border-dashed border-border py-16 text-center">
              <Database className="mx-auto h-8 w-8 text-muted-foreground/50" />
              <p className="mt-3 text-sm text-muted-foreground">
                Belum ada release tersedia.
              </p>
            </div>
          )}
        </div>
      </section>

      {/* ═══════════════════ ORGANIZATIONS ═══════════════════ */}
      <section className="border-t border-border/30 bg-gradient-to-b from-muted/20 to-background py-20">
        <div className="mx-auto max-w-7xl px-4">
          <div className="mb-10 flex flex-col gap-2 sm:flex-row sm:items-end sm:justify-between">
            <div>
              <h2 className="text-2xl font-bold tracking-tight">Organisasi</h2>
              <p className="mt-1 text-sm text-muted-foreground">
                Unit kerja yang berkontribusi dalam penyediaan data
              </p>
            </div>
            <Link
              href="/releases"
              className="group inline-flex items-center gap-1.5 text-sm font-medium text-primary transition-colors hover:text-primary/80"
            >
              Lihat Semua
              <ArrowRight className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
            </Link>
          </div>

          {orgsLoading ? (
            <div className="grid gap-4 md:grid-cols-3">
              {[...Array(3)].map((_, i) => (
                <Skeleton key={i} className="h-24 rounded-2xl" />
              ))}
            </div>
          ) : orgs && orgs.length > 0 ? (
            <div className="grid gap-4 md:grid-cols-3">
              {orgs.slice(0, 6).map((org, i) => (
                <div key={org.id} className="animate-fade-in-up" style={{ animationDelay: `${i * 0.08}s` }}>
                  <Card className="group h-full border-border/50 transition-all duration-300 hover:-translate-y-0.5 hover:border-primary/20 hover:shadow-md">
                    <CardContent className="flex items-center gap-4 p-5">
                      <div className="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-gradient-to-br from-primary/10 to-primary/5 shadow-sm transition-transform duration-300 group-hover:scale-110">
                        <Building2 className="h-6 w-6 text-primary" />
                      </div>
                      <div className="min-w-0 flex-1">
                        <h3 className="font-semibold truncate">{org.name}</h3>
                        {org.dataset_count !== undefined && (
                          <p className="mt-0.5 text-xs text-muted-foreground">
                            {org.dataset_count} dataset
                          </p>
                        )}
                      </div>
                      <ArrowRight className="h-4 w-4 shrink-0 text-muted-foreground opacity-0 transition-all group-hover:translate-x-0.5 group-hover:opacity-100" />
                    </CardContent>
                  </Card>
                </div>
              ))}
            </div>
          ) : null}
        </div>
      </section>

      {/* ═══════════════════ CATEGORIES ═══════════════════ */}
      <section className="py-20">
        <div className="mx-auto max-w-7xl px-4">
          <div className="mb-10 text-center">
            <h2 className="text-2xl font-bold tracking-tight">Kategori</h2>
            <p className="mt-1.5 text-sm text-muted-foreground">
              Jelajahi data berdasarkan kategori
            </p>
          </div>

          {groupsLoading ? (
            <div className="flex flex-wrap justify-center gap-3">
              {[...Array(4)].map((_, i) => (
                <Skeleton key={i} className="h-12 w-36 rounded-full" />
              ))}
            </div>
          ) : groups && groups.length > 0 ? (
            <div className="flex flex-wrap justify-center gap-3">
              {groups.map((group, i) => (
                <Link
                  key={group.id}
                  href={`/releases?group=${group.slug}`}
                  className="animate-fade-in-up group inline-flex items-center gap-2 rounded-full border border-border/50 bg-card px-5 py-2.5 text-sm font-medium text-muted-foreground shadow-sm transition-all duration-200 hover:border-primary/30 hover:bg-primary/5 hover:text-primary hover:shadow-md"
                  style={{ animationDelay: `${i * 0.06}s` }}
                >
                  <Tags className="h-4 w-4 transition-transform group-hover:scale-110" />
                  {group.name}
                </Link>
              ))}
            </div>
          ) : null}
        </div>
      </section>
    </div>
  );
}
