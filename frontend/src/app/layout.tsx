import type { Metadata, Viewport } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { Providers } from "@/components/shared/providers";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: {
    default: "Aceh Besar Satu - Portal Data Aceh Besar",
    template: "%s - Aceh Besar Satu",
  },
  description: "Portal Satu Data Aceh Besar merupakan pusat integrasi dan penyebarluasan data Pemerintah Kabupaten Aceh Besar yang menjamin data akurat, mutakhir, terstandar, dan dapat dipertanggungjawabkan.",
  keywords: ["satu data", "Aceh Besar", "portal data", "open data", "statistik sektoral"],
  authors: [{ name: "Aceh Besar Satu" }],
  openGraph: {
    type: "website",
    locale: "id_ID",
    siteName: "Aceh Besar Satu",
    title: "Aceh Besar Satu - Portal Data Aceh Besar",
    description: "Portal Satu Data Aceh Besar - pusat integrasi dan penyebarluasan data Pemerintah Kabupaten Aceh Besar.",
  },
  twitter: {
    card: "summary_large_image",
    title: "Aceh Besar Satu - Portal Data Aceh Besar",
    description: "Portal Satu Data Aceh Besar - pusat integrasi dan penyebarluasan data Pemerintah Kabupaten Aceh Besar.",
  },
  robots: {
    index: true,
    follow: true,
  },
  manifest: "/manifest.json",
  icons: {
    icon: [
      { url: "/favicon.ico" },
    ],
    apple: [],
  },
};

export const viewport: Viewport = {
  themeColor: "#006792",
  width: "device-width",
  initialScale: 1,
  maximumScale: 1,
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html
      lang="id"
      suppressHydrationWarning
      data-scroll-behavior="smooth"
      className={`${geistSans.variable} ${geistMono.variable} h-full antialiased`}
    >
      <body className="min-h-full flex flex-col">
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
