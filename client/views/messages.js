import { h } from 'hyperapp';
import Markdown from './markdown';

export const Message = (props, children) => {
    return <Markdown class={props['@']} text={props.msg}></Markdown>;
};

export default (props, children) => {
    if (!props.messages) {
        return null;
    }

    return <div>
        { props.messages.map(m => <Message id={m['!id']} {...m}></Message>) }
    </div>;
};
