"use client";

import Link from "next/link";
import { BarChart3, Menu, X } from "lucide-react";
import { useState } from "react";
import { useAuthStore } from "@/store/authStore";
import { ThemeToggle } from "./ThemeToggle";

export function Header() {
  const [open, setOpen] = useState(false);
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);

  return (
    <header className="sticky top-0 z-50 border-b bg-background/80 backdrop-blur">
      <div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-4">
        <Link href="/" className="flex items-center gap-2 font-bold text-lg">
          <BarChart3 className="h-6 w-6 text-primary" />
          Aceh Besar Satu
        </Link>

        {/* Desktop nav */}
        <nav className="hidden items-center gap-6 md:flex">
          <Link href="/" className="text-sm font-medium hover:text-primary">
            Beranda
          </Link>
          <Link
            href="/releases"
            className="text-sm font-medium hover:text-primary"
          >
            Open Data
          </Link>
          <Link
            href="/statistik-sektoral"
            className="text-sm font-medium hover:text-primary"
          >
            Statistik Sektoral
          </Link>
          <Link href="/ppid" className="text-sm font-medium hover:text-primary">
            PPID
          </Link>
          <ThemeToggle />
          {isAuthenticated ? (
            <Link
              href="/admin"
              className="rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
            >
              Dashboard
            </Link>
          ) : (
            <Link
              href="/login"
              className="rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
            >
              Masuk
            </Link>
          )}
        </nav>

        {/* Mobile toggle */}
        <button
          className="md:hidden"
          onClick={() => setOpen(!open)}
          aria-label="Toggle menu"
        >
          {open ? (
            <X className="h-6 w-6" />
          ) : (
            <Menu className="h-6 w-6" />
          )}
        </button>
      </div>

      {/* Mobile nav */}
      {open && (
        <nav className="border-t p-4 md:hidden">
          <div className="flex flex-col gap-3">
            <Link href="/" className="text-sm font-medium" onClick={() => setOpen(false)}>
              Beranda
            </Link>
            <Link href="/releases" className="text-sm font-medium" onClick={() => setOpen(false)}>
              Open Data
            </Link>
            <Link href="/statistik-sektoral" className="text-sm font-medium" onClick={() => setOpen(false)}>
              Statistik Sektoral
            </Link>
            <Link href="/ppid" className="text-sm font-medium" onClick={() => setOpen(false)}>
              PPID
            </Link>
            <div className="flex items-center gap-2">
              <ThemeToggle />
            </div>
            {isAuthenticated ? (
              <Link href="/admin" className="rounded-lg bg-primary px-4 py-2 text-center text-sm font-medium text-primary-foreground">
                Dashboard
              </Link>
            ) : (
              <Link href="/login" className="rounded-lg bg-primary px-4 py-2 text-center text-sm font-medium text-primary-foreground">
                Masuk
              </Link>
            )}
          </div>
        </nav>
      )}
    </header>
  );
}
