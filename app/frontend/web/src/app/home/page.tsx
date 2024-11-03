"use client";

import React from "react";
import Image from "next/image";
import Link from "next/link";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { LogoIcon } from "@/components/ui/logoIcon";
import { SideMenu } from "@/app/home/components/sidemenu";
import { useMutation, useQuery } from "@tanstack/react-query";
import { getWorkspaces } from "@/app/home/lib/get-workspaces";
import { Workspaces } from "@/app/home/lib/types";
import { newTrigger } from "@/app/home/lib/new-trigger";
import { Loader2 } from "lucide-react";
import { MdAddBox } from "react-icons/md";


export default function Page() {
  const [workspaces, setWorkspaces] = React.useState<Workspaces>([])
  const [loading, setLoading] = React.useState<boolean>(false);

  const { data, isPending, error } = useQuery({
    queryKey: ["workspaces"],
    queryFn: () => getWorkspaces(),
  });

  const mutation = useMutation({
    mutationFn: newTrigger,
    onSuccess: (trigger) => {
      window.location.href = `/trigger/${trigger.id}`;
    },
  });

  React.useEffect(() => {
    if (data) {
      setWorkspaces(data);
    }
  }, [data]);

  if (isPending) return <div>Loading...</div>
  if (error) return <div>Could not get user workspaces.</div>

  const handleClick = () => {
    setLoading(true)
    mutation.mutate(`Workspace ${workspaces.length}`);
  };

  return (
    <div className="flex h-screen w-full overflow-hidden">
      <SideMenu />
      <div className="flex flex-col w-full p-5 overflow-x-auto">
        <Card className="py-6 h-full">
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
              <div className="flex flex-row flex-wrap gap-4 p-5 items-center justify-start">
                {workspaces.length > 0 ? workspaces.map((trigger, index) => (
                  <div key={index}>
                    <Link href={`/trigger/${trigger.id}`}>
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
                        {trigger.name}
                      </p>
                    </Link>
                  </div>
                )) : (
                  <div className="flex flex-col justify-center items-center h-full w-full">
                    <p className="font-bold text-xl py-2 text-center">No triggers yet? Let{"'"}s change that! Click below to create your first one.</p>
                    <Button
                      onClick={handleClick}
                      className="bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 hover:bg-gradient-to-r hover:from-blue-600 hover:via-violet-600 hover:to-fuchsia-600 animate-gradient text-white mb-5 lg:w-1/4 md:w-1/2 w-2/3 h-1/5"
                      disabled={loading}
                    >
                      {loading ? (
                        <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                      ) : (
                        <MdAddBox className="text-white mr-2 w-5 h-5" />
                      )}
                      <p className="text-xl">Create Trigger</p>
                    </Button>
                  </div>
                )}
              </div>
            </ScrollArea>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
