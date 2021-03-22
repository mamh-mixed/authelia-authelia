package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/authelia/authelia/internal/authentication"
	"github.com/authelia/authelia/internal/authorization"
	"github.com/authelia/authelia/internal/logging"
	"github.com/authelia/authelia/internal/middlewares"
	"github.com/authelia/authelia/internal/session"
)

func authorizeHandler(ctx *middlewares.AutheliaCtx, rw http.ResponseWriter, r *http.Request) {
	ar, err := ctx.Providers.OpenIDConnect.Fosite.NewAuthorizeRequest(ctx, r)
	if err != nil {
		logging.Logger().Errorf("Error occurred in NewAuthorizeRequest: %+v", err)
		ctx.Providers.OpenIDConnect.Fosite.WriteAuthorizeError(rw, ar, err)

		return
	}

	clientID := ar.GetClient().GetID()

	if !ctx.Providers.OpenIDConnect.IsValidClientID(clientID) {
		err := fmt.Errorf("Unable to find related client configuration with name %s", ar.GetID())
		ctx.Logger.Error(err)
		ctx.Providers.OpenIDConnect.Fosite.WriteAuthorizeError(rw, ar, err)

		return
	}

	targetURL := ar.GetRedirectURI()
	userSession := ctx.GetSession()

	// Resolve the required level of authorizations to proceed with OIDC
	requiredAuthorizationLevel := ctx.Providers.Authorizer.GetRequiredLevel(authorization.Subject{
		Username: userSession.Username,
		Groups:   userSession.Groups,
		IP:       ctx.RemoteIP(),
	}, authorization.NewObjectRaw(targetURL, []byte("GET")))

	requestedScopes := ar.GetRequestedScopes()
	requestedAudience := ar.GetRequestedAudience()

	isAuthInsufficient := !ctx.Providers.OpenIDConnect.IsAuthenticationLevelSufficient(clientID, userSession.AuthenticationLevel)

	if isAuthInsufficient || (IsConsentMissing(userSession.OIDCWorkflowSession, requestedScopes, requestedAudience)) {
		forwardedURI, err := ctx.GetOriginalURL()
		if err != nil {
			ctx.Logger.Errorf("%v", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)

			return
		}

		if userSession.AuthenticationLevel == authentication.NotAuthenticated {
			// Reset all values from previous session before regenerating the cookie. We do this here because it's
			// skipped for the OIDC workflow on the 1FA post handler.
			err = ctx.SaveSession(session.NewDefaultUserSession())

			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)

				return
			}
		}

		ctx.Logger.Debugf("User %s must consent with scopes %s", userSession.Username, strings.Join(ar.GetRequestedScopes(), ", "))
		userSession.OIDCWorkflowSession = new(session.OIDCWorkflowSession)
		userSession.OIDCWorkflowSession.ClientID = clientID
		userSession.OIDCWorkflowSession.RequestedScopes = requestedScopes
		userSession.OIDCWorkflowSession.RequestedAudience = requestedAudience
		userSession.OIDCWorkflowSession.AuthURI = forwardedURI.String()
		userSession.OIDCWorkflowSession.TargetURI = targetURL.String()
		userSession.OIDCWorkflowSession.RequiredAuthorizationLevel = requiredAuthorizationLevel

		if err := ctx.SaveSession(userSession); err != nil {
			ctx.Logger.Errorf("%v", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)

			return
		}

		uri, err := ctx.ForwardedProtoHost()
		if err != nil {
			ctx.Logger.Errorf("%v", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)

			return
		}

		if isAuthInsufficient {
			http.Redirect(rw, r, uri, http.StatusFound)
		} else {
			http.Redirect(rw, r, fmt.Sprintf("%s/consent", uri), http.StatusFound)
		}

		return
	}

	for _, scope := range requestedScopes {
		ar.GrantScope(scope)
	}

	for _, a := range requestedAudience {
		ar.GrantAudience(a)
	}

	userSession.OIDCWorkflowSession = nil
	if err := ctx.SaveSession(userSession); err != nil {
		ctx.Logger.Errorf("%v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)

		return
	}

	oauthSession := newSession(ctx, ar)
	response, err := ctx.Providers.OpenIDConnect.Fosite.NewAuthorizeResponse(ctx, ar, oauthSession)

	if err != nil {
		ctx.Logger.Errorf("Error occurred in NewAuthorizeResponse: %+v", err)
		ctx.Providers.OpenIDConnect.Fosite.WriteAuthorizeError(rw, ar, err)

		return
	}

	ctx.Providers.OpenIDConnect.Fosite.WriteAuthorizeResponse(rw, ar, response)
}
