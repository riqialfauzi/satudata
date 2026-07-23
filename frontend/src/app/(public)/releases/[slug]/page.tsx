"use client";

import { useParams } from "next/navigation";
import { useReleaseBySlug, useRelatedReleases } from "@/hooks/useReleases";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Card, CardContent } from "@/components/ui/card";
import {
  Calendar,
  Download,
  FileText,
  TrendingUp,
  Database,
  User,
  Clock,
  ArrowRight,
} from "lucide-react";
import Link from "next/link";
import { ArrowLeft } from "lucide-react";

export default function ReleaseDetailPage() {
  const { slug } = useParams<{ slug: string }>();
  const { data: release, isLoading, error } = useReleaseBySlug(slug);
  const { data: related } = useRelatedReleases(release?.id || "", 4);

  if (isLoading) {
    return (
      <div className="mx-auto max-w-4xl px-4 py-8">
        <Skeleton className="h-8 w-64" />
        <Skeleton className="mt-4 h-64 w-full rounded-xl" />
      </div>
    );
  }

  if (error || !release) {
    return (
      <div className="mx-auto max-w-4xl px-4 py-20 text-center">
        <p className="text-lg text-destructive">Release tidak ditemukan</p>
        <Link href="/releases" className="mt-4 inline-block text-sm text-primary hover:underline">
          &larr; Kembali ke daftar
        </Link>
      </div>
    );
  }

  const isDataset = release.release_type === "dataset";
  const dm = release.dataset_metadata;
  const am = release.article_metadata;

  return (
    <div className="mx-auto max-w-4xl px-4 py-8">
      {/* Back */}
      <Link
        href="/releases"
        className="mb-6 inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground"
      >
        <ArrowLeft className="h-4 w-4" /> Kembali
      </Link>

      {/* Header */}
      <div className="flex flex-wrap items-center gap-3">
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
        <Badge variant="outline">{release.year}</Badge>
        {release.status === "published" && (
          <Badge variant="default" className="bg-green-600">
            Published
          </Badge>
        )}
      </div>

      <h1 className="mt-4 text-3xl font-bold">{release.title}</h1>

      {release.description && (
        <p className="mt-2 text-lg text-muted-foreground">
          {release.description}
        </p>
      )}

      {/* Meta */}
      <div className="mt-4 flex flex-wrap gap-4 text-sm text-muted-foreground">
        <span className="flex items-center gap-1">
          <Calendar className="h-4 w-4" />
          {new Date(release.created_at).toLocaleDateString("id-ID", {
            year: "numeric",
            month: "long",
            day: "numeric",
          })}
        </span>
        <span className="flex items-center gap-1">
          <TrendingUp className="h-4 w-4" />
          {release.view_count} dilihat
        </span>
        {am?.author_name && (
          <span className="flex items-center gap-1">
            <User className="h-4 w-4" />
            {am.author_name}
          </span>
        )}
        {am?.reading_time_minutes && (
          <span className="flex items-center gap-1">
            <Clock className="h-4 w-4" />
            {am.reading_time_minutes} menit baca
          </span>
        )}
      </div>

      {release.tags && release.tags.length > 0 && (
        <div className="mt-4 flex flex-wrap gap-2">
          {release.tags.map((tag) => (
            <Badge key={tag} variant="outline" className="text-xs">
              {tag}
            </Badge>
          ))}
        </div>
      )}

      {/* Dataset view */}
      {isDataset && dm && (
        <Card className="mt-8">
          <CardContent className="p-6">
            <h2 className="flex items-center gap-2 text-lg font-semibold">
              <Database className="h-5 w-5 text-primary" />
              Metadata Dataset
            </h2>
            <div className="mt-4 grid gap-4 sm:grid-cols-2">
              <div>
                <span className="text-xs text-muted-foreground">Format</span>
                <p className="font-medium uppercase">{dm.file_format}</p>
              </div>
              <div>
                <span className="text-xs text-muted-foreground">Ukuran</span>
                <p className="font-medium">
                  {dm.file_size
                    ? `${(dm.file_size / 1024).toFixed(1)} KB`
                    : "-"}
                </p>
              </div>
              {dm.row_count !== undefined && dm.row_count !== null && (
                <div>
                  <span className="text-xs text-muted-foreground">Baris</span>
                  <p className="font-medium">
                    {dm.row_count.toLocaleString()}
                  </p>
                </div>
              )}
              {dm.column_count !== undefined && dm.column_count !== null && (
                <div>
                  <span className="text-xs text-muted-foreground">Kolom</span>
                  <p className="font-medium">{dm.column_count}</p>
                </div>
              )}
              {dm.data_source && (
                <div>
                  <span className="text-xs text-muted-foreground">
                    Sumber Data
                  </span>
                  <p className="font-medium">{dm.data_source}</p>
                </div>
              )}
              {dm.update_frequency && (
                <div>
                  <span className="text-xs text-muted-foreground">
                    Frekuensi Update
                  </span>
                  <p className="font-medium capitalize">
                    {dm.update_frequency}
                  </p>
                </div>
              )}
            </div>
            {dm.file_url && (
              <a
                href={dm.file_url}
                target="_blank"
                rel="noopener noreferrer"
                className="mt-6 inline-flex items-center gap-2 rounded-lg bg-primary px-6 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
              >
                <Download className="h-4 w-4" />
                Download Dataset
              </a>
            )}
          </CardContent>
        </Card>
      )}

      {/* Article view */}
      {am?.content && (
        <Card className="mt-8">
          <CardContent className="p-6">
            <h2 className="flex items-center gap-2 text-lg font-semibold">
              <FileText className="h-5 w-5 text-primary" />
              Konten Artikel
            </h2>
            <div
              className="prose prose-sm dark:prose-invert mt-4 max-w-none"
              dangerouslySetInnerHTML={{ __html: am.content }}
            />
          </CardContent>
        </Card>
      )}

      {/* Related Releases */}
      {related && related.length > 0 && (
        <section className="mt-12">
          <div className="mb-6 flex items-center justify-between">
            <h2 className="text-xl font-bold">Terkait</h2>
            <Link
              href="/releases"
              className="flex items-center gap-1 text-sm text-primary hover:underline"
            >
              Lihat Semua <ArrowRight className="h-4 w-4" />
            </Link>
          </div>
          <div className="grid gap-4 md:grid-cols-2">
            {related.map((item) => (
              <Link key={item.id} href={`/releases/${item.slug}`}>
                <Card className="h-full transition-shadow hover:shadow-md">
                  <CardContent className="p-5">
                    <div className="flex items-start justify-between">
                      <Badge
                        variant={
                          item.release_type === "dataset"
                            ? "default"
                            : item.release_type === "article"
                            ? "secondary"
                            : "outline"
                        }
                      >
                        {item.release_type}
                      </Badge>
                      <span className="text-xs text-muted-foreground">
                        {item.year}
                      </span>
                    </div>
                    <h3 className="mt-3 font-semibold line-clamp-2">
                      {item.title}
                    </h3>
                    {item.description && (
                      <p className="mt-1 text-sm text-muted-foreground line-clamp-2">
                        {item.description}
                      </p>
                    )}
                  </CardContent>
                </Card>
              </Link>
            ))}
          </div>
        </section>
      )}
    </div>
  );
}
