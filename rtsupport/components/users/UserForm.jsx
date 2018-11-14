import React, { Component } from 'react'
import PropTypes from 'prop-types';

class UserForm extends Component {
    onSubmit(e) {
        e.preventDefault();
        const node = this.refs.UserName;
        const userName = node.value;
        this.props.setUserName(userName);
        node.value = '';
    }
    render() {
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <div className="form-group">
                    <input
                        className="form-control"
                        placeholder="Set your name..."
                        type="text"
                        ref="UserName"
                    />
                </div>
            </form>
        )
    }
}

UserForm.propTypes = {
    setUserName: PropTypes.func.isRequired,
}

export default UserForm;