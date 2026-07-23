"use client";

import { useStandards } from "@/hooks/useStandards";
import { standardsApi } from "@/lib/api/standards";
import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Plus, FileText, Download, CheckCircle2 } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog";

const standardSchema = z.object({
  title: z.string().min(1, "Judul wajib"),
  year: z.number().min(2000).max(2100),
  description: z.string().optional(),
  file_url: z.string().optional(),
  version: z.string().optional(),
  is_current: z.boolean().optional(),
});

type StandardForm = z.infer<typeof standardSchema>;

export default function AdminStandardsPage() {
  const { data: standards, isLoading, error } = useStandards();
  const queryClient = useQueryClient();
  const [open, setOpen] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [formError, setFormError] = useState<string | null>(null);

  const form = useForm<StandardForm>({
    resolver: zodResolver(standardSchema),
    defaultValues: {
      year: new Date().getFullYear(),
      version: "1.0",
      is_current: false,
    },
  });

  const onSubmit = async (data: StandardForm) => {
    setSubmitting(true);
    setFormError(null);
    try {
      await standardsApi.create({
        title: data.title,
        year: data.year,
        description: data.description || undefined,
        file_url: data.file_url || undefined,
        version: data.version || "1.0",
        is_current: data.is_current || false,
      });
      queryClient.invalidateQueries({ queryKey: ["standards"] });
      setOpen(false);
      form.reset();
    } catch (err: unknown) {
      setFormError(err instanceof Error ? err.message : "Gagal menyimpan");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div>
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold">Standar Data</h1>
          <p className="text-sm text-muted-foreground">
            Kelola standar data nasional.
          </p>
        </div>
        <button
          onClick={() => setOpen(true)}
          className="flex items-center gap-1 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground"
        >
          <Plus className="h-4 w-4" /> Tambah
        </button>
      </div>

      {/* List */}
      <div className="mt-4 space-y-3">
        {isLoading &&
          [...Array(3)].map((_, i) => (
            <Skeleton key={i} className="h-20 rounded-xl" />
          ))}
        {error && (
          <p className="text-destructive">Gagal memuat data</p>
        )}
        {standards?.length === 0 && (
          <p className="text-muted-foreground">Belum ada standar data</p>
        )}
        {standards?.map((std) => (
          <div
            key={std.id}
            className="flex items-center gap-4 rounded-lg border p-4"
          >
            <FileText className="h-8 w-8 text-primary" />
            <div className="flex-1 min-w-0">
              <div className="flex items-center gap-2">
                <span className="font-medium">{std.title}</span>
                {std.is_current && (
                  <CheckCircle2 className="h-4 w-4 text-green-600" />
                )}
                <Badge variant="outline">v{std.version}</Badge>
              </div>
              <p className="text-sm text-muted-foreground">
                Tahun {std.year} &middot; {std.status}
              </p>
            </div>
            {std.file_url && (
              <a
                href={std.file_url}
                target="_blank"
                className="rounded-lg border p-2 hover:bg-muted"
              >
                <Download className="h-4 w-4" />
              </a>
            )}
          </div>
        ))}
      </div>

      {/* Create Dialog */}
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Tambah Standar Data</DialogTitle>
          </DialogHeader>
          {formError && (
            <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3 text-sm text-destructive">
              {formError}
            </div>
          )}
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
            <div className="space-y-1">
              <label className="text-sm font-medium">Judul *</label>
              <input
                {...form.register("title")}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
              />
            </div>
            <div className="grid grid-cols-2 gap-3">
              <div className="space-y-1">
                <label className="text-sm font-medium">Tahun *</label>
                <input
                  type="number"
                  {...form.register("year", { valueAsNumber: true })}
                  className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                />
              </div>
              <div className="space-y-1">
                <label className="text-sm font-medium">Versi</label>
                <input
                  {...form.register("version")}
                  className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                />
              </div>
            </div>
            <div className="space-y-1">
              <label className="text-sm font-medium">Deskripsi</label>
              <textarea
                {...form.register("description")}
                rows={2}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
              />
            </div>
            <div className="space-y-1">
              <label className="text-sm font-medium">File URL</label>
              <input
                {...form.register("file_url")}
                className="w-full rounded-lg border bg-background px-3 py-2 text-sm"
                placeholder="https://..."
              />
            </div>
            <label className="flex items-center gap-2 text-sm">
              <input type="checkbox" {...form.register("is_current")} />
              Jadikan standar aktif
            </label>
            <DialogFooter>
              <button
                type="submit"
                disabled={submitting}
                className="rounded-lg bg-primary px-4 py-2 text-sm text-primary-foreground disabled:opacity-50"
              >
                {submitting ? "Menyimpan..." : "Simpan"}
              </button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>
    </div>
  );
}
