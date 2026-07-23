import { BarChart3 } from "lucide-react";
import Link from "next/link";

export function Footer() {
  return (
    <footer className="border-t bg-muted/50">
      <div className="mx-auto max-w-7xl px-4 py-8">
        <div className="grid gap-8 md:grid-cols-3">
          <div>
            <div className="flex items-center gap-2 font-bold text-lg">
              <BarChart3 className="h-5 w-5 text-primary" />
              Satudata
            </div>
            <p className="mt-2 text-sm text-muted-foreground">
              Portal data terbuka untuk informasi statistik dan publik Indonesia.
            </p>
          </div>
          <div>
            <h3 className="mb-3 text-sm font-semibold">Navigasi</h3>
            <ul className="space-y-2 text-sm text-muted-foreground">
              <li>
                <Link href="/" className="hover:text-foreground">
                  Beranda
                </Link>
              </li>
              <li>
                <Link href="/releases" className="hover:text-foreground">
                  Data & Artikel
                </Link>
              </li>
              <li>
                <Link href="/standards" className="hover:text-foreground">
                  Standar Data
                </Link>
              </li>
              <li>
                <Link href="/tentang" className="hover:text-foreground">
                  Tentang
                </Link>
              </li>
              <li>
                <Link href="/faq" className="hover:text-foreground">
                  FAQ
                </Link>
              </li>
            </ul>
          </div>
          <div>
            <h3 className="mb-3 text-sm font-semibold">Kontak</h3>
            <ul className="space-y-2 text-sm text-muted-foreground">
              <li>Email: info@satudata.go.id</li>
              <li>Telepon: (021) 1234-5678</li>
            </ul>
          </div>
        </div>
        <div className="mt-8 border-t pt-4 text-center text-xs text-muted-foreground">
          &copy; {new Date().getFullYear()} Satudata. All rights reserved.
        </div>
      </div>
    </footer>
  );
}
