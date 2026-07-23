"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useAuthStore } from "@/store/authStore";
import {
  BarChart3,
  Database,
  FileText,
  Users,
  ClipboardList,
  UserCircle,
  LayoutDashboard,
  LogOut,
} from "lucide-react";
import { cn } from "@/lib/utils";

const navItems = [
  { href: "/admin", label: "Dashboard", icon: LayoutDashboard },
  { href: "/admin/releases", label: "Releases", icon: Database },
  { href: "/admin/standards", label: "Standar Data", icon: FileText },
  { href: "/admin/users", label: "Users", icon: Users, adminOnly: true },
  { href: "/admin/audit-logs", label: "Audit Logs", icon: ClipboardList, adminOnly: true },
  { href: "/admin/profile", label: "Profil", icon: UserCircle },
];

export function DashboardSidebar() {
  const pathname = usePathname();
  const user = useAuthStore((s) => s.user);

  return (
    <aside className="hidden w-64 border-r bg-card md:flex md:flex-col">
      {/* Logo */}
      <div className="flex h-16 items-center gap-2 border-b px-6">
        <BarChart3 className="h-6 w-6 text-primary" />
        <span className="font-bold">Aceh Besar Satu</span>
      </div>

      {/* Nav */}
      <nav className="flex-1 space-y-1 p-4">
        {navItems
          .filter((item) => !item.adminOnly || user?.role === "admin")
          .map((item) => {
            const Icon = item.icon;
            const active = pathname === item.href || pathname.startsWith(item.href + "/");
            return (
              <Link
                key={item.href}
                href={item.href}
                className={cn(
                  "flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors",
                  active
                    ? "bg-primary text-primary-foreground"
                    : "text-muted-foreground hover:bg-muted hover:text-foreground"
                )}
              >
                <Icon className="h-4 w-4" />
                {item.label}
              </Link>
            );
          })}
      </nav>

      {/* User info */}
      <div className="border-t p-4">
        <div className="flex items-center gap-3">
          <div className="flex h-8 w-8 items-center justify-center rounded-full bg-primary/10 text-sm font-bold text-primary">
            {user?.full_name?.charAt(0) || "U"}
          </div>
          <div className="flex-1 min-w-0">
            <p className="text-sm font-medium truncate">{user?.full_name}</p>
            <p className="text-xs text-muted-foreground capitalize">{user?.role}</p>
          </div>
        </div>
      </div>
    </aside>
  );
}
