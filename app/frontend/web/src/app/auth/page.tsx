"use client";

import React from "react";
import { useSearchParams } from "next/navigation";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Login } from "@/app/auth/components/login";
import { Register } from "@/app/auth/components/register";

const page = () => {
  const searchParams = useSearchParams();
  const type = searchParams.get("type");
  const defaultValue = type === "login" ? "login" : "register";

  const tabs = [
    {
      name: "Log In",
      value: "login",
      Component: <Login />,
    },
    {
      name: "Register",
      value: "register",
      Component: <Register />,
    },
  ] as const;

  return (
    <div className="flex justify-center items-center h-full">
      <Tabs defaultValue={defaultValue} className="w-3/4 md:w-1/2 lg:w-1/3">
        <TabsList className="grid w-full grid-cols-2">
          {tabs.map((t, index) => (
            <TabsTrigger key={index} value={t.value}>
              {t.name}
            </TabsTrigger>
          ))}
        </TabsList>
        {tabs.map((t, index) => (
          <TabsContent key={index} value={t.value}>
            {t.Component}
          </TabsContent>
        ))}
      </Tabs>
    </div>
  );
};

export default page;
