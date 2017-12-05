(function (global, factory) {
	typeof exports === 'object' && typeof module !== 'undefined' ? factory() :
	typeof define === 'function' && define.amd ? define(factory) :
	(factory());
}(this, (function () { 'use strict';

function app(props, container) {
  var root = (container = container || document.body).children[0];
  var node = vnode(root, [].map);
  var appState = {};
  var appActions = {};
  var lifecycle = [];
  var patchLock;

  repaint(init(appState, appActions, props, []));

  return appActions;

  function vnode(element, map) {
    return element && h(element.tagName.toLowerCase(), {}, map.call(element.childNodes, function (element) {
      return element.nodeType === 3 ? element.nodeValue : vnode(element, map);
    }));
  }

  function repaint() {
    if (props.view && !patchLock) {
      setTimeout(render, patchLock = !patchLock);
    }
  }

  function render(next) {
    patchLock = !patchLock;
    if ((next = props.view(appState, appActions)) && !patchLock) {
      root = patch(container, root, node, node = next);
    }
    while (next = lifecycle.pop()) {
      next();
    }
  }

  function init(state, actions, from, path) {
    var modules = from.modules;

    initDeep(state, actions, from.actions, path);
    set(state, from.state);

    for (var i in modules) {
      init(state[i] = {}, actions[i] = {}, modules[i], path.concat(i));
    }
  }

  function initDeep(state, actions, from, path) {
    Object.keys(from || {}).map(function (key) {
      if (typeof from[key] === "function") {
        actions[key] = function (data) {
          var result = from[key](state = get(path, appState), actions);

          if (typeof result === "function") {
            result = result(data);
          }

          if (result && result !== state && !result.then) {
            repaint(appState = setDeep(path, merge(state, result), appState));
          }

          return result;
        };
      } else {
        initDeep(state[key] || (state[key] = {}), actions[key] = {}, from[key], path.concat(key));
      }
    });
  }

  function merge(to, from) {
    return set(set({}, to), from);
  }

  function set(to, from) {
    for (var i in from) {
      to[i] = from[i];
    }
    return to;
  }

  function setDeep(path, value, from) {
    var to = {};
    return path.length === 0 ? value : (to[path[0]] = 1 < path.length ? setDeep(path.slice(1), value, from[path[0]]) : value, merge(from, to));
  }

  function get(path, from) {
    for (var i = 0; i < path.length; i++) {
      from = from[path[i]];
    }
    return from;
  }

  function createElement(node, isSVG) {
    if (typeof node === "string") {
      var element = document.createTextNode(node);
    } else {
      var element = (isSVG = isSVG || node.type === "svg") ? document.createElementNS("http://www.w3.org/2000/svg", node.type) : document.createElement(node.type);

      if (node.props.oncreate) {
        lifecycle.push(function () {
          node.props.oncreate(element);
        });
      }

      for (var i = 0; i < node.children.length; i++) {
        element.appendChild(createElement(node.children[i], isSVG));
      }

      for (var i in node.props) {
        setElementProp(element, i, node.props[i]);
      }
    }
    return element;
  }

  function setElementProp(element, name, value, oldValue) {
    if (name === "key") {} else if (name === "style") {
      for (var name in merge(oldValue, value = value || {})) {
        element.style[name] = value[name] || "";
      }
    } else {
      try {
        element[name] = null == value ? "" : value;
      } catch (_) {}

      if (typeof value !== "function") {
        if (null == value || false === value) {
          element.removeAttribute(name);
        } else {
          element.setAttribute(name, value);
        }
      }
    }
  }

  function updateElement(element, oldProps, props) {
    for (var i in merge(oldProps, props)) {
      var value = props[i];
      var oldValue = i === "value" || i === "checked" ? element[i] : oldProps[i];

      if (value !== oldValue) {
        setElementProp(element, i, value, oldValue);
      }
    }

    if (props.onupdate) {
      lifecycle.push(function () {
        props.onupdate(element, oldProps);
      });
    }
  }

  function removeElement(parent, element, props) {
    if (props && props.onremove) {
      props.onremove(element, done);
    } else {
      done();
    }

    function done() {
      parent.removeChild(element);
    }
  }

  function getKey(node) {
    if (node && node.props) {
      return node.props.key;
    }
  }

  function patch(parent, element, oldNode, node, isSVG, nextSibling) {
    if (oldNode === node) {} else if (oldNode == null) {
      element = parent.insertBefore(createElement(node, isSVG), element);
    } else if (node.type != null && node.type === oldNode.type) {
      updateElement(element, oldNode.props, node.props);

      isSVG = isSVG || node.type === "svg";

      var len = node.children.length;
      var oldLen = oldNode.children.length;
      var oldKeyed = {};
      var oldElements = [];
      var keyed = {};

      for (var i = 0; i < oldLen; i++) {
        var oldElement = oldElements[i] = element.childNodes[i];
        var oldChild = oldNode.children[i];
        var oldKey = getKey(oldChild);

        if (null != oldKey) {
          oldKeyed[oldKey] = [oldElement, oldChild];
        }
      }

      var i = 0;
      var j = 0;

      while (j < len) {
        var oldElement = oldElements[i];
        var oldChild = oldNode.children[i];
        var newChild = node.children[j];

        var oldKey = getKey(oldChild);
        if (keyed[oldKey]) {
          i++;
          continue;
        }

        var newKey = getKey(newChild);
        var keyedNode = oldKeyed[newKey] || [];

        if (null == newKey) {
          if (null == oldKey) {
            patch(element, oldElement, oldChild, newChild, isSVG);
            j++;
          }
          i++;
        } else {
          if (oldKey === newKey) {
            patch(element, keyedNode[0], keyedNode[1], newChild, isSVG);
            i++;
          } else if (keyedNode[0]) {
            element.insertBefore(keyedNode[0], oldElement);
            patch(element, keyedNode[0], keyedNode[1], newChild, isSVG);
          } else {
            patch(element, oldElement, null, newChild, isSVG);
          }

          j++;
          keyed[newKey] = newChild;
        }
      }

      while (i < oldLen) {
        var oldChild = oldNode.children[i];
        var oldKey = getKey(oldChild);
        if (null == oldKey) {
          removeElement(element, oldElements[i], oldChild.props);
        }
        i++;
      }

      for (var i in oldKeyed) {
        var keyedNode = oldKeyed[i];
        var reusableNode = keyedNode[1];
        if (!keyed[reusableNode.props.key]) {
          removeElement(element, keyedNode[0], reusableNode.props);
        }
      }
    } else if (element && node !== element.nodeValue) {
      if (typeof node === "string" && typeof oldNode === "string") {
        element.nodeValue = node;
      } else {
        element = parent.insertBefore(createElement(node, isSVG), nextSibling = element);
        removeElement(parent, nextSibling, oldNode.props);
      }
    }

    return element;
  }
}

function h(type, props) {
  var node;
  var stack = [];
  var children = [];

  for (var i = arguments.length; i-- > 2;) {
    stack.push(arguments[i]);
  }

  while (stack.length) {
    if (Array.isArray(node = stack.pop())) {
      for (i = node.length; i--;) {
        stack.push(node[i]);
      }
    } else if (node != null && node !== true && node !== false) {
      children.push(typeof node === "number" ? node = node + "" : node);
    }
  }

  return typeof type === "string" ? {
    type: type,
    props: props || {},
    children: children
  } : type(props || {}, children);
}

var settings = {
    toggleSound: function toggleSound(state, actions) {
        return function (value) {
            var soundEnabled = value === undefined ? !state.soundEnabled : value;
            return Promise.resolve({ soundEnabled: soundEnabled }).then(actions.applyState);
        };
    },

    toggleNotification: function toggleNotification(state, actions) {
        return function (value) {
            var notificationEnabled = value === undefined ? !state.soundEnabled : value;
            return Promise.resolve({ notificationEnabled: notificationEnabled }).then(actions.applyState);
        };
    },

    applyState: function applyState(state, actions) {
        return function (state) {
            return state;
        };
    }
};

var config = {
    load: function load(state, actions) {
        return function (c) {
            return c;
        };
    }
};

var DEFAULT_CHANNEL = 'SERVER';

var MessageCommonProperties = function MessageCommonProperties(m) {
    return {
        id: m['!id'],
        timestamp: m.utc_timestamp,
        command: m['@'].toLowerCase()
    };
};

var _extends = Object.assign || function (target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i];

    for (var key in source) {
      if (Object.prototype.hasOwnProperty.call(source, key)) {
        target[key] = source[key];
      }
    }
  }

  return target;
};

var _LogStateCommands;

var LogStateCommands = (_LogStateCommands = {
    'pong': function pong(m) {
        return null;
    },
    'nick-set': function nickSet(m) {
        return null;
    },

    'default': function _default(message) {
        console.error('Invalid message', message);
        return null;
    },

    'client-meta': function clientMeta(m) {
        return _extends({}, MessageCommonProperties(m), {
            to: DEFAULT_CHANNEL,
            from: DEFAULT_CHANNEL,
            type: 'meta',
            command: 'client-meta',
            message: m.msg
        });
    }

}, _LogStateCommands[DEFAULT_CHANNEL.toLowerCase()] = function (m) {
    return _extends({}, MessageCommonProperties(m), {
        to: DEFAULT_CHANNEL,
        from: DEFAULT_CHANNEL,
        message: m.msg,
        type: 'meta'
    });
}, _LogStateCommands['group-list'] = function groupList(m) {
    return _extends({}, MessageCommonProperties(m), {
        to: m.to,
        from: m.from,
        message: 'members: ' + m.pack_msg.join(', '),
        type: 'meta'
    });
}, _LogStateCommands['group-join'] = function groupJoin(m) {
    return _extends({}, MessageCommonProperties(m), {
        to: m.to,
        from: m.from,
        message: m.from + ' has joined ' + m.to,
        type: 'meta'
    });
}, _LogStateCommands['group-leave'] = function groupLeave(m) {
    return _extends({}, MessageCommonProperties(m), {
        to: m.to,
        from: m.from,
        message: m.from + ' has left ' + m.to,
        type: 'meta'
    });
}, _LogStateCommands['member-nick-set'] = function memberNickSet(m) {
    return _extends({}, MessageCommonProperties(m), {
        to: m.to,
        from: DEFAULT_CHANNEL,
        message: m.pack_msg.oldNick + ' changed nick to ' + m.pack_msg.newNick,
        type: 'meta'
    });
}, _LogStateCommands['group-message'] = function groupMessage(m) {
    return _extends({}, MessageCommonProperties(m), {
        to: m.to,
        from: m.from,
        message: m.msg
    });
}, _LogStateCommands);

var messageParser = function (logs, message) {
    var _ref2;

    var command = message['@'].toLowerCase();
    var targetChannel = message.to || DEFAULT_CHANNEL;
    var channelLog = logs[targetChannel] || [];
    var parser = LogStateCommands[command] || LogStateCommands['default'];
    var parsedMessage = parser(message);

    if (!parsedMessage) {
        var _ref;

        return _ref = {}, _ref[targetChannel] = channelLog, _ref;
    }

    return _ref2 = {}, _ref2[targetChannel] = channelLog.concat([_extends({}, parsedMessage, {
        _raw: message
    })]), _ref2;
};

var DefaultViewState = { active: false, userInput: '' };

var view = {
    init: function init(state, actions) {
        return function (name) {
            var _babelHelpers$extends;

            return _extends({}, state, (_babelHelpers$extends = {}, _babelHelpers$extends[name] = _extends({}, DefaultViewState), _babelHelpers$extends));
        };
    },

    clear: function clear(state, actions) {
        return function (name) {
            var _babelHelpers$extends2;

            return _extends({}, state, (_babelHelpers$extends2 = {}, _babelHelpers$extends2[name] = null, _babelHelpers$extends2));
        };
    },

    reset: function reset(state, actions) {
        return function (name) {
            return actions.update({ name: name, userInput: '' });
        };
    },

    update: function update(state, actions) {
        return function (_ref) {
            var _babelHelpers$extends3;

            var name = _ref.name,
                userInput = _ref.userInput,
                active = _ref.active;

            var old = state[name] || _extends({}, DefaultViewState);
            if (userInput === undefined) userInput = old.userInput;
            if (active === undefined) active = old.active;

            var newState = _extends({}, state, (_babelHelpers$extends3 = {}, _babelHelpers$extends3[name] = { userInput: userInput, active: active }, _babelHelpers$extends3));

            return newState;
        };
    }
};

var logs = {
    add: function add(state, actions) {
        return function (message) {
            return messageParser(state, message);
        };
    },
    clear: function clear(state, actions) {
        return function (name) {
            var _babelHelpers$extends;

            return _extends({}, state, (_babelHelpers$extends = {}, _babelHelpers$extends[name] = null, _babelHelpers$extends));
        };
    }
};

var messageId = 0;

function serverCommandMessage(command, msg) {
    return {
        '@': command,
        to: DEFAULT_CHANNEL,
        msg: msg
    };
}

function meta(msg, extra) {
    extra = extra || {};

    return _extends({
        '@': 'client-meta',
        '!id': 'm' + messageId++,
        'utc_timestamp': new Date().getTime(),
        to: DEFAULT_CHANNEL,
        msg: msg
    }, extra);
}

function join(room, extra) {
    extra = extra || {};
    return _extends({}, serverCommandMessage('join-group', room), extra);
}

function leave(room, extra) {
    extra = extra || {};
    return _extends({}, serverCommandMessage('leave-group', room), extra);
}

function nick(nick, extra) {
    extra = extra || {};
    return _extends({}, serverCommandMessage('set-nick', nick), extra);
}

function list(room, extra) {
    extra = extra || {};
    return _extends({}, serverCommandMessage('list-group', room), extra);
}

function message(room, message, extra) {
    extra = extra || {};
    return _extends({
        '@': 'send-msg',
        to: room,
        msg: message
    }, extra);
}

var commandRegex = /^\/(join|leave|nick)\s+(.+)$/i;
var commandTransformMap = {
    'join': function join$$1(r) {
        return join(r);
    },
    'leave': function leave$$1(r) {
        return leave(r);
    },
    'nick': function nick$$1(r) {
        return nick(r);
    }
};

var UserInputParser = function (roomName, text) {
    var regexMatch = commandRegex.exec(text);
    if (regexMatch && commandTransformMap[regexMatch[1]]) {
        return commandTransformMap[regexMatch[1]](regexMatch[2].trim());
    }

    var trimmedText = text.trim().toLowerCase();
    if (trimmedText === '/leave') {
        return leave(roomName);
    }

    if (trimmedText === '/list') {
        return list(roomName);
    }

    // Parse rest of commands here...
    return message(roomName, text);
};

var HelpMessage = '\n```\nSystem connected, valid commands:\n\n/join <room> - to join a room\n/leave <room> - to leave a room \n/nick <name> -  to set your nick\n/list - to list room members\n\n```\n\n# \xA0\n# Welcome to RaspChat\n\xA0\xA0\xA0\n';

var chat = {
    view: view,
    logs: logs,

    connect: function connect(state, actions) {
        return function (config) {
            var connection = new WebSocket(config.webSocketConnectionUri);
            connection.onopen = actions.onOpen(connection);
            connection.onclose = actions.onClose(connection);
            connection.onmessage = actions.onMessage(connection);

            return { connectionAttempts: state.connectionAttempts + 1, connection: connection };
        };
    },

    sendMessageObject: function sendMessageObject(state, actions) {
        return function (message$$1) {
            if (!state.connection) {
                return {};
            }

            try {
                state.connection.send(JSON.stringify(message$$1));
            } catch (e) {
                console.error('Unable to send message', e);
            }

            return {};
        };
    },

    sendToRoom: function sendToRoom(state, actions) {
        return function (_ref) {
            var roomName = _ref.name,
                text = _ref.text;

            return Promise.resolve(UserInputParser(roomName, text)).then(function (m) {
                actions.view.reset(roomName);
                actions.sendMessageObject(m);
                return m;
            }).then(function (_ref2) {
                var command = _ref2['@'],
                    group = _ref2['msg'];

                if (command === 'leave-group') {
                    actions.logs.clear(group);
                    actions.view.clear(group);
                }
            });
        };
    },

    stayConnected: function stayConnected(state, actions) {
        return function (_ref3) {
            var timeout = _ref3.timeout,
                config = _ref3.config;
            return function () {
                var c = state.connection;
                var alreadyConnected = c && (c.readyState === 1 || c.readyState === 0);

                // Do not retry if already connected or a connection was never attempted
                if (state.connectionAttempts === 0 || alreadyConnected) {
                    window.setTimeout(actions.stayConnected({ config: config, timeout: timeout }), timeout);
                    return {};
                }

                var ret = actions.connect(config);
                window.setTimeout(actions.stayConnected({ config: config, timeout: timeout }), timeout);
                return ret;
            };
        };
    },

    onOpen: function onOpen(state, actions) {
        return function (connection) {
            return function (event) {
                actions.logs.add(meta(HelpMessage));
                window.setTimeout(function () {
                    actions.sendMessageObject(join('Lounge'));
                }, 1);
                return { connectionAttempts: 1, connection: connection };
            };
        };
    },

    onClose: function onClose(state, actions) {
        return function (connection) {
            return function (event) {
                console.error('Connection closed', event);
                return { connection: connection };
            };
        };
    },

    onMessage: function onMessage(state, actions) {
        return function (connection) {
            return function (event) {
                try {
                    var msg = JSON.parse(event.data);
                    actions.logs.add(msg);
                } catch (e) {
                    console.warn('Error while parsing message', e, event.data);
                }

                return { connection: connection };
            };
        };
    }
};

var actions = {
    config: config,
    settings: settings,
    chat: chat
};

var state = {
    config: {},
    settings: {
        soundEnabled: true,
        notificationEnabled: true
    },
    chat: {
        connectionAttempts: 0,
        connection: null,
        view: {},
        logs: {}
    }
};

var Panel = (function (props, children) {
    return h(
        "div",
        { "class": "panel" },
        h(
            "div",
            { "class": "panel-content-container" },
            children
        )
    );
});

var TAGS = {
	'': ['<em>', '</em>'],
	_: ['<strong>', '</strong>'],
	'\n': ['<br />'],
	' ': ['<br />'],
	'-': ['<hr />']
};

/** Outdent a string based on the first indented line's leading whitespace
 *	@private
 */
function outdent(str) {
	return str.replace(RegExp('^' + (str.match(/^(\t| )+/) || '')[0], 'gm'), '');
}

/** Encode special attribute characters to HTML entities in a String.
 *	@private
 */
function encodeAttr(str) {
	return (str + '').replace(/"/g, '&quot;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
}

/** Parse Markdown into an HTML String. */
function parse(md) {
	var tokenizer = /((?:^|\n+)(?:\n---+|\* \*(?: \*)+)\n)|(?:^```(\w*)\n([\s\S]*?)\n```$)|((?:(?:^|\n+)(?:\t|  {2,}).+)+\n*)|((?:(?:^|\n)([>*+-]|\d+\.)\s+.*)+)|(?:\!\[([^\]]*?)\]\(([^\)]+?)\))|(\[)|(\](?:\(([^\)]+?)\))?)|(?:(?:^|\n+)([^\s].*)\n(\-{3,}|={3,})(?:\n+|$))|(?:(?:^|\n+)(#{1,3})\s*(.+)(?:\n+|$))|(?:`([^`].*?)`)|(  \n\n*|\n{2,}|__|\*\*|[_*])/gm,
	    context = [],
	    out = '',
	    last = 0,
	    links = {},
	    chunk,
	    prev,
	    token,
	    inner,
	    t;

	function tag(token) {
		var desc = TAGS[token.replace(/\*/g, '_')[1] || ''],
		    end = context[context.length - 1] == token;
		if (!desc) {
			return token;
		}
		if (!desc[1]) {
			return desc[0];
		}
		context[end ? 'pop' : 'push'](token);
		return desc[end | 0];
	}

	function flush() {
		var str = '';
		while (context.length) {
			str += tag(context[context.length - 1]);
		}
		return str;
	}

	md = md.replace(/^\[(.+?)\]:\s*(.+)$/gm, function (s, name, url) {
		links[name.toLowerCase()] = url;
		return '';
	}).replace(/^\n+|\n+$/g, '');

	while (token = tokenizer.exec(md)) {
		prev = md.substring(last, token.index);
		last = tokenizer.lastIndex;
		chunk = token[0];
		if (prev.match(/[^\\](\\\\)*\\$/)) {}
		// escaped

		// Code/Indent blocks:
		else if (token[3] || token[4]) {
				chunk = '<pre class="code ' + (token[4] ? 'poetry' : token[2].toLowerCase()) + '">' + outdent(encodeAttr(token[3] || token[4]).replace(/^\n+|\n+$/g, '')) + '</pre>';
			}
			// > Quotes, -* lists:
			else if (token[6]) {
					t = token[6];
					if (t.match(/\./)) {
						token[5] = token[5].replace(/^\d+/gm, '');
					}
					inner = parse(outdent(token[5].replace(/^\s*[>*+.-]/gm, '')));
					if (t === '>') {
						t = 'blockquote';
					} else {
						t = t.match(/\./) ? 'ol' : 'ul';
						inner = inner.replace(/^(.*)(\n|$)/gm, '<li>$1</li>');
					}
					chunk = '<' + t + '>' + inner + '</' + t + '>';
				}
				// Images:
				else if (token[8]) {
						chunk = "<img src=\"" + encodeAttr(token[8]) + "\" alt=\"" + encodeAttr(token[7]) + "\">";
					}
					// Links:
					else if (token[10]) {
							out = out.replace('<a>', "<a href=\"" + encodeAttr(token[11] || links[prev.toLowerCase()]) + "\">");
							chunk = flush() + '</a>';
						} else if (token[9]) {
							chunk = '<a>';
						}
						// Headings:
						else if (token[12] || token[14]) {
								t = 'h' + (token[14] ? token[14].length : token[13][0] === '=' ? 1 : 2);
								chunk = '<' + t + '>' + parse(token[12] || token[15]) + '</' + t + '>';
							}
							// `code`:
							else if (token[16]) {
									chunk = '<code>' + encodeAttr(token[16]) + '</code>';
								}
								// Inline formatting: *em*, **strong** & friends
								else if (token[17] || token[1]) {
										chunk = tag(token[17] || '--');
									}
		out += prev;
		out += chunk;
	}

	return (out + md.substring(last) + flush()).trim();
}


//# sourceMappingURL=snarkdown.es.js.map

var applyMarkdownHtml = function applyMarkdownHtml(text) {
    return function (e) {
        if (!text) {
            return;
        }

        e.innerHTML = parse(text);
    };
};

var Markdown = (function (props) {
    return h('div', { 'class': props.type + '-markdown',
        oncreate: applyMarkdownHtml(props.md || '') });
});

var Message = function Message(props, children) {
    return h(
        'div',
        { 'class': props.command + '-message message-log' },
        h(
            'div',
            { 'class': 'from-container' },
            props.from
        ),
        h(
            'div',
            { 'class': 'message-container' },
            h(Markdown, { type: props.command, md: props.message })
        )
    );
};

var Messages = (function (props, children) {
    if (!props.messages) {
        return null;
    }

    return h(
        'div',
        { 'class': 'log-messages ' + (props.class || '') },
        props.messages.map(function (m) {
            return h(Message, _extends({ id: m.id }, m));
        })
    );
});

var sendWhenCommited = function sendWhenCommited(onSend) {
    return function (ev) {
        if (ev.keyCode == 13 && !ev.shiftKey) {
            var val = ev.target.value.trim();
            ev.preventDefault();
            ev.stopPropagation();
            if (val && onSend) {
                onSend(val, ev);
            }
        }
    };
};

var notifyInputChange = function notifyInputChange(cb) {
    return function (ev) {
        return cb(ev.target.value);
    };
};

var ChatPanel = (function (props, children) {
    return h(
        Panel,
        null,
        h(
            'div',
            { 'class': 'panel-title' },
            props.name
        ),
        h(Messages, { messages: props.logs }),
        h('textarea', {
            oninput: notifyInputChange(props.onchange),
            onchange: notifyInputChange(props.onchange),
            onkeydown: sendWhenCommited(props.onsend),
            value: props.view.userInput,
            'class': 'user-message' })
    );
});

var RenderChatPanel = function RenderChatPanel(state, actions) {
    return function (name) {
        if (!state.chat.view[name]) {
            actions.chat.view.init(name);
            return null;
        }

        return h(ChatPanel, {
            name: name,
            onchange: function onchange(userInput) {
                return actions.chat.view.update({ name: name, userInput: userInput });
            },
            onsend: function onsend(text) {
                return actions.chat.sendToRoom({ name: name, text: text });
            },
            logs: state.chat.logs[name],
            view: state.chat.view[name] });
    };
};

var view$1 = (function (state, actions) {
    return h(
        'div',
        { 'class': 'root-container', oncreate: function oncreate(e) {
                return actions.chat.connect(state.config);
            } },
        h(
            'div',
            { 'class': 'panel-container' },
            Object.getOwnPropertyNames(state.chat.logs).filter(function (name) {
                return state.chat.logs[name];
            }).map(RenderChatPanel(state, actions))
        )
    );
});

var boundActions = app({ state: state, actions: actions, view: view$1 });
boundActions.config.load(window['RaspConfig']);
boundActions.chat.stayConnected({ timeout: 2000, config: window['RaspConfig'] })();

})));
//# sourceMappingURL=app.js.map
