"use client";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import React from "react";
import { useMenu } from "@/app/trigger/components/MenuProvider";
import { ActionType, NodeItem } from "@/app/trigger/lib/types";
import { Combox, Status } from "@/components/ui/combox";
import { BsHourglassTop, BsHourglassBottom, BsHourglassSplit } from "react-icons/bs";
import { AiOutlineIssuesClose } from "react-icons/ai";
import { FaCodeCommit } from "react-icons/fa6";
import { FaCodePullRequest } from "react-icons/fa6";

function GithubSettings({ node, type, actions }: { node: NodeItem, type: string, actions: ActionType }) {
  const { setFields, setNodes } = useMenu();

  React.useEffect(() => {
    const githubTriggerAction = actions.find(
      (action) => action.provider === "github" && action.type === type
    );
    if (!githubTriggerAction) return;
    if (node.action_id !== githubTriggerAction.id) {
      setNodes({
        [node.id]: { ...node, action_id: githubTriggerAction.id },
      });
    }
    if (node.fields?.type !== type)
      setFields(node.id, { ...node.fields, type });
  }, [type, actions, node, setNodes, setFields]);

  const handleFieldChange = (fieldType: string, index: string, value: string) => {
    setFields(node.id, { ...node.fields, [index]: value, type: fieldType });
  };

  const triggerInputs = [
    { label: "Owner", json: "owner", placeholder: "John Doe" },
    { label: "Repository", json: "repo", placeholder: "example_repository" },
  ];

  const reactionInputs = [
    { label: "Owner", json: "owner", placeholder: "John Doe" },
    { label: "Repository", json: "repo", placeholder: "example_repository" },
    { label: "Title", json: "title", placeholder: "Example title" },
    { label: "Description", json: "description", placeholder: "This is a new issue" },
  ];

  return (
    <div>
      {type === "reaction" ? (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Creates a new issue on the specified repository.</p>
          {reactionInputs.map((item, key) => (
            <div key={`${node.id}-${key}`}>
              <Label>{item.label}</Label>
              <Input
                placeholder={item.placeholder}
                onChange={(e) => handleFieldChange(type, item.json, e.target.value)}
                value={node.fields["type"] !== null
                  ? (node.fields[item.json] as string | number | undefined) || ""
                  : ""}
              />
            </div>
          ))}
        </div>
      ) : (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Waits for a push to happen.</p>
          {triggerInputs.map((item, key) => (
            <div key={`${node.id}-${key}`}>
              <Label>{item.label}</Label>
              <Input
                placeholder={item.placeholder}
                onChange={(e) => handleFieldChange(type, item.json, e.target.value)}
                value={node.fields["type"] !== null
                  ? (node.fields[item.json] as string | number | undefined) || ""
                  : ""}
              />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

function TwitchSettings({ node, type, actions }: { node: NodeItem, type: string, actions: ActionType }) {
  const { setFields, setNodes } = useMenu();

  React.useEffect(() => {
    const twitchTriggerAction = actions.find(
      (action) => action.provider === "twitch" && action.type === type
    );
    if (!twitchTriggerAction) return;
    if (node.action_id !== twitchTriggerAction.id) {
      setNodes({
        [node.id]: { ...node, action_id: twitchTriggerAction.id },
      });
    }
    if (node.fields?.type !== type)
      setFields(node.id, { ...node.fields, type });
  }, [type, actions, node, setNodes, setFields]);

  const handleFieldChange = (fieldType: string, index: string, value: string) => {
    setFields(node.id, { ...node.fields, [index]: value, type: fieldType });
  };

  if (!node) return <div>No node found</div>;

  return (
    <>
      {type === "reaction" ? (
        <div className="flex flex-col gap-y-4">
          <Label>Message to send</Label>
          <Input
            placeholder="your followers count increased!!"
            onChange={(e) => handleFieldChange(type, "message", e.target.value)}
            value={node.fields["type"] !== null
              ? (node.fields["message"] as string | number | undefined) || ""
              : ""}
          />
        </div>
      ) : (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Waits for the follower count to increase.</p>
        </div>
      )}
    </>
  );
}

function EmailSettings({ node, type, actions }: { node: NodeItem, type: string, actions: ActionType }) {
  const { setFields, setNodes } = useMenu();

  React.useEffect(() => {
    const gmailTriggerAction = actions.find(
      (action) => action.provider === "gmail" && action.type === type
    );

    if (!gmailTriggerAction) return;

    if (node.action_id !== gmailTriggerAction.id) {
      setNodes({
        [node.id]: { ...node, action_id: gmailTriggerAction.id },
      });
    }

    if (node.fields?.type !== type)
      setFields(node.id, { ...node.fields, type });
  }, [type, actions, node, setNodes, setFields]);

  const handleFieldChange = (fieldType: string, index: string, value: string) => {
    setFields(node.id, { ...node.fields, [index]: value, type: fieldType });
  };

  const inputs = [
    { label: "Destination", json: "to", placeholder: "example@example.com", type: "email" },
    { label: "Subject", json: "subject", placeholder: "This is an email subject" },
  ];

  if (!node) return <div>No node found</div>;

  return (
    <>
      {type === "reaction" ? (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Sends an email to the desired destination.</p>
          {inputs.map((item, key) => (
            <div key={`${node.id}-${key}`}>
              <Label>{item.label}</Label>
              <Input
                placeholder={item.placeholder}
                onChange={(e) => handleFieldChange(type, item.json, e.target.value)}
                value={(node.fields[item.json] as string | number | undefined) || ""}
                type={item.type}
              />
            </div>
          ))}
          <div>
            <Label>Email body</Label>
            <Textarea
              placeholder="Hey there! Just wanted to check in and see how you’re doing..."
              className="resize-none h-[200px]"
              onChange={(e) => handleFieldChange(type, "body", e.target.value)}
              value={(node.fields["body"] as string | number | undefined) || ""}
            />
          </div>
        </div>
      ) : (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Waits for a gmail action to happen {"(email arrived, sent, deleted, etc)"}.</p>
        </div>
      )}
    </>
  );
}


function SpotifySettings({ node, type, actions }: { node: NodeItem, type: string, actions: ActionType }) {
  const { setFields, setNodes } = useMenu();

  React.useEffect(() => {
    const spotifyTriggerAction = actions.find(
      (action) => action.provider === "spotify" && action.type === type
    );
    if (!spotifyTriggerAction) return;
    if (node.action_id !== spotifyTriggerAction.id) {
      setNodes({
        [node.id]: { ...node, action_id: spotifyTriggerAction.id },
      });
    }
    if (node.fields?.type !== type)
      setFields(node.id, { ["type"]: type });
  }, [type, actions, node, setNodes, setFields]);


  if (!node) return <div>No node found</div>;


  return (
    <>
      {type === "reaction" ? (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Plays music on your device.</p>
        </div>
      ) : (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Waits for the follower count to change.</p>
        </div>
      )}
    </>
  );
}


function DiscordSettings({ node, type, actions }: { node: NodeItem, type: string, actions: ActionType }) {
  const { setFields, setNodes } = useMenu();

  React.useEffect(() => {
    const discordTriggerAction = actions.find(
      (action) => action.provider === "discord" && action.type === type
    );

    if (!discordTriggerAction) return;

    if (node.action_id !== discordTriggerAction.id) {
      setNodes({
        [node.id]: { ...node, action_id: discordTriggerAction.id },
      });
    }

    if (node.fields?.type !== type)
      setFields(node.id, { ...node.fields, type });
  }, [type, actions, node, setNodes, setFields]);

  const handleFieldChange = (fieldType: string, index: string, value: string) => {
    setFields(node.id, { ...node.fields, [index]: value, type: fieldType });
  };

  if (!node) return <div>No node found</div>;

  return (
    <>
      {type === "reaction" ? (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Sends an message to the desired discord channel.</p>
          <Label>Channel ID</Label>
          <Input
            placeholder="1234567890"
            onChange={(e) => handleFieldChange(type, "channel_id", e.target.value)}
            value={(node.fields["channel_id"] as string | number | undefined) || ""}
          />
          <div>
            <Label>Message to send</Label>
            <Textarea
              placeholder="This is an example mesagge"
              className="resize-none h-[200px]"
              onChange={(e) => handleFieldChange(type, "content", e.target.value)}
              value={(node.fields["content"] as string | number | undefined) || ""}
            />
          </div>
        </div>
      ) : (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Waits for a message to be sent on a discord channel.</p>
          <Label>Channel ID</Label>
          <Input
            placeholder="1234567890"
            onChange={(e) => handleFieldChange(type, "channel_id", e.target.value)}
            value={(node.fields["channel_id"] as string | number | undefined) || ""}
          />
        </div>
      )}
    </>
  );
}

const bitbucketStatuses = [
  {
    label:
      <div className="flex flex-row">
        <AiOutlineIssuesClose className="mr-2 w-5 h-5" />
        <p className="font-bold">New issue</p>
      </div>,
    value: "watch_issue_created"
  },
  {
    label:
      <div className="flex flex-row">
        <FaCodeCommit className="mr-2 w-5 h-5" />
        <p className="font-bold">New commit</p>
      </div>,
    value: "watch_repo_push"
  },
  {
    label:
      <div className="flex flex-row">
        <FaCodePullRequest className="mr-2 w-5 h-5" />
        <p className="font-bold">Pull request</p>
      </div>,
    value: "watch_pull_request_created"
  },
];

function BitBucketSettings({ node, type, actions }: { node: NodeItem, type: string, actions: ActionType }) {
  const [messageType, setMessageType] = React.useState<Status | null>(bitbucketStatuses[0]);
  const { setFields, setNodes } = useMenu();

  React.useEffect(() => {
    const bitbucketTriggerAction = actions.find(
      (action) => action.provider === "bitbucket" && action.type === type && (action.type === "trigger" ? action.action === messageType?.value : true)
    );

    if (!bitbucketTriggerAction) return;

    if (node.action_id !== bitbucketTriggerAction.id) {
      setNodes({
        [node.id]: { ...node, action_id: bitbucketTriggerAction.id },
      });
    }

    if (node.fields?.type !== type)
      setFields(node.id, { ...node.fields, type });
  }, [type, actions, node, setNodes, setFields, messageType]);

  const handleFieldChange = (fieldType: string, index: string, value: string) => {
    setFields(node.id, { ...node.fields, [index]: value, type: fieldType });
  };

  const inputs = [
    { label: "Workspace", json: "workspace", placeholder: "Example workspace" },
    { label: "Repository", json: "repository", placeholder: "Example repo" },
    { label: "Title", json: "title", placeholder: "Example title" },
    { label: "Source Branch", json: "source_branch", placeholder: "Example branch" },
    { label: "Destination Branch", json: "destination_branch", placeholder: "Example branch" },
  ];

  const inputsTrigger = [
    { label: "Workspace", json: "workspace", placeholder: "Example workspace" },
    { label: "Repository", json: "repository", placeholder: "Example repo" },
  ];

  if (!node) return <div>No node found</div>;

  return (
    <>
      {type === "reaction" ? (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Creates a pull request.</p>
          {inputs.map((item, key) => (
            <div key={`${node.id}-${key}`}>
              <Label>{item.label}</Label>
              <Input
                placeholder={item.placeholder}
                onChange={(e) => handleFieldChange(type, item.json, e.target.value)}
                value={(node.fields[item.json] as string | number | undefined) || ""}
              />
            </div>
          ))}
        </div>
      ) : (
        <div className="flex flex-col gap-y-4">
          <p className="text-zinc-500">Waits for a commit to happen.</p>
          <Combox
            statuses={bitbucketStatuses}
            selectedStatus={messageType}
            setSelectedStatus={setMessageType}
            label="Select type trigger"
          />
          {inputsTrigger.map((item, key) => (
            <div key={`${node.id}-${key}`}>
              <Label>{item.label}</Label>
              <Input
                placeholder={item.placeholder}
                onChange={(e) => handleFieldChange(type, item.json, e.target.value)}
                value={(node.fields[item.json] as string | number | undefined) || ""}
              />
            </div>
          ))}
        </div>
      )}
    </>
  );
}

const timerStatuses = [
  {
    label:
      <div className="flex flex-row">
        <BsHourglassTop className="mr-2 w-5 h-5" />
        <p className="font-bold">Minute</p>
      </div>,
    value: "watch_minute"
  },
  {
    label:
      <div className="flex flex-row">
        <BsHourglassSplit className="mr-2 w-5 h-5" />
        <p className="font-bold">Hour</p>
      </div>,
    value: "watch_hour"
  },
  {
    label:
      <div className="flex flex-row">
        <BsHourglassBottom className="mr-2 w-5 h-5" />
        <p className="font-bold">Day</p>
      </div>,
    value: "watch_day"
  },
];

function TimerSettings({ node, type, actions }: { node: NodeItem, type: string, actions: ActionType }) {
  const [messageType, setMessageType] = React.useState<Status | null>(timerStatuses[0]);
  const { setNodes } = useMenu();

  React.useEffect(() => {
    const timmerTriggerAction = actions.find(
      (action) => action.provider === "timer" && action.type === "trigger" && action.action === messageType?.value
    );
    if (!timmerTriggerAction) return;
    if (node.action_id !== timmerTriggerAction.id) {
      setNodes({
        [node.id]: { ...node, action_id: timmerTriggerAction.id },
      });
    }
  }, [type, actions, node, setNodes, messageType]);


  const descriptions = {
    "watch_minute": "Triggers each time the minute changes.",
    "watch_hour": "Triggers each time the hour changes.",
    "watch_day": "Triggers each time the day changes.",
  }

  if (!node) return <div>No node found</div>;

  return (
    <div>
      <div className="mb-4">
        <Label
          htmlFor="message-type-dropdown"
          className="block text-sm font-medium text-gray-700"
        >
          Select Message Type
        </Label>
        <Combox
          statuses={timerStatuses}
          selectedStatus={messageType}
          setSelectedStatus={setMessageType}
          label="Select Time for trigger"
        />
      </div>
      <p className="text-zinc-500">{descriptions[messageType?.value as keyof typeof descriptions || "watch_minute"]}</p>
    </div>
  );
}

export { EmailSettings, DiscordSettings, GithubSettings, SpotifySettings, TwitchSettings, TimerSettings, BitBucketSettings };

