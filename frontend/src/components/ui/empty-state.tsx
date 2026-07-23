import { cn } from "@/lib/utils";
import { FileQuestion, SearchX, FolderOpen, LucideIcon } from "lucide-react";

interface EmptyStateProps {
  icon?: LucideIcon;
  title: string;
  description?: string;
  action?: React.ReactNode;
  className?: string;
}

const iconMap = {
  default: FileQuestion,
  search: SearchX,
  empty: FolderOpen,
};

export function EmptyState({
  icon: Icon = FileQuestion,
  title,
  description,
  action,
  className,
}: EmptyStateProps) {
  return (
    <div
      className={cn(
        "flex flex-col items-center justify-center py-16 px-4 text-center",
        className
      )}
    >
      <div className="rounded-full bg-muted p-4 mb-4">
        <Icon className="h-8 w-8 text-muted-foreground" />
      </div>
      <h3 className="text-lg font-semibold text-foreground mb-1">{title}</h3>
      {description && (
        <p className="text-sm text-muted-foreground max-w-sm mb-4">
          {description}
        </p>
      )}
      {action && <div>{action}</div>}
    </div>
  );
}

export function SearchEmptyState({
  query,
  className,
}: {
  query?: string;
  className?: string;
}) {
  return (
    <EmptyState
      icon={SearchX}
      title="Tidak ada hasil ditemukan"
      description={
        query
          ? `Pencarian "${query}" tidak menemukan hasil. Coba gunakan kata kunci lain.`
          : "Tidak ada data yang tersedia saat ini."
      }
      className={className}
    />
  );
}

export function DataEmptyState({
  title = "Belum ada data",
  description,
  action,
  className,
}: {
  title?: string;
  description?: string;
  action?: React.ReactNode;
  className?: string;
}) {
  return (
    <EmptyState
      icon={FolderOpen}
      title={title}
      description={description}
      action={action}
      className={className}
    />
  );
}
