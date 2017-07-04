package rascore

import (
    "sync"
)

type gcmTransportContainer struct {
    sync.Mutex
    transportMap map[string]chan string
}

var pGCMTransportContainer *gcmTransportContainer = &gcmTransportContainer{
    transportMap: make(map[string]chan string),
}

func (p *gcmTransportContainer) saveChannelForId(id string, c chan string) {
    p.Lock()
    defer p.Unlock()

    p.transportMap[id] = c
}

func (p *gcmTransportContainer) updateChannelForId(oldId, newId string) bool {
    p.Lock()
    defer p.Unlock()

    if c, ok := p.transportMap[oldId]; ok {
        p.transportMap[newId] = c
        delete(p.transportMap, oldId)
        return true
    }

    return false
}

func (p *gcmTransportContainer) deleteChannelForId(id string) {
    p.Lock()
    defer p.Unlock()

    delete(p.transportMap, id)
}

func (p *gcmTransportContainer) getChannelForId(id string) (chan string, bool) {
    p.Lock()
    defer p.Unlock()

    if c, ok := p.transportMap[id]; ok {
        return c, true
    }

    return make(chan string), false
}

type GCMTransport struct {
    clientId       string
    worker         *GCMWorker
    request        chan string
    pendingBatches map[uint64]interface{}
}

func NewGCMTransport(client string, worker *GCMWorker) *GCMTransport {
    ch, _ := pGCMTransportContainer.getChannelForId(client)

    t := &GCMTransport{
        clientId:       client,
        worker:         worker,
        request:        ch,
        pendingBatches: make(map[uint64]interface{}),
    }

    pGCMTransportContainer.saveChannelForId(client, t.request)
    return t
}

func (h *GCMTransport) ReadMessage() (IEventMessage, error) {
    str := <-h.request
    return transportDecodeMessage([]byte(str))
}

func (h *GCMTransport) WriteMessage(id uint64, msg IEventMessage) error {
    h.worker.Enqueue(h.clientId, id, msg)
    return nil
}

var __empty struct{}

func (h *GCMTransport) BeginBatch(id uint64, msg IEventMessage) {
    h.pendingBatches[id] = __empty
}

func (h *GCMTransport) FlushBatch(id uint64) {
    go h.worker.Deliver(id)
}

func (h *GCMTransport) PostMessage(msg string) {
    h.request <- msg
}

func (h *GCMTransport) Unregister() {
    pGCMTransportContainer.deleteChannelForId(h.clientId)
}
