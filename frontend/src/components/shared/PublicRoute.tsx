"use client";

import { useAuthStore } from "@/store/authStore";

/**
 * PublicRoute — No restriction, content visible to everyone.
 * Optionally redirects authenticated users away (e.g. login page).
 */
export function PublicRoute({
  children,
  redirectIfAuthenticated,
}: {
  children: React.ReactNode;
  redirectIfAuthenticated?: string;
}) {
  // Public routes are always accessible. If you need to redirect
  // authenticated users (e.g. /login → /admin), handle it in the page.
  return <>{children}</>;
}
