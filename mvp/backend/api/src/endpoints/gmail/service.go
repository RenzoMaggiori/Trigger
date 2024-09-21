package gmail

import (
	"context"
	"errors"
)

var _ Gmail = Model{}

/* "auth": {
  "oauth2": {
    "scopes": {
      "https://mail.google.com/": {
        "description": "Read, compose, send, and permanently delete all your email from Gmail"
      },
      "https://www.googleapis.com/auth/gmail.addons.current.action.compose": {
        "description": "Manage drafts and send emails when you interact with the add-on"
      },
      "https://www.googleapis.com/auth/gmail.addons.current.message.action": {
        "description": "View your email messages when you interact with the add-on"
      },
      "https://www.googleapis.com/auth/gmail.addons.current.message.metadata": {
        "description": "View your email message metadata when the add-on is running"
      },
      "https://www.googleapis.com/auth/gmail.addons.current.message.readonly": {
        "description": "View your email messages when the add-on is running"
      },
      "https://www.googleapis.com/auth/gmail.compose": {
        "description": "Manage drafts and send emails"
      },
      "https://www.googleapis.com/auth/gmail.insert": {
        "description": "Add emails into your Gmail mailbox"
      },
      "https://www.googleapis.com/auth/gmail.labels": {
        "description": "See and edit your email labels"
      },
      "https://www.googleapis.com/auth/gmail.metadata": {
        "description": "View your email message metadata such as labels and headers, but not the email body"
      },
      "https://www.googleapis.com/auth/gmail.modify": {
        "description": "Read, compose, and send emails from your Gmail account"
      },
      "https://www.googleapis.com/auth/gmail.readonly": {
        "description": "View your email messages and settings"
      },
      "https://www.googleapis.com/auth/gmail.send": {
        "description": "Send email on your behalf"
      },
      "https://www.googleapis.com/auth/gmail.settings.basic": {
        "description": "See, edit, create, or change your email settings and filters in Gmail"
      },
      "https://www.googleapis.com/auth/gmail.settings.sharing": {
        "description": "Manage your sensitive mail settings, including who can manage your mail"
      }
    }
  }
} */

func (m Model) Auth() error {
	// POST /gmail/v1/users/{userId}/watch
	// https://gmail.googleapis.com
	return nil
}

/* "watch": {
"id": "gmail.users.watch",
"path": "gmail/v1/users/{userId}/watch",
"flatPath": "gmail/v1/users/{userId}/watch",
"httpMethod": "POST",
"parameters": {
"userId": {
"description": "The user's email address. The special value `me` can be used to indicate the authenticated user.",
"default": "me",
"location": "path",
"required": true,
"type": "string"
}
},
"parameterOrder": [
"userId"
],
"request": {
"$ref": "WatchRequest"
},
"response": {
"$ref": "WatchResponse"
},
"scopes": [
"https://mail.google.com/",
"https://www.googleapis.com/auth/gmail.metadata",
"https://www.googleapis.com/auth/gmail.modify",
"https://www.googleapis.com/auth/gmail.readonly"
],
"description": "Set up or update a push notification watch on the given user mailbox."
}, */

/* "WatchRequest": {
"id": "WatchRequest",
"description": "Set up or update a new push notification watch on this user's mailbox.",
"type": "object",
"properties": {
"labelIds": {3 items},
"labelFilterAction": {5 items},
"labelFilterBehavior": {4 items},
"topicName": {2 items}
}
}, */

/* "WatchResponse": {
"id": "WatchResponse",
"description": "Push notification watch response.",
"type": "object",
"properties": {
"historyId": {
"description": "The ID of the mailbox's current history record.",
"type": "string",
"format": "uint64"
},
"expiration": {
"description": "When Gmail will stop sending notifications for mailbox updates (epoch millis). Call `watch` again before this time to renew the watch.",
"type": "string",
"format": "int64"
}
}
}, */

func (m Model) Register(ctx context.Context) error {
	return errors.New("Not implemented")
}

func (m Model) Webhook() error {
	return errors.New("Not implemented")
}

func (m Model) Send(email Email) error {
	return errors.New("Not implemented")
}
