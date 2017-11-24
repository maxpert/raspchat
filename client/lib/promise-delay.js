export default function (delay, maybeValue) {
    return new Promise(ok => 
        window.setTimeout(() => ok(maybeValue), delay)
    );
}
