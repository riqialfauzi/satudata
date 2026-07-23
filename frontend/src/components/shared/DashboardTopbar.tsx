"use client";

import { useAuth } from "@/hooks/useAuth";
import { ThemeToggle } from "./ThemeToggle";
import { LogOut, Menu } from "lucide-react";

export function DashboardTopbar({
  onMenuClick,
}: {
  onMenuClick?: () => void;
}) {
  const { logout, user } = useAuth();

  return (
    <header className="flex h-16 items-center justify-between border-b bg-card px-4 md:px-6">
      <button className="md:hidden" onClick={onMenuClick} aria-label="Menu">
        <Menu className="h-5 w-5" />
      </button>

      <div className="hidden md:block" />

      <div className="flex items-center gap-3">
        <ThemeToggle />
        <span className="hidden text-sm text-muted-foreground md:inline">
          {user?.email}
        </span>
        <button
          onClick={logout}
          className="flex items-center gap-1 rounded-lg border px-3 py-1.5 text-sm hover:bg-muted"
          title="Logout"
        >
          <LogOut className="h-4 w-4" />
          <span className="hidden md:inline">Keluar</span>
        </button>
      </div>
    </header>
  );
}
