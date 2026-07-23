"use client";

import { useStandards } from "@/hooks/useStandards";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Card, CardContent } from "@/components/ui/card";
import { FileText, Download, CheckCircle2 } from "lucide-react";

export default function StandardsPage() {
  const { data: standards, isLoading, error } = useStandards();

  return (
    <div className="mx-auto max-w-5xl px-4 py-8">
      <h1 className="text-3xl font-bold">Standar Data</h1>
      <p className="mt-1 text-muted-foreground">
        Standar data nasional yang digunakan untuk menjaga konsistensi dan
        kualitas data.
      </p>

      {isLoading && (
        <div className="mt-6 space-y-4">
          {[...Array(3)].map((_, i) => (
            <Skeleton key={i} className="h-24 rounded-xl" />
          ))}
        </div>
      )}

      {error && (
        <div className="mt-6 rounded-lg border border-destructive/50 bg-destructive/10 p-4 text-sm text-destructive">
          Gagal memuat standar data.
        </div>
      )}

      {standards && standards.length === 0 && (
        <div className="mt-20 text-center text-muted-foreground">
          <p>Belum ada standar data tersedia.</p>
        </div>
      )}

      {standards && standards.length > 0 && (
        <div className="mt-6 space-y-4">
          {standards.map((std) => (
            <Card
              key={std.id}
              className={`transition-shadow hover:shadow-md ${
                std.is_current ? "ring-2 ring-primary/20" : ""
              }`}
            >
              <CardContent className="flex items-start gap-4 p-5">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-primary/10">
                  <FileText className="h-5 w-5 text-primary" />
                </div>
                <div className="flex-1 min-w-0">
                  <div className="flex flex-wrap items-center gap-2">
                    <h3 className="font-semibold">{std.title}</h3>
                    {std.is_current && (
                      <Badge className="bg-green-600 flex items-center gap-1">
                        <CheckCircle2 className="h-3 w-3" /> Aktif
                      </Badge>
                    )}
                    <Badge variant="outline" className="text-xs">
                      v{std.version}
                    </Badge>
                  </div>
                  {std.description && (
                    <p className="mt-1 text-sm text-muted-foreground line-clamp-2">
                      {std.description}
                    </p>
                  )}
                  <div className="mt-2 flex flex-wrap items-center gap-3 text-xs text-muted-foreground">
                    <span>Tahun: {std.year}</span>
                    <span>Status: {std.status}</span>
                    {std.file_size ? (
                      <span>
                        Ukuran: {(std.file_size / 1024).toFixed(1)} KB
                      </span>
                    ) : null}
                  </div>
                </div>
                {std.file_url && (
                  <a
                    href={std.file_url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex shrink-0 items-center gap-1 rounded-lg border px-3 py-2 text-sm hover:bg-muted"
                  >
                    <Download className="h-4 w-4" />
                    Unduh
                  </a>
                )}
              </CardContent>
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}
