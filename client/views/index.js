import { h } from 'hyperapp';
import ChatPanel from './chat-panel';

const RenderChatPanel = (state, actions) => name => {
    if (!state.chat.view[name]) {
        actions.chat.view.init(name);
        return null;
    }

    return <ChatPanel 
        name={name}
        onchange={userInput => actions.chat.view.update({name, userInput})}
        onsend={text => actions.chat.sendToRoom({name, text})}
        logs={state.chat.logs[name]} 
        view={state.chat.view[name]}>
    </ChatPanel>;
};

export default (state, actions) => {
    return <div class="root-container" oncreate={e => actions.chat.connect(state.config)}>
        <div class="panel-container">
            { Object.getOwnPropertyNames(state.chat.logs)
                .filter(name => state.chat.logs[name])
                .map(RenderChatPanel(state, actions)) }
        </div>
    </div>;
};
