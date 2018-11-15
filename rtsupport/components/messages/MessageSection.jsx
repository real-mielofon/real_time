import React, { Component } from 'react';
import PropTypes from 'prop-types';

import MessageForm from './MessageForm.jsx';
import MessageList from './MessageList.jsx';


class MessageSection extends Component {
    render() {
        let {activeChannel} = this.props;
        return (
            <div className="message-container card">
                <div className="card-header bg-primary text-white">
                    <strong>{activeChannel.name}</strong>
                </div>
                <div className="card-body messages">
                    <MessageList {...this.props}/>
                    <MessageForm {...this.props} />
                </div>
            </div>
        )
    }
}



MessageSection.propTypes = {
    messages: PropTypes.array.isRequired,
    activeChannel: PropTypes.object.isRequired,
    addMessage: PropTypes.func.isRequired,
}

export default MessageSection;