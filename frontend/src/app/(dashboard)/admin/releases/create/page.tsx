"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { releasesApi } from "@/lib/api/releases";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { ArrowLeft } from "lucide-react";
import Link from "next/link";

const createSchema = z.object({
  title: z.string().min(1, "Judul wajib diisi"),
  release_type: z.enum(["dataset", "article", "infographic"]),
  year: z.number().min(2000, "Tahun tidak valid").max(2100),
  description: z.string().optional(),
  tags: z.string().optional(),
  // Dataset fields
  file_url: z.string().optional(),
  file_format: z.string().optional(),
  data_source: z.string().optional(),
  // Article fields
  content: z.string().optional(),
  excerpt: z.string().optional(),
  author_name: z.string().optional(),
  category: z.string().optional(),
});

type CreateForm = z.infer<typeof createSchema>;

export default function CreateReleasePage() {
  const router = useRouter();
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<CreateForm>({
    resolver: zodResolver(createSchema),
    defaultValues: {
      release_type: "dataset",
      year: new Date().getFullYear(),
    },
  });

  const releaseType = watch("release_type");

  const onSubmit = async (data: CreateForm) => {
    setSubmitting(true);
    setError(null);
    try {
      const tags = data.tags
        ? data.tags.split(",").map((t) => t.trim()).filter(Boolean)
        : undefined;

      await releasesApi.create({
        title: data.title,
        release_type: data.release_type,
        year: data.year,
        description: data.description || undefined,
        tags,
        file_url: data.file_url || undefined,
        file_format: data.file_format || undefined,
        data_source: data.data_source || undefined,
        content: data.content || undefined,
        excerpt: data.excerpt || undefined,
        author_name: data.author_name || undefined,
        category: data.category || undefined,
      });

      router.push("/admin/releases");
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : "Gagal membuat release");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="max-w-2xl">
      <Link
        href="/admin/releases"
        className="mb-4 inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground"
      >
        <ArrowLeft className="h-4 w-4" /> Kembali
      </Link>

      <h1 className="text-2xl font-bold">Buat Release Baru</h1>

      {error && (
        <div className="mt-4 rounded-lg border border-destructive/50 bg-destructive/10 p-3 text-sm text-destructive">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit(onSubmit)} className="mt-6 space-y-4">
        {/* Title */}
        <div className="space-y-1">
          <label className="text-sm font-medium">Judul *</label>
          <input
            {...register("title")}
            className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            placeholder="Judul release"
          />
          {errors.title && (
            <p className="text-xs text-destructive">{errors.title.message}</p>
          )}
        </div>

        {/* Type & Year */}
        <div className="grid gap-4 sm:grid-cols-2">
          <div className="space-y-1">
            <label className="text-sm font-medium">Tipe *</label>
            <select
              {...register("release_type")}
              className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            >
              <option value="dataset">Dataset</option>
              <option value="article">Artikel</option>
              <option value="infographic">Infografis</option>
            </select>
          </div>
          <div className="space-y-1">
            <label className="text-sm font-medium">Tahun *</label>
            <input
              type="number"
              {...register("year", { valueAsNumber: true })}
              className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            />
            {errors.year && (
              <p className="text-xs text-destructive">{errors.year.message}</p>
            )}
          </div>
        </div>

        {/* Description */}
        <div className="space-y-1">
          <label className="text-sm font-medium">Deskripsi</label>
          <textarea
            {...register("description")}
            rows={3}
            className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            placeholder="Deskripsi release"
          />
        </div>

        {/* Tags */}
        <div className="space-y-1">
          <label className="text-sm font-medium">Tags</label>
          <input
            {...register("tags")}
            className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            placeholder="pisahkan dengan koma, contoh: statistik, ekonomi"
          />
        </div>

        {/* Dataset fields */}
        {releaseType === "dataset" && (
          <>
            <div className="grid gap-4 sm:grid-cols-2">
              <div className="space-y-1">
                <label className="text-sm font-medium">File URL</label>
                <input
                  {...register("file_url")}
                  className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                  placeholder="https://..."
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm font-medium">Format File</label>
                <select
                  {...register("file_format")}
                  className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                >
                  <option value="">Pilih format</option>
                  <option value="csv">CSV</option>
                  <option value="json">JSON</option>
                  <option value="xlsx">XLSX</option>
                  <option value="parquet">Parquet</option>
                </select>
              </div>
            </div>
            <div className="space-y-1">
              <label className="text-sm font-medium">Sumber Data</label>
              <input
                {...register("data_source")}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                placeholder="Sumber data"
              />
            </div>
          </>
        )}

        {/* Article fields */}
        {releaseType === "article" && (
          <>
            <div className="space-y-1">
              <label className="text-sm font-medium">Penulis</label>
              <input
                {...register("author_name")}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                placeholder="Nama penulis"
              />
            </div>
            <div className="space-y-1">
              <label className="text-sm font-medium">Kategori</label>
              <input
                {...register("category")}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                placeholder="Kategori artikel"
              />
            </div>
            <div className="space-y-1">
              <label className="text-sm font-medium">Excerpt</label>
              <textarea
                {...register("excerpt")}
                rows={2}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                placeholder="Ringkasan singkat"
              />
            </div>
            <div className="space-y-1">
              <label className="text-sm font-medium">Konten (HTML)</label>
              <textarea
                {...register("content")}
                rows={8}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm font-mono"
                placeholder="<h1>Judul</h1><p>Konten artikel...</p>"
              />
            </div>
          </>
        )}

        <button
          type="submit"
          disabled={submitting}
          className="rounded-lg bg-primary px-6 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90 disabled:opacity-50"
        >
          {submitting ? "Menyimpan..." : "Simpan"}
        </button>
      </form>
    </div>
  );
}
