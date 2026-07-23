import Link from "next/link";
import { ChevronRight, Home } from "lucide-react";
import { cn } from "@/lib/utils";

export interface BreadcrumbItem {
  label: string;
  href?: string;
}

interface BreadcrumbProps {
  items: BreadcrumbItem[];
  className?: string;
  showHome?: boolean;
}

export function Breadcrumb({ items, className, showHome = true }: BreadcrumbProps) {
  return (
    <nav aria-label="Breadcrumb" className={cn("flex items-center text-sm text-muted-foreground", className)}>
      <ol className="flex items-center gap-1.5">
        {showHome && (
          <li className="flex items-center gap-1.5">
            <Link href="/" className="hover:text-foreground transition-colors">
              <Home className="h-4 w-4" />
              <span className="sr-only">Beranda</span>
            </Link>
            <ChevronRight className="h-4 w-4" />
          </li>
        )}
        {items.map((item, index) => (
          <li key={index} className="flex items-center gap-1.5">
            {item.href ? (
              <Link
                href={item.href}
                className="hover:text-foreground transition-colors"
              >
                {item.label}
              </Link>
            ) : (
              <span className="text-foreground font-medium truncate max-w-[200px]">
                {item.label}
              </span>
            )}
            {index < items.length - 1 && (
              <ChevronRight className="h-4 w-4" />
            )}
          </li>
        ))}
      </ol>
    </nav>
  );
}
