"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Switch } from "@/components/ui/switch";
import React from "react";
import { FcGoogle } from "react-icons/fc";

import { FaDiscord } from "react-icons/fa";
import { IoLogoGithub } from "react-icons/io";
import { FaSlack } from "react-icons/fa6";
import { PiMicrosoftOutlookLogo } from "react-icons/pi";

import { FaCircle } from "react-icons/fa6";
import { env } from "@/lib/env";
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

/* import { useMutation } from "@tanstack/react-query";
import { getConnections } from "./lib/get-conections"; */

type SettingsProps = {
  name: string;
  fields: Record<string, boolean>;
  icon?: React.JSX.Element;
};

const services: SettingsProps[] = [
  {
    name: "Google",
    icon: <FcGoogle className="w-5 h-5" />,
    fields: {
      "Show on Profile": true,
      Connected: true,
    },
  },
  {
    name: "Discord",
    icon: <FaDiscord className="w-5 h-5 text-blue-500" />,
    fields: {
      "Show on Profile": true,
      Connected: false,
    },
  },
  {
    name: "Slack",
    icon: <FaSlack className="w-5 h-5" />,
    fields: {
      "Show on Profile": true,
      Connected: true,
    },
  },
  {
    name: "Outlook",
    icon: <PiMicrosoftOutlookLogo className="w-5 h-5 text-black" />,
    fields: {
      "Show on Profile": true,
      Connected: false,
    },
  },
  {
    name: "Github",
    icon: <IoLogoGithub className="w-5 h-5" />,
    fields: {
      "Show on Profile": true,
      Connected: true,
    },
  },
];

export default function Page() {
  const [serviceList, setServiceList] =
    React.useState<SettingsProps[]>(services);

  const handleSwitchChange = (
    serviceIndex: number,
    fieldKey: string,
    provider: string,
  ) => {
    const updatedServices = [...serviceList];
    if (fieldKey === "Connection") {
      const href = updatedServices[serviceIndex].fields[fieldKey]
        ? `${env.NEXT_PUBLIC_SYNC_SERVICE_URL}/api/auth/sync/sync-with?provider=${provider}`
        : `${env.NEXT_PUBLIC_SYNC_SERVICE_URL}/api/auth/sync/sync-with?provider=${provider}`;
      window.location.href = href;
    }
    updatedServices[serviceIndex].fields[fieldKey] =
      !updatedServices[serviceIndex].fields[fieldKey];
    setServiceList(updatedServices);
  };

  /* const mutation = useMutation({
    mutationFn: getConnections,
  }); */

  const handleConnectionClick = (
    active: boolean,
    serviceIndex: number,
    fieldKey: string,
    provider: string,
  ) => {
    const updatedServices = [...serviceList];
    if (!active) {
      const href = updatedServices[serviceIndex].fields[fieldKey]
        ? `${env.NEXT_PUBLIC_SYNC_SERVICE_URL}/api/auth/sync/sync-with?provider=${provider}`
        : `${env.NEXT_PUBLIC_SYNC_SERVICE_URL}/api/auth/sync/sync-with?provider=${provider}`;
      window.location.href = href;
      updatedServices[serviceIndex].fields[fieldKey] = true;
    } else updatedServices[serviceIndex].fields[fieldKey] = false;
    setServiceList(updatedServices);
  };

  return (
    <div className="flex flex-col w-full h-full items-center justify-center gap-5">
      <ScrollArea className="w-full md:w-2/3 lg:w-1/2 max-h-[80vh] items-center justify-center border p-5 rounded-md">
        <div className="flex flex-col items-center justify-center gap-5 w-full">
          {services.map((item, key) => (
            <Card className="flex flex-col w-full h-auto" key={key}>
              <CardHeader className={`rounded-t-md border-b`}>
                <CardTitle className="flex items-center justify-between text-xl text-start font-bold">
                  <div className="flex items-center gap-x-2">
                    {item.icon}
                    {item.name}
                  </div>
                  <div
                    className={`flex items-center ${item.fields["Connection"] ? "text-green-500" : "text-red-500"}`}
                  >
                    <div className="hidden md:block">
                      {item.fields["Connection"] ? "Connected" : "Disconnected"}
                    </div>
                    <FaCircle
                      className={`ml-2 ${item.fields["Connection"] ? "text-green-500" : "text-red-500"}`}
                    />
                  </div>
                </CardTitle>
              </CardHeader>
              <CardContent
                className={`flex w-full h-full rounded-b-md items-end justify-center`}
              >
                <div className="w-full flex flex-col text-lg items-start justify-start gap-y-3 mt-5">
                  {Object.entries(item.fields).map(
                    ([fieldName, isActive], index) => (
                      <div
                        key={index}
                        className="flex flex-row w-full items-center justify-between text-black font-bold"
                      >
                        {fieldName}
                        {fieldName !== "Connected" ? (
                          <Switch
                            className="data-[state=checked]:bg-green-400 data-[state=unchecked]:bg-red-500"
                            checked={isActive}
                            onClick={() =>
                              handleSwitchChange(
                                key,
                                fieldName,
                                item.name.toLowerCase(),
                              )
                            }
                          />
                        ) : (
                          <Dialog>
                            <DialogTrigger asChild>
                              <Button variant="outline">
                                {isActive ? "Disconect" : "Connect"}
                              </Button>
                            </DialogTrigger>
                            <DialogContent className="">
                              <DialogHeader>
                                <DialogTitle className="text-2xl">
                                  {isActive ? "Disconect" : "Connect"}{" "}
                                  {item.name}
                                </DialogTitle>
                                <DialogDescription className="text-xl">
                                  Are you sure you want to{" "}
                                  {isActive ? "disconect" : "connect"}{" "}
                                  {item.name}?
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
                                        isActive,
                                        key,
                                        fieldName,
                                        item.name.toLowerCase(),
                                      )
                                    }
                                  >
                                    {isActive ? "Disconect" : "Connect"}
                                  </Button>
                                </DialogClose>
                              </DialogFooter>
                            </DialogContent>
                          </Dialog>
                        )}
                      </div>
                    ),
                  )}
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </ScrollArea>
    </div>
  );
}
