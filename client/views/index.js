import { h } from 'hyperapp';
import Panel from './panel';
import Messages from './messages';

export default (state, actions) => {
    return <Panel init={e => actions.channels.connect(state.config)}>
        <Messages messages={state.channels.logs.SERVER}></Messages>
    </Panel>;
};
