"use client";

import Image from "next/image";
import Link from "next/link";
import { Menu, X } from "lucide-react";
import { useState } from "react";
import { useAuthStore } from "@/store/authStore";
import { ThemeToggle } from "./ThemeToggle";

const navLinks = [
  { href: "/", label: "Beranda" },
  { href: "/releases", label: "Open Data" },
  { href: "/statistik-sektoral", label: "Statistik Sektoral" },
  { href: "/ppid", label: "PPID" },
];

export function Header() {
  const [open, setOpen] = useState(false);
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);

  return (
    <header className="sticky top-0 z-50 border-b border-border/50 bg-background/70 backdrop-blur-xl supports-[backdrop-filter]:bg-background/60">
      <div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6">
        {/* Logo */}
        <Link href="/" className="group flex items-center gap-2.5">
          <div className="flex h-9 w-9 items-center justify-center rounded-lg bg-gradient-to-br from-primary to-primary/70 shadow-sm transition-transform duration-300 group-hover:scale-105">
            <Image
              src="/lambang-aceh-besar.png"
              alt="Lambang Aceh Besar"
              width={24}
              height={24}
              className="h-6 w-6 object-contain"
            />
          </div>
          <span className="text-lg font-bold tracking-tight">
            Aceh Besar <span className="text-primary">Satu</span>
          </span>
        </Link>

        {/* Desktop nav */}
        <nav className="hidden items-center gap-1 md:flex">
          {navLinks.map((link) => (
            <Link
              key={link.href}
              href={link.href}
              className="relative rounded-lg px-3.5 py-2 text-sm font-medium text-muted-foreground transition-all duration-200 hover:bg-accent hover:text-foreground"
            >
              {link.label}
            </Link>
          ))}
          <div className="ml-2 flex items-center gap-2 border-l border-border/50 pl-4">
            <ThemeToggle />
            {isAuthenticated ? (
              <Link
                href="/admin"
                className="rounded-lg bg-gradient-to-r from-primary to-primary/80 px-4 py-2 text-sm font-medium text-primary-foreground shadow-sm transition-all duration-200 hover:shadow-md hover:brightness-110"
              >
                Dashboard
              </Link>
            ) : (
              <Link
                href="/login"
                className="rounded-lg bg-gradient-to-r from-primary to-primary/80 px-4 py-2 text-sm font-medium text-primary-foreground shadow-sm transition-all duration-200 hover:shadow-md hover:brightness-110"
              >
                Masuk
              </Link>
            )}
          </div>
        </nav>

        {/* Mobile toggle */}
        <button
          className="inline-flex items-center justify-center rounded-lg p-2 text-muted-foreground transition-colors hover:bg-accent hover:text-foreground md:hidden"
          onClick={() => setOpen(!open)}
          aria-label={open ? "Tutup menu" : "Buka menu"}
        >
          {open ? <X className="h-5 w-5" /> : <Menu className="h-5 w-5" />}
        </button>
      </div>

      {/* Mobile nav */}
      {open && (
        <nav className="animate-fade-in border-t border-border/50 px-4 pb-4 pt-2 md:hidden">
          <div className="flex flex-col gap-1">
            {navLinks.map((link) => (
              <Link
                key={link.href}
                href={link.href}
                onClick={() => setOpen(false)}
                className="rounded-lg px-3 py-2.5 text-sm font-medium text-muted-foreground transition-colors hover:bg-accent hover:text-foreground"
              >
                {link.label}
              </Link>
            ))}
            <div className="mt-2 flex items-center gap-3 border-t border-border/50 pt-3">
              <ThemeToggle />
              {isAuthenticated ? (
                <Link
                  href="/admin"
                  onClick={() => setOpen(false)}
                  className="flex-1 rounded-lg bg-gradient-to-r from-primary to-primary/80 px-4 py-2.5 text-center text-sm font-medium text-primary-foreground shadow-sm"
                >
                  Dashboard
                </Link>
              ) : (
                <Link
                  href="/login"
                  onClick={() => setOpen(false)}
                  className="flex-1 rounded-lg bg-gradient-to-r from-primary to-primary/80 px-4 py-2.5 text-center text-sm font-medium text-primary-foreground shadow-sm"
                >
                  Masuk
                </Link>
              )}
            </div>
          </div>
        </nav>
      )}
    </header>
  );
}
