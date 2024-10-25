"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import React from "react";
import { useMenu } from "@/app/trigger/components/MenuProvider";
import { NodeItem } from "@/app/trigger/lib/types";
import { Combox, Status } from "@/components/ui/combox";

const InputComponent = ({
  label,
  placeholder,
  type,
}: {
  label: string;
  placeholder?: string;
  type?: string;
}) => {
  const inputType = type ? type : undefined;
  const inputPlaceholder = placeholder ? placeholder : undefined;

  return (
    <div>
      <Label>{label}</Label>
      <Input placeholder={inputPlaceholder} type={inputType} />
    </div>
  );
};

function GithubSettings({}: { node: NodeItem, type: string }) {
  return <div></div>;
}

function EmailSettings({ node, type }: { node: NodeItem, type: string }) {
  const { setFields } = useMenu();

  const handleFieldChange = (type: string, index: string, value: string) => {
    const currentField = node.fields["type"];

    if (currentField !== null && type === currentField) {
      setFields(node.id, { ...node.fields, [index]: value });
    } else {
      setFields(node.id, { ["type"]: type, [index]: value });
    }
  };

  const inputs = [
    { label: "Destination", placeholder: "example@example.com", type: "email" },
    { label: "Title", placeholder: "Example title..." },
    { label: "Subject", placeholder: "Example subject..." },
  ];

  if (!node) return <div>No node found</div>;

  return (
    <>
      {type === "reaction" ? (
        <div className="flex flex-col gap-y-4">
          {inputs.map((item, key) => (
            <div key={`${node.id}-${key}`}>
              <Label>{item.label}</Label>
              <Input
                placeholder={item.placeholder}
                onChange={(e) => handleFieldChange(type, item.label, e.target.value)}
                value={ node.fields["type"] !== null
                  ? (node.fields[item.label] as string | number | undefined) || ""
                  : ""}
                type={item.type}
              />
            </div>
          ))}
          <div>
            <Label>Email body</Label>
            <Textarea
              placeholder="Example body..."
              className="resize-none h-[200px]"
              onChange={(e) => handleFieldChange(type, "Body", e.target.value)}
              value={node.fields["type"] !== null
                ? (node.fields["Body"] as string | number | undefined) || ""
                : ""}
            />
          </div>
        </div>
      ) : (
        <div className="flex flex-col gap-y-4">
          <Label>Source</Label>
          <Input
            placeholder="example@example.com"
            onChange={(e) => handleFieldChange(type, "Source", e.target.value)}
            value={node.fields["type"] !== null
              ? (node.fields["Source"] as string | number | undefined) || ""
              : ""}
          />
        </div>
      )}
    </>
  );
}


function DiscordSettings({}: { node: NodeItem, type: string }) {
  const [messageType, setMessageType] = React.useState<Status | null>({ label: "Normal Message", value: "Normal" });
  const [embedFields, setEmbedFields] = React.useState<
    { name: string; value: string }[]
  >([]);

  const handleAddField = () => {
    setEmbedFields([...embedFields, { name: "", value: "" }]);
  };

  const handleFieldChange = (
    index: number,
    fieldType: "name" | "value",
    value: string,
  ) => {
    const updatedFields = [...embedFields];
    updatedFields[index][fieldType] = value;
    setEmbedFields(updatedFields);
  };

  const handleRemoveField = (index: number) => {
    const updatedFields = embedFields.filter((_, i) => i !== index);
    setEmbedFields(updatedFields);
  };

  const inputs: {
    label: string;
    placeholder?: string;
    type?: string;
  }[] = [
      { label: "Embed Color", placeholder: "Example title...", type: "color" },
      { label: "Embed Title", placeholder: "Example embed title" },
    ];

  const fieldInputs = [
    { placeholder: "Field Name", fieldType: "name" },
    { placeholder: "Field Value", fieldType: "value" },
  ];

  return (
    <>
      <div className="mb-4">
        <Label
          htmlFor="message-type-dropdown"
          className="block text-sm font-medium text-gray-700"
        >
          Select Message Type
        </Label>
        <Combox
          statuses={[{ label: "Normal Message", value: "Normal" }, { label: "Embeded Message", value: "Embed" }]}
          selectedStatus={messageType}
          setSelectedStatus={setMessageType}
          label="Select Message Type"
        />
      </div>

      {messageType?.value === "Normal" ? (
        <div className="normal-message-settings">
          <Label
            htmlFor="normal-message-content"
            className="block text-sm font-medium text-gray-700"
          >
            Message Content
          </Label>
          <Textarea
            id="normal-message-content"
            placeholder="Enter your message content here..."
            className="mt-1 block w-full border-gray-300 rounded-md shadow-sm sm:text-sm resize-none h-[150px]"
          />
        </div>
      ) : (
        <div className="embed-message-settings">
          {inputs.map((item, key) => (
            <div key={key}>
              <InputComponent
                label={item.label}
                placeholder={item.placeholder}
                type={item.type}
              />
            </div>
          ))}

          <Label
            htmlFor="embed-description"
            className="block text-sm font-medium text-gray-700 mt-4"
          >
            Embed Description
          </Label>
          <Textarea
            id="embed-description"
            placeholder="Enter embed description"
            className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
          <div className="mt-4">
            <h4 className="font-bold text-lg">Custom Fields</h4>
            {embedFields.map((field, index) => (
              <div key={index} className="flex gap-4 items-center mb-2">
                {fieldInputs.map((item, key) => (
                  <Input
                    key={key}
                    placeholder={item.placeholder}
                    onChange={(e) =>
                      handleFieldChange(
                        index,
                        item.fieldType === "name" ? "name" : "value",
                        e.target.value,
                      )
                    }
                    className="w-1/2"
                  />
                ))}
                <Button
                  variant="destructive"
                  onClick={() => handleRemoveField(index)}
                  className="ml-2"
                >
                  Remove
                </Button>
              </div>
            ))}
            <Button variant="outline" onClick={handleAddField} className="mt-2">
              Add Field
            </Button>
          </div>
        </div>
      )}
    </>
  );
}

export { EmailSettings, DiscordSettings, GithubSettings };

