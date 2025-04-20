import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Coffee Coin",
  description: "A platform to revolutionize coffee trading with blockchain technology.",
  openGraph: {
    title: "Coffee Coin",
    description: "A platform to revolutionize coffee trading with blockchain technology.",
    images: [
      {
        url: "/preview.png",
        width: 1200,
        height: 630,
        alt: "Coffee Coin - Revolutionizing coffee trading with blockchain",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "Coffee Coin",
    description: "A platform to revolutionize coffee trading with blockchain technology.",
    images: ["/preview.png"],
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
