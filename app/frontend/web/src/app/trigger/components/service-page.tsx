"use client";
import React from "react";
import { Card, CardContent } from "@/components/ui/card";
import { TriggerDraggable } from "@/app/trigger/components/trigger-draggable";
import { Service } from "@/app/trigger/lib/types";
import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-react";

interface ServicesProps {
  services: Service[];
  handleDragStart: (
    e: React.DragEvent<HTMLDivElement>,
    service: Service,
  ) => void;
  handleOnClick: () => void;
}

export const ServicesComponent: React.FC<ServicesProps> = ({
  services,
  handleDragStart,
  handleOnClick,
}) => {
  const [loading, setLoading] = React.useState<boolean>(false);

  return (
    <div className="w-auto p-5">
      <Card className="h-full">
        <p className="font-bold text-2xl p-3">Services</p>
        <CardContent className="flex flex-col items-center justify-start h-full py-5 gap-4">
          {services.map((item, key) => (
            <div
              key={key}
              draggable
              onDragStart={(e) => handleDragStart(e, item)}
              className="cursor-move"
            >
              <TriggerDraggable service={item} className="w-[200px]" />
            </div>
          ))}
          <Button
            className="w-full text-md rounded-full bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 hover:bg-gradient-to-r hover:from-blue-600 hover:via-violet-600 hover:to-fuchsia-600 animate-gradient text-white"
            onClick={() => {
              handleOnClick();
              setLoading(true);
            }}
            disabled={loading}
          >
            {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            Deploy Trigger
          </Button>
        </CardContent>
      </Card>
    </div>
  );
};
