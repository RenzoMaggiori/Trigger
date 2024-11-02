package errors

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	// Not Found Errors
	ErrWorkspaceNotFound           error = errors.New("workspace not found")
	ErrActionNotFound              error = errors.New("action not found")
	ErrSessionNotFound             error = errors.New("session not found")
	ErrUserNotFound                error = errors.New("user not found")
	ErrGithubStopModelNotFound     error = errors.New("github stop model not found")
	ErrNodeNotFound                error = errors.New("action node not found")
	ErrAuthorizationHeaderNotFound error = errors.New("authorization header not found")
	ErrSyncNotFound                error = errors.New("sync not found")
	ErrCollectionNotFound          error = errors.New("collection not found")

	// Bad Request Errors
	ErrBadWorkspaceId        error = errors.New("bad workspace id")
	ErrBadUserId             error = errors.New("bad user id")
	ErrBadActionId           error = errors.New("bad action id")
	ErrSessionTypeNone       error = errors.New("could not decypher session type")
	ErrInvalidGithubStatus   error = errors.New("invalid github status")
	ErrInvalidReactionInput  error = errors.New("invalid reaction input")
	ErrInvalidReactionOutput error = errors.New("invalid reaction output")
	ErrNoTokenInRequest      error = errors.New("token could not be found in the request")

	// Decyphering Errors
	ErrActionTypeNone       error = errors.New("could not decypher action type")
	ErrActionNodeTypeNone   error = errors.New("could not decypher action node type")
	ErrUserTypeNone         error = errors.New("could not decypher user type")
	ErrGmailHistoryTypeNone error = errors.New("could not decypher gmail history type")
	ErrHistoryInt           error = errors.New("error converting historyto a number")

	// Context Errors
	ErrAccessTokenCtx error = errors.New("could not retrieve access token from context")
	ErrEventCtx       error = errors.New("could not retrieve event")

	// Creation Errors
	ErrCreatingWorkspace error = errors.New("error while creating workspace")
	ErrCreatingSession   error = errors.New("error while creating session")
	ErrCreatingEmail     error = errors.New("failed to create raw email")
	ErrCreatingSync      error = errors.New("failed to create a new sync")
	ErrCreatingSetting   error = errors.New("failed to create a new setting")
	ErrCreatingNode      error = errors.New("failed to create workspace node")

	// Updating Errors
	ErrUpdatingWorkspace error = errors.New("error while updating workspace")

	// Setting/Completing Errors
	ErrSettingAction         error = errors.New("error while setting trigger or reaction")
	ErrCompletingAction      error = errors.New("error while completing action")
	ErrCompletingWatchAction error = errors.New("error while completing watch action")
	ErrStopingWorkspace      error = errors.New("error while stopping workspace")
	// Retrieval/Fetching Errors
	ErrFetchingSession        error = errors.New("error while retrieving session")
	ErrFetchingActions        error = errors.New("error while retrieving actions")
	ErrFetchingSync           error = errors.New("error while retrieving sync")
	ErrActionProviderNotFound error = errors.New("could not find action provider")

	// Email Errors
	ErrFailedToSendEmail  error = errors.New("failed to send email")
	ErrGmailSendEmail     error = errors.New("error while sending email through gmail")
	ErrGmailWatch         error = errors.New("error while watching gmail")
	ErrGmailStop          error = errors.New("error while stopping gmail")
	ErrGmailHistory       error = errors.New("error while fetching gmail history")
	ErrInvalidGoogleToken error = errors.New("token provided is not valid")

	// Twitch Errors
	ErrTwitchUser             error = errors.New("error while fetching twitch user")
	ErrTwitchUserFound        error = errors.New("twitch user found")
	ErrTwitchAppAccessToken   error = errors.New("error while fetching twitch app access token")
	ErrTwitchSendMessage      error = errors.New("error while sending twitch channel message")
	ErrWebhookVerificationCtx error = errors.New("could not find webhook verification in ctx")
	ErrTwitchWatch            error = errors.New("error while watching twitch")

	// Sync Errors
	ErrSyncAccessTokenNotFound error = errors.New("error could not find sync access token")
	ErrSyncModelTypeNone       error = errors.New("error could not decode sync model")

	// Spotify Errors
	ErrSpotifyBadStatus error = errors.New("invalid response status from spotify")

	// Webhook
	ErrBadWebhookData error = errors.New("could not parse the webhook data")

	// DATA
	ErrMarshalData error = errors.New("could not marshal data")
	ErrDecodeData  error = errors.New("could not decode data")

	// Discord
	ErrDiscordGuilds              error = errors.New("could not retrieve guilds")
	ErrDiscordMe                  error = errors.New("could not retrieve user/@me data")
	ErrCreateDiscordGoSession     error = errors.New("could not create discord session")
	ErrOpeningDiscordConnection   error = errors.New("error opening discord connection")
	ErrBotAlreadyRunning          error = errors.New("bot is already running for this user")
	ErrBotNotRunning              error = errors.New("bot is not running for this user")
	ErrDiscordUserSessionNotFound error = errors.New("discord user session not found")
	ErrGuildIdNotFound            error = errors.New("guild id not found")
	ErrAddDiscordSession          error = errors.New("error storing discord session in db")
	ErrUpdateDiscordSession       error = errors.New("error updating discord session in db")
	ErrDeleteDiscordSession       error = errors.New("error deleting discord session in db")

	// State
	ErrMalformedState error = errors.New("malformed state")

	//Github
	ErrGithubUserInfo       error = errors.New("could not get github user info")
	ErrGithubCommitInfo     error = errors.New("could not get github commit info")
	ErrGithubSendingWebhook error = errors.New("error while sending github webhook")
	ErrGithubCommitData     error = errors.New("error getting github commit data")

	ErrCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		ErrWorkspaceNotFound: {
			Message: ErrWorkspaceNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrNodeNotFound: {
			Message: ErrNodeNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrBadWorkspaceId: {
			Message: ErrBadWorkspaceId.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrBadUserId: {
			Message: ErrBadUserId.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrBadActionId: {
			Message: ErrBadActionId.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrSessionNotFound: {
			Message: ErrSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrSessionTypeNone: {
			Message: ErrSessionTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		ErrCreatingWorkspace: {
			Message: ErrCreatingWorkspace.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCreatingNode: {
			Message: ErrCreatingNode.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrFetchingActions: {
			Message: ErrFetchingActions.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrActionTypeNone: {
			Message: ErrActionTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		ErrActionNodeTypeNone: {
			Message: ErrActionNodeTypeNone.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrUserNotFound: {
			Message: ErrUserNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrUserTypeNone: {
			Message: ErrUserTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		ErrSettingAction: {
			Message: ErrSettingAction.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCompletingAction: {
			Message: ErrCompletingAction.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCompletingWatchAction: {
			Message: ErrCompletingWatchAction.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrActionNotFound: {
			Message: ErrActionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrFetchingSession: {
			Message: ErrFetchingSession.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCreatingSession: {
			Message: ErrCreatingSession.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrAccessTokenCtx: {
			Message: ErrAccessTokenCtx.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCreatingEmail: {
			Message: ErrCreatingEmail.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrFailedToSendEmail: {
			Message: ErrFailedToSendEmail.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGmailSendEmail: {
			Message: ErrGmailSendEmail.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGmailWatch: {
			Message: ErrGmailWatch.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGmailStop: {
			Message: ErrGmailStop.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGmailHistory: {
			Message: ErrGmailHistory.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGmailHistoryTypeNone: {
			Message: ErrGmailHistoryTypeNone.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGithubStopModelNotFound: {
			Message: ErrGithubStopModelNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrInvalidGithubStatus: {
			Message: ErrInvalidGithubStatus.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrInvalidReactionInput: {
			Message: ErrInvalidReactionInput.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrInvalidReactionOutput: {
			Message: ErrInvalidReactionOutput.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrHistoryInt: {
			Message: ErrHistoryInt.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrInvalidGoogleToken: {
			Message: ErrInvalidGoogleToken.Error(),
			Code:    http.StatusUnauthorized,
		},
		ErrAuthorizationHeaderNotFound: {
			Message: ErrAuthorizationHeaderNotFound.Error(),
			Code:    http.StatusForbidden,
		},
		ErrUpdatingWorkspace: {
			Message: ErrUpdatingWorkspace.Error(),
			Code:    http.StatusNotFound,
		},
		ErrTwitchUser: {
			Message: ErrTwitchUser.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrTwitchUserFound: {
			Message: ErrTwitchUserFound.Error(),
			Code:    http.StatusOK,
		},
		ErrTwitchAppAccessToken: {
			Message: ErrTwitchAppAccessToken.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrTwitchSendMessage: {
			Message: ErrTwitchSendMessage.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrWebhookVerificationCtx: {
			Message: ErrWebhookVerificationCtx.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrSyncAccessTokenNotFound: {
			Message: ErrUpdatingWorkspace.Error(),
			Code:    http.StatusNotFound,
		},
		ErrSyncModelTypeNone: {
			Message: ErrSyncModelTypeNone.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrSpotifyBadStatus: {
			Message: ErrSpotifyBadStatus.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrBadWebhookData: {
			Message: ErrBadWebhookData.Error(),
			Code:    http.StatusUnprocessableEntity,
		},
		ErrMalformedState: {
			Message: ErrMalformedState.Error(),
			Code:    http.StatusUnprocessableEntity,
		},
		ErrActionProviderNotFound: {
			Message: ErrActionProviderNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrSyncNotFound: {
			Message: ErrSyncNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrCreatingSync: {
			Message: ErrCreatingSync.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrFetchingSync: {
			Message: ErrFetchingSync.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrMarshalData: {
			Message: ErrMarshalData.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCreatingSetting: {
			Message: ErrCreatingSetting.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrMarshalData: {
			Message: ErrMarshalData.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrDecodeData: {
			Message: ErrDecodeData.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrDiscordGuilds: {
			Message: ErrDiscordGuilds.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrDiscordMe: {
			Message: ErrDiscordMe.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGuildIdNotFound: {
			Message: ErrGuildIdNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrCreateDiscordGoSession: {
			Message: ErrCreateDiscordGoSession.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrOpeningDiscordConnection: {
			Message: ErrOpeningDiscordConnection.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrBotAlreadyRunning: {
			Message: ErrBotAlreadyRunning.Error(),
			Code:    http.StatusConflict,
		},
		ErrBotNotRunning: {
			Message: ErrBotNotRunning.Error(),
			Code:    http.StatusConflict,
		},
		ErrDiscordUserSessionNotFound: {
			Message: ErrDiscordUserSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrAddDiscordSession: {
			Message: ErrAddDiscordSession.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrUpdateDiscordSession: {
			Message: ErrUpdateDiscordSession.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrDeleteDiscordSession: {
			Message: ErrDeleteDiscordSession.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCollectionNotFound: {
			Message: ErrCollectionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrGithubUserInfo: {
			Message: ErrGithubUserInfo.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGithubCommitInfo: {
			Message: ErrGithubCommitInfo.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGithubSendingWebhook: {
			Message: ErrGithubSendingWebhook.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrGithubCommitData: {
			Message: ErrGithubCommitData.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrStopingWorkspace: {
			Message: ErrStopingWorkspace.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrTwitchWatch: {
			Message: ErrTwitchWatch.Error(),
			Code:    http.StatusInternalServerError,
		},
	}
)
