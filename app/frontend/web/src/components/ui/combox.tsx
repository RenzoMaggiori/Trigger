import React from "react";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";

type Status = {
  label: string | React.ReactNode;
  value: string;
};

function StatusList({
  setOpen,
  setSelectedStatus,
  statuses,
}: {
  setOpen: (open: boolean) => void;
  setSelectedStatus: (status: Status | null) => void;
  statuses: Status[];
}) {
  return (
    <Command>
      <CommandInput placeholder="Filter status..." />
      <CommandList>
        <CommandEmpty>No results found.</CommandEmpty>
        <CommandGroup>
          {statuses.map((status) => (
            <CommandItem
              key={status.value}
              value={status.value}
              onSelect={(value) => {
                setSelectedStatus(
                  statuses.find((priority) => priority.value === value) || null,
                );
                setOpen(false);
              }}
            >
              {status.label}
            </CommandItem>
          ))}
        </CommandGroup>
      </CommandList>
    </Command>
  );
}

interface ComboxProps {
  statuses: Status[];
  selectedStatus: Status | null;
  setSelectedStatus: (status: Status | null) => void;
  label: string;
  icon?: React.ReactNode;
}

const Combox: React.FC<ComboxProps> = ({
  statuses,
  setSelectedStatus,
  selectedStatus,
  label,
  icon,
}) => {
  const [open, setOpen] = React.useState(false);
  const Icon = icon ? icon : <></>;
  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button variant="outline" className="w-[150px] justify-start">
          {selectedStatus ? (
            <>{selectedStatus.label}</>
          ) : (
            <div className="flex flex-row items-center justify-start text-center font-bold text-lg">
              {Icon}
              {label}
            </div>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0" align="start">
        <StatusList
          setOpen={setOpen}
          setSelectedStatus={setSelectedStatus}
          statuses={statuses}
        />
      </PopoverContent>
    </Popover>
  );
};

export { Combox, StatusList, type Status };

