"use client";

import React from "react";
import Image from "next/image";
import Link from "next/link";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { LogoIcon } from "@/components/ui/logoIcon";
import { SideMenu } from "@/app/home/components/sidemenu";
import { useQuery } from "@tanstack/react-query";
import { getWorkspaces } from "@/app/home/lib/get-workspaces";
// import { Workspaces } from "@/app/home/lib/types";


export default function Page() {
  // const [workspaces, setWorkspaces] = React.useState<Workspaces>([])

  const {data, isPending, error} = useQuery({
    queryKey: ["workspaces"],
    queryFn: getWorkspaces,
  });

  // React.useEffect(() => {
  //   if (data) {
  //     setWorkspaces(data);
  //   }
  // }, [data]);

  console.log(error)

  if (isPending) return <div>Loading...</div>
  if (error) return <div>Could not get user workspaces.</div>

    console.log(data)

  return (
    <div className="flex h-screen w-full overflow-hidden">
      <SideMenu />
      <div className="flex flex-col w-full p-5 overflow-x-auto">
        <Card className="py-6">
          <CardContent>
            <ScrollArea>
              <div className="px-8 py-4 w-full">
                <Card className="w-full h-[200px] bg-gradient-to-tr from-blue-500 via-violet-500 to-fuchsia-500 p-5 rounded-lg shadow-lg animate-gradient">
                  <CardContent className="flex flex-col h-full items-center justify-center text-white text-center font-bold text-lg lg:text-3xl gap-y-3">
                    <LogoIcon className="w-[150px] fill-yellow-500" />
                    <p>Try Trigger for 30 days free</p>
                    <Button className="bg-zinc-200 text-black p-5 hover:bg-zinc-100">
                      Start free trial
                    </Button>
                  </CardContent>
                </Card>
              </div>
              <p className="text-3xl font-bold p-5">Your Triggers</p>
              <div className="flex flex-row flex-wrap gap-4 p-5 items-center justify-center">
                {data.map((trigger, index) => (
                  <div key={index}>
                    <Link href={`/triggers/${trigger.id}`}>
                      <Card
                        className="flex flex-col bg-zinc-100 shadow-md rounded-lg w-[200px]"
                        key={index}
                      >
                        <CardHeader className="p-4 border-b">
                          <CardTitle className="text-xl font-bold">
                            <Image
                              src="https://fakeimg.pl/300x200"
                              alt={trigger.id}
                              width={500}
                              height={500}
                              layout="responsive"
                            />

                          </CardTitle>
                        </CardHeader>
                      </Card>
                      <p className="font-bold text-md text-start p-1">
                        Workspace{" "}{index}
                      </p>
                    </Link>
                  </div>
                ))}
              </div>
            </ScrollArea>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
