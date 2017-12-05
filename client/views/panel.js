import { h } from 'hyperapp';

export default (props, children) => {
    return <div class="panel">
        <div class="panel-content-container">
            {children}
        </div>
    </div>;
};
