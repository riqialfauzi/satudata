"use client";

import { useState } from "react";
import { useReleases } from "@/hooks/useReleases";
import { releasesApi } from "@/lib/api/releases";
import { useQueryClient } from "@tanstack/react-query";
import Link from "next/link";
import { Plus, Search, TrendingUp, Trash2, Edit } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { useAuthStore } from "@/store/authStore";

export default function AdminReleasesPage() {
  const [search, setSearch] = useState("");
  const [page, setPage] = useState(1);
  const [deleteId, setDeleteId] = useState<string | null>(null);
  const [deleting, setDeleting] = useState(false);
  const user = useAuthStore((s) => s.user);
  const queryClient = useQueryClient();

  const { data, isLoading, error } = useReleases({
    q: search || undefined,
    page,
    limit: 10,
  });

  const handleDelete = async () => {
    if (!deleteId) return;
    setDeleting(true);
    try {
      await releasesApi.delete(deleteId);
      queryClient.invalidateQueries({ queryKey: ["releases"] });
      setDeleteId(null);
    } catch (err) {
      console.error(err);
    } finally {
      setDeleting(false);
    }
  };

  return (
    <div>
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold">Releases</h1>
          <p className="text-sm text-muted-foreground">
            Kelola dataset, artikel, dan infografis.
          </p>
        </div>
        <Link
          href="/admin/releases/create"
          className="flex items-center gap-1 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
        >
          <Plus className="h-4 w-4" /> Tambah
        </Link>
      </div>

      {/* Search */}
      <div className="relative mt-4 max-w-sm">
        <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
        <input
          type="text"
          placeholder="Cari..."
          value={search}
          onChange={(e) => {
            setSearch(e.target.value);
            setPage(1);
          }}
          className="w-full rounded-lg border bg-background py-2 pl-10 pr-3 text-sm"
        />
      </div>

      {/* Table */}
      <div className="mt-4 overflow-x-auto rounded-lg border">
        <table className="w-full text-sm">
          <thead className="bg-muted/50">
            <tr>
              <th className="px-4 py-3 text-left font-medium">Judul</th>
              <th className="px-4 py-3 text-left font-medium">Tipe</th>
              <th className="px-4 py-3 text-left font-medium">Status</th>
              <th className="px-4 py-3 text-left font-medium">Tahun</th>
              <th className="px-4 py-3 text-left font-medium">Dilihat</th>
              <th className="px-4 py-3 text-right font-medium">Aksi</th>
            </tr>
          </thead>
          <tbody className="divide-y">
            {isLoading &&
              [...Array(5)].map((_, i) => (
                <tr key={i}>
                  <td colSpan={6} className="px-4 py-3">
                    <Skeleton className="h-6 w-full" />
                  </td>
                </tr>
              ))}
            {error && (
              <tr>
                <td colSpan={6} className="px-4 py-8 text-center text-destructive">
                  Gagal memuat data
                </td>
              </tr>
            )}
            {data?.data?.length === 0 && !isLoading && (
              <tr>
                <td colSpan={6} className="px-4 py-8 text-center text-muted-foreground">
                  Tidak ada release
                </td>
              </tr>
            )}
            {data?.data?.map((release) => (
              <tr key={release.id} className="hover:bg-muted/50">
                <td className="px-4 py-3 font-medium max-w-xs truncate">
                  {release.title}
                </td>
                <td className="px-4 py-3">
                  <Badge variant="outline">{release.release_type}</Badge>
                </td>
                <td className="px-4 py-3">
                  <Badge
                    variant={
                      release.status === "published"
                        ? "default"
                        : release.status === "draft"
                        ? "secondary"
                        : "outline"
                    }
                    className={
                      release.status === "published"
                        ? "bg-green-600"
                        : undefined
                    }
                  >
                    {release.status}
                  </Badge>
                </td>
                <td className="px-4 py-3">{release.year}</td>
                <td className="px-4 py-3">{release.view_count}</td>
                <td className="px-4 py-3 text-right">
                  <div className="flex justify-end gap-1">
                    <Link
                      href={`/admin/releases/${release.id}/edit`}
                      className="rounded-lg border p-1.5 hover:bg-muted"
                    >
                      <Edit className="h-4 w-4" />
                    </Link>
                    {user?.role === "admin" && (
                      <button
                        onClick={() => setDeleteId(release.id)}
                        className="rounded-lg border p-1.5 text-destructive hover:bg-destructive/10"
                      >
                        <Trash2 className="h-4 w-4" />
                      </button>
                    )}
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      {data && (
        <div className="mt-4 flex items-center justify-between text-sm">
          <span className="text-muted-foreground">
            Total: {data.total}
          </span>
          <div className="flex gap-2">
            <button
              onClick={() => setPage((p) => Math.max(1, p - 1))}
              disabled={page <= 1}
              className="rounded-lg border px-3 py-1 disabled:opacity-50"
            >
              Sebelumnya
            </button>
            <button
              onClick={() => setPage((p) => p + 1)}
              disabled={(data?.data?.length || 0) < 10}
              className="rounded-lg border px-3 py-1 disabled:opacity-50"
            >
              Selanjutnya
            </button>
          </div>
        </div>
      )}

      {/* Delete confirmation */}
      <Dialog open={!!deleteId} onOpenChange={() => setDeleteId(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Hapus Release</DialogTitle>
            <DialogDescription>
              Apakah Anda yakin ingin menghapus release ini? Tindakan ini tidak
              dapat dibatalkan.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <button
              onClick={() => setDeleteId(null)}
              className="rounded-lg border px-4 py-2 text-sm"
            >
              Batal
            </button>
            <button
              onClick={handleDelete}
              disabled={deleting}
              className="rounded-lg bg-destructive px-4 py-2 text-sm text-destructive-foreground hover:bg-destructive/90 disabled:opacity-50"
            >
              {deleting ? "Menghapus..." : "Hapus"}
            </button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
