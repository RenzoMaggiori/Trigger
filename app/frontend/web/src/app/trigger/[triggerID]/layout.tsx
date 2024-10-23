"use client";

import React from "react";

import { MenuProvider } from "@/app/trigger/components/MenuProvider";

export default function Layout({ children }: { children: React.ReactNode }) {
  const [mounted, setMounted] = React.useState(false);

  React.useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) return null;
  return <MenuProvider>{children}</MenuProvider>;
}
