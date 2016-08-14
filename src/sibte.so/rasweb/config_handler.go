package rasweb

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"

    "github.com/julienschmidt/httprouter"
    "sibte.so/rasconfig"
)

type configRouteHandler struct {
}

// NewConfigRouteHandler returns instance of configuration route handler
func NewConfigRouteHandler() RouteHandler {
    return &configRouteHandler{}
}

func (h *configRouteHandler) Register(r *httprouter.Router) error {
    r.GET("/config/:type", h.getChatConfigurationHalder)
    return nil
}

// GetChatConfigurationHalder handles the /config/client.(js|json) calls
func (h *configRouteHandler) getChatConfigurationHalder(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    appConfig := rasconfig.CurrentAppConfig
    isJs := false

    if strings.HasSuffix(params.ByName("type"), ".js") {
        isJs = true
    }

    if isJs {
        w.Header().Add("Content-Type", "text/javascript")
    } else {
        w.Header().Add("Content-Type", "application/json")
    }

    config := make(map[string]interface{})
    config["webSocketConnectionUri"] = appConfig.WebSocketURL
    config["webSocketSecureConnectionUri"] = appConfig.WebSocketSecureURL
    config["externalSignIn"] = appConfig.ExternalSignIn
    config["hasAuthProviders"] = appConfig.HasAuthProviders

    if isJs {
        fmt.Fprint(w, "window.RaspConfig=")
    }

    json.NewEncoder(w).Encode(config)
}
