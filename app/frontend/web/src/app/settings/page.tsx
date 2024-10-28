"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import React from "react";
import { FcGoogle } from "react-icons/fc";

import { FaDiscord } from "react-icons/fa";
import { IoLogoGithub } from "react-icons/io";
import { FaSlack } from "react-icons/fa6";
import { PiMicrosoftOutlookLogo } from "react-icons/pi";

import { FaCircle } from "react-icons/fa6";
// import { env } from "@/lib/env";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useQuery } from "@tanstack/react-query";
import { getConnections } from "@/app/settings/lib/get-conections";
import { SettingsType } from "@/app/settings/lib/types";
// import { sync } from "./lib/sync";

// import { useMutation } from "@tanstack/react-query";
// import { getConnections } from "./lib/get-conections";

const services = {
  "Google": <FcGoogle className="w-5 h-5" />,
  "Discord": <FaDiscord className="w-5 h-5 text-blue-500" />,
  "Slack": <FaSlack className="w-5 h-5" />,
  "Outlook": <PiMicrosoftOutlookLogo className="w-5 h-5 text-black" />,
  "Github": <IoLogoGithub className="w-5 h-5" />,
} as const;

export default function Page() {
  const [settings, setSettings] = React.useState<SettingsType[]>([]);
  /* const mutation = useMutation({
    mutationFn: sync,
  }); */

  const { data, isPending, error } = useQuery({
    queryKey: ["settings"],
    queryFn: getConnections,
  });

  if (error) return <div>failed to get user settings.</div>
  if (isPending) return <div>loading...</div>

  setSettings(
    data.map((s) => ({
      id: s.id,
      userId: s.userId,
      providerName: s.providerName,
      active: s.active,
    }))
  );

  const handleConnectionClick = (
    active: boolean,
    provider: string,
  ) => {
    const updatedServices = settings.find((s) => s.providerName.toLocaleLowerCase() === provider.toLocaleLowerCase());
    if (!updatedServices)
      return
    if (!active) {
      updatedServices.active = true;
      window.location.href = `/api/redirect?provider=${provider}`;
      // mutation.mutate(provider);
    } else
      updatedServices.active = false;
    setSettings([...settings, updatedServices]);
  };

  return (
    <div className="flex flex-col w-full h-full items-center justify-center gap-5">
      <ScrollArea className="w-full md:w-2/3 lg:w-1/2 max-h-[80vh] items-center justify-center border p-5 rounded-md">
        <div className="flex flex-col items-center justify-center gap-5 w-full">
          {Object.entries(services).map(([providerName, icon]) => {
            const matchingSetting = settings.find((s) => s.providerName.toLowerCase() === providerName.toLowerCase());
            return {
              providerName,
              icon,
              active: matchingSetting?.active || false,
            };
          }).map((item, key) => (
            <Card className="flex flex-col w-full h-auto" key={key}>
              <CardHeader className={`rounded-t-md border-b`}>
                <CardTitle className="flex items-center justify-between text-xl text-start font-bold">
                  <div className="flex items-center gap-x-2">
                    {item.icon}
                    {item.providerName}
                  </div>
                  <div
                    className={`flex items-center ${item.active ? "text-green-500" : "text-red-500"}`}
                  >
                    <div className="hidden md:block">
                      {item.active ? "Connected" : "Disconnected"}
                    </div>
                    <FaCircle
                      className={`ml-2 ${item.active ? "text-green-500" : "text-red-500"}`}
                    />
                  </div>
                </CardTitle>
              </CardHeader>
              <CardContent
                className={`flex w-full h-full rounded-b-md items-end justify-center`}
              >
                <div className="w-full flex flex-col text-lg items-start justify-start gap-y-3 mt-5">
                  <Dialog>
                    <DialogTrigger asChild>
                      <Button variant="outline">
                        {item.active ? "Disconect" : "Connect"}
                      </Button>
                    </DialogTrigger>
                    <DialogContent className="">
                      <DialogHeader>
                        <DialogTitle className="text-2xl">
                          {item.active ? "Disconect" : "Connect"}{" "}
                          {item.providerName}
                        </DialogTitle>
                        <DialogDescription className="text-xl">
                          Are you sure you want to{" "}
                          {item.active ? "disconect" : "connect"}{" "}
                          {item.providerName}?
                        </DialogDescription>
                      </DialogHeader>
                      <DialogFooter className="mt-5">
                        <DialogClose asChild>
                          <Button variant="outline">Cancel</Button>
                        </DialogClose>
                        <DialogClose>
                          <Button
                            onClick={() =>
                              handleConnectionClick(
                                item.active,
                                item.providerName,
                              )
                            }
                          >
                            {item.active ? "Disconect" : "Connect"}
                          </Button>
                        </DialogClose>
                      </DialogFooter>
                    </DialogContent>
                  </Dialog>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </ScrollArea>
    </div>
  );
}
