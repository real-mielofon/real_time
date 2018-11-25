import React, { Component } from 'react'
import PropTypes from 'prop-types'
import fecha from 'fecha'

class Message extends Component {
    render() {
        let {message} = this.props
        console.log("render message = "+message)
        let createdAtTime = fecha.parse(message.createdAt, "YYYY-MM-DDTHH:mm:ss.SSSZZ")
        let createdAt = fecha.format(createdAtTime, "YYYY-MM-DD hh:mm:ss")
//        let createdAt = message.createdAt
        console.log("render message created = "+createdAt)
        return (
            <li className="message">
                <div className="author">
                    <strong>{message.author}</strong>
                    <i className="timestamp">{createdAt}</i> 
                </div>
                <div className="body">
                    {message.body}
                </div>
            </li>
        )
    }
}

Message.propTypes = {
    message: PropTypes.object.isRequired
}

export default Message