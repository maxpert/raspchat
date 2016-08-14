package rica

import (
    "bytes"
    "crypto/tls"
    "encoding/json"
    "errors"
    "net/http"
    "sync"
)

type GCMDeliveryWork struct {
    DeliveryIds []string
    Message     interface{}
}

type GCMWorker struct {
    sync.Mutex
    deliveryMap map[uint64]*GCMDeliveryWork
    serverKey   string
}

type GCMJsonMessage struct {
    RegistrationIds       []string    `json:"registration_ids"`
    CollapseKey           string      `json:"collapse_key,omitempty"`
    Priority              string      `json:"priority,omitempty"`
    ContentAvailable      string      `json:"content_available,omitempty"`
    DelayWhileIdle        string      `json:"delay_while_idle,omitempty"`
    TTL                   uint32      `json:"time_to_live"`
    RestrictedPackageName string      `json:"restricted_package_name,omitempty"`
    DryRun                bool        `json:"dry_run"`
    Data                  interface{} `json:"data"`
    Notification          interface{} `json:"notification"`
}

func NewGCMWorker(serverKey string) *GCMWorker {
    return &GCMWorker{
        deliveryMap: make(map[uint64]*GCMDeliveryWork),
        serverKey:   serverKey,
    }
}

func (g *GCMWorker) Enqueue(to string, id uint64, message interface{}) {
    g.Lock()

    var work *GCMDeliveryWork = nil
    ok := false
    if work, ok = g.deliveryMap[id]; !ok {
        work = &GCMDeliveryWork{
            DeliveryIds: []string{},
            Message:     message,
        }
    }

    work.DeliveryIds = append(work.DeliveryIds, to)
    g.deliveryMap[id] = work
    g.Unlock()
}

func (g *GCMWorker) Deliver(id uint64) error {
    work := g.dequeue(id)

    if work == nil {
        return nil
    }

    msg, err := json.Marshal(work.Message)
    if err != nil {
        return err
    }

    return g.sendPushRequest(work.DeliveryIds, msg)
}

func (g *GCMWorker) dequeue(id uint64) *GCMDeliveryWork {
    g.Lock()
    defer g.Unlock()

    if work, ok := g.deliveryMap[id]; ok {
        delete(g.deliveryMap, id)
        return work
    }

    return nil
}

func (g *GCMWorker) sendPushRequest(registrationIds []string, body []byte) error {
    strBody := string(body)
    gReq := &GCMJsonMessage{
        Data: map[string]string{
            "json_msg": strBody,
        },
        RegistrationIds: registrationIds,
        TTL:             3 * 60 * 60,
    }
    gReqBytes, err := json.Marshal(gReq)
    if err != nil {
        return err
    }

    postBody := bytes.NewBuffer(gReqBytes)

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := http.Client{Transport: tr}

    req, err := http.NewRequest("POST", "https://android.googleapis.com/gcm/send", postBody)
    if err != nil {
        return err
    }

    req.Header.Add("Authorization", "key="+g.serverKey)
    req.Header.Add("Content-Type", "application/json")

    resp, err := client.Do(req)

    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return errors.New(resp.Status)
    }

    return nil
}
