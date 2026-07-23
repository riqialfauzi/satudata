"use client";

import { useParams, useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { releasesApi } from "@/lib/api/releases";
import { useReleaseBySlug } from "@/hooks/useReleases";
import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { ArrowLeft } from "lucide-react";
import Link from "next/link";
import { Skeleton } from "@/components/ui/skeleton";

const editSchema = z.object({
  title: z.string().min(1).optional(),
  description: z.string().optional(),
  status: z.enum(["draft", "published", "archived"]).optional(),
  year: z.number().min(2000).max(2100).optional(),
  tags: z.string().optional(),
});

type EditForm = z.infer<typeof editSchema>;

export default function EditReleasePage() {
  const { id } = useParams<{ id: string }>();
  const router = useRouter();
  const queryClient = useQueryClient();
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Fetch release to get slug for detail reference
  // For edit, we use the ID directly
  const { data: release, isLoading } = useReleaseBySlug(id);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<EditForm>({
    resolver: zodResolver(editSchema),
    values: release
      ? {
          title: release.title,
          description: release.description,
          status: release.status as "draft" | "published" | "archived",
          year: release.year,
          tags: release.tags?.join(", ") || "",
        }
      : undefined,
  });

  const onSubmit = async (data: EditForm) => {
    setSubmitting(true);
    setError(null);
    try {
      const tags = data.tags
        ? data.tags.split(",").map((t) => t.trim()).filter(Boolean)
        : undefined;

      await releasesApi.update(id, {
        title: data.title,
        description: data.description,
        status: data.status,
        year: data.year,
        tags,
      });

      queryClient.invalidateQueries({ queryKey: ["releases"] });
      router.push("/admin/releases");
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : "Gagal mengupdate release");
    } finally {
      setSubmitting(false);
    }
  };

  if (isLoading) {
    return (
      <div className="max-w-2xl">
        <Skeleton className="h-8 w-48" />
        <Skeleton className="mt-4 h-96 w-full rounded-xl" />
      </div>
    );
  }

  if (!release) {
    return (
      <div className="py-20 text-center">
        <p className="text-destructive">Release tidak ditemukan</p>
        <Link href="/admin/releases" className="mt-2 inline-block text-sm text-primary">
          Kembali
        </Link>
      </div>
    );
  }

  return (
    <div className="max-w-2xl">
      <Link
        href="/admin/releases"
        className="mb-4 inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground"
      >
        <ArrowLeft className="h-4 w-4" /> Kembali
      </Link>

      <h1 className="text-2xl font-bold">Edit Release</h1>

      {error && (
        <div className="mt-4 rounded-lg border border-destructive/50 bg-destructive/10 p-3 text-sm text-destructive">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit(onSubmit)} className="mt-6 space-y-4">
        <div className="space-y-1">
          <label className="text-sm font-medium">Judul</label>
          <input
            {...register("title")}
            className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
          />
          {errors.title && (
            <p className="text-xs text-destructive">{errors.title.message}</p>
          )}
        </div>

        <div className="grid gap-4 sm:grid-cols-2">
          <div className="space-y-1">
            <label className="text-sm font-medium">Status</label>
            <select
              {...register("status")}
              className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            >
              <option value="draft">Draft</option>
              <option value="published">Published</option>
              <option value="archived">Archived</option>
            </select>
          </div>
          <div className="space-y-1">
            <label className="text-sm font-medium">Tahun</label>
            <input
              type="number"
              {...register("year")}
              className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            />
          </div>
        </div>

        <div className="space-y-1">
          <label className="text-sm font-medium">Deskripsi</label>
          <textarea
            {...register("description")}
            rows={3}
            className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
          />
        </div>

        <div className="space-y-1">
          <label className="text-sm font-medium">Tags</label>
          <input
            {...register("tags")}
            className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
            placeholder="pisahkan dengan koma"
          />
        </div>

        <button
          type="submit"
          disabled={submitting}
          className="rounded-lg bg-primary px-6 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90 disabled:opacity-50"
        >
          {submitting ? "Menyimpan..." : "Simpan Perubahan"}
        </button>
      </form>
    </div>
  );
}
