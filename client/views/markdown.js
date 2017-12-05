import { h } from 'hyperapp';
import snarkdown from 'snarkdown';

const applyMarkdownHtml = text => e => {
    if (!text) {
        return;
    }

    e.innerHTML = snarkdown(text);
};

export default (props) => (
    <div class={props.type + '-markdown'} 
        oncreate={applyMarkdownHtml(props.md || '')}>
    </div>
);
