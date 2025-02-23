import type { Metadata } from "next";
import { Ovo, Acme } from "next/font/google";
import "./globals.css";
import { Toaster } from "@/components/ui/sonner";
import { ThemeProvider } from "@/components/theme-provider";

const ovo = Ovo({
  variable: "--font-ovo",
  subsets: ["latin"],
  weight: "400",
});

const sansCustom = Acme({
  variable: "--font-sans-custom",
  subsets: ["latin"],
  weight: "400",
});

export const metadata: Metadata = {
  title: "Zen0-tren",
  description: "",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={`${ovo.variable} ${sansCustom.variable} antialiased`}>
        <ThemeProvider attribute="class" defaultTheme="system">
          {children}
          <Toaster />
        </ThemeProvider>
      </body>
    </html>
  );
}
