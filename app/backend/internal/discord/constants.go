package discord

type StopCtx string
type WorkspaceCtx string

const StopCtxKey = StopCtx("StopCtxKey")

const WorkspaceCtxKey WorkspaceCtx = WorkspaceCtx("WorkspaceCtxKey")

const AccessTokenCtxKey WorkspaceCtx = WorkspaceCtx("AuthorizationCtxKey")

const DiscordEventCtxKey WorkspaceCtx = WorkspaceCtx("DiscordEventCtxKey")