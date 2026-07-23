import { BarChart3, Mail, Phone, MapPin } from "lucide-react";
import Link from "next/link";

const navLinks = [
  { href: "/", label: "Beranda" },
  { href: "/releases", label: "Open Data" },
  { href: "/statistik-sektoral", label: "Statistik Sektoral" },
  { href: "/ppid", label: "PPID" },
];

export function Footer() {
  return (
    <footer className="border-t border-border/50 bg-gradient-to-b from-background to-muted/30">
      {/* Wave decoration */}
      <div className="h-1 w-full bg-gradient-to-r from-primary/20 via-primary/40 to-primary/20" />

      <div className="mx-auto max-w-7xl px-4 py-12 sm:px-6">
        <div className="grid gap-10 md:grid-cols-12">
          {/* Brand */}
          <div className="md:col-span-5">
            <Link href="/" className="group inline-flex items-center gap-2.5">
              <div className="flex h-9 w-9 items-center justify-center rounded-lg bg-gradient-to-br from-primary to-primary/70 shadow-sm">
                <BarChart3 className="h-5 w-5 text-primary-foreground" />
              </div>
              <span className="text-lg font-bold tracking-tight">
                Aceh Besar <span className="text-primary">Satu</span>
              </span>
            </Link>
            <p className="mt-3 max-w-sm text-sm leading-relaxed text-muted-foreground">
              Portal Satu Data Aceh Besar — pusat integrasi dan penyebarluasan data Pemerintah Kabupaten Aceh Besar yang menjamin data akurat, mutakhir, terstandar, dan dapat dipertanggungjawabkan.
            </p>
          </div>

          {/* Navigasi */}
          <div className="md:col-span-3">
            <h3 className="mb-4 text-sm font-semibold uppercase tracking-wider text-foreground">
              Navigasi
            </h3>
            <ul className="space-y-2.5">
              {navLinks.map((link) => (
                <li key={link.href}>
                  <Link
                    href={link.href}
                    className="text-sm text-muted-foreground transition-colors hover:text-foreground"
                  >
                    {link.label}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          {/* Kontak */}
          <div className="md:col-span-4">
            <h3 className="mb-4 text-sm font-semibold uppercase tracking-wider text-foreground">
              Kontak
            </h3>
            <ul className="space-y-3">
              <li className="flex items-start gap-2.5 text-sm text-muted-foreground">
                <Mail className="mt-0.5 h-4 w-4 shrink-0 text-primary" />
                <span>info@acehbesarsatu.go.id</span>
              </li>
              <li className="flex items-start gap-2.5 text-sm text-muted-foreground">
                <Phone className="mt-0.5 h-4 w-4 shrink-0 text-primary" />
                <span>(0651) 1234-5678</span>
              </li>
              <li className="flex items-start gap-2.5 text-sm text-muted-foreground">
                <MapPin className="mt-0.5 h-4 w-4 shrink-0 text-primary" />
                <span>Kantor Bupati Aceh Besar, Kota Jantho</span>
              </li>
            </ul>
          </div>
        </div>

        {/* Bottom bar */}
        <div className="mt-10 flex flex-col items-center justify-between gap-4 border-t border-border/50 pt-6 sm:flex-row">
          <p className="text-xs text-muted-foreground">
            &copy; {new Date().getFullYear()} Pemerintah Kabupaten Aceh Besar. All rights reserved.
          </p>
          <div className="flex gap-4 text-xs text-muted-foreground">
            <Link href="/tentang" className="hover:text-foreground transition-colors">
              Kebijakan Privasi
            </Link>
            <Link href="/faq" className="hover:text-foreground transition-colors">
              FAQ
            </Link>
          </div>
        </div>
      </div>
    </footer>
  );
}
