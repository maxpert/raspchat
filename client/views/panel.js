import { h } from 'hyperapp';

export default (props, children) => {
    return <div onclick={props.onClick} oncreate={props.init}>
        {children}
    </div>;
};
