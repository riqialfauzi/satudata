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
    default: "Satudata - Portal Data Terbuka",
    template: "%s - Satudata",
  },
  description: "Portal data statistik dan informasi publik Indonesia. Akses dataset, artikel analitis, infografis, dan standar data dari berbagai unit kerja.",
  keywords: ["data terbuka", "statistik", "Indonesia", "dataset", "portal data", "open data"],
  authors: [{ name: "Satudata" }],
  openGraph: {
    type: "website",
    locale: "id_ID",
    siteName: "Satudata",
    title: "Satudata - Portal Data Terbuka",
    description: "Portal data statistik dan informasi publik Indonesia.",
  },
  twitter: {
    card: "summary_large_image",
    title: "Satudata - Portal Data Terbuka",
    description: "Portal data statistik dan informasi publik Indonesia.",
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
      className={`${geistSans.variable} ${geistMono.variable} h-full antialiased`}
    >
      <body className="min-h-full flex flex-col">
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
