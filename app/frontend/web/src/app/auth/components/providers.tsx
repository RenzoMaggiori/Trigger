"use client";

import React from "react";
import Link from "next/link";

import { Button } from "@/components/ui/button";

import { cn } from "@/lib/utils";
import { env } from "@/lib/env";

interface ProvidersProps {
  providers: {
    icon: React.JSX.Element;
    name: string;
    text: string;
    className?: string;
  }[];
}

export function Providers({ providers }: ProvidersProps) {
  return (
    <div className="flex flex-col w-full justify-center items-center text-center gap-5">
      {providers.map((p, key) => (
        <Button
          asChild
          key={key}
          variant="outline"
          className={cn(
            "w-2/3 rounded-full py-5 text-sm md:text-lg",
            p.className ?? "",
          )}
        >
          <Link
            href={`${env.NEXT_PUBLIC_SERVER_URL}/api/oauth2/login?provider=${p.name}&redirect=${env.NEXT_PUBLIC_WEB_URL}/auth/token`}
          >
            {p.icon}
            {p.text}
          </Link>
        </Button>
      ))}
    </div>
  );
}
