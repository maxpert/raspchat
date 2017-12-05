import { h } from 'hyperapp';
import Markdown from './markdown';

export const Message = (props, children) => {
    return <div class={`${props.command}-message message-log`}>
        <div class="from-container">{props.from}</div>
        <div class="message-container">
            <Markdown type={props.command} md={props.message}></Markdown>
        </div>
    </div>;
};

export default (props, children) => {
    if (!props.messages) {
        return null;
    }

    return <div class={'log-messages ' + (props.class || '')}>
        { props.messages.map(m => <Message id={m.id} {...m}></Message>) }
    </div>;
};
