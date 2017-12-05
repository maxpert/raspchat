import { h } from 'hyperapp';
import Panel from './panel';
import Messages from './messages';

const sendWhenCommited = onSend => ev => {
    if (ev.keyCode == 13 && !ev.shiftKey) {
        const val = ev.target.value.trim();
        ev.preventDefault();
        ev.stopPropagation();
        if (val && onSend) {
            onSend(val, ev);
        }
    }
};

const notifyInputChange = cb => ev => cb(ev.target.value);

export default (props, children) => {
    return <Panel>
        <div class="panel-title">
            {props.name}
        </div>

        <Messages messages={props.logs}></Messages>
        <textarea 
            oninput={notifyInputChange(props.onchange)}
            onchange={notifyInputChange(props.onchange)}
            onkeydown={sendWhenCommited(props.onsend)}
            value={props.view.userInput}
            class="user-message"></textarea>
    </Panel>;
};
